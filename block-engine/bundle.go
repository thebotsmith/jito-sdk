package block_engine

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	bundle_pb "github.com/Prophet-Solutions/jito-sdk/pb/bundle"
	searcher_pb "github.com/Prophet-Solutions/jito-sdk/pb/searcher"
	"github.com/Prophet-Solutions/jito-sdk/pkg"
	"github.com/gagliardetto/solana-go"
	"google.golang.org/grpc"
)

// Constants for retry and timeout configurations
const (
	CheckBundleRetries               = 10               // Number of times to retry checking bundle status
	CheckBundleRetryDelay            = 2 * time.Second  // Delay between retries for checking bundle status
	SignaturesConfirmationTimeout    = 15 * time.Second // Timeout for confirming signatures
	SignaturesConfirmationRetryDelay = 1 * time.Second  // Delay between retries for confirming signatures
)

// SendBundleWithConfirmation sends a bundle of transactions and waits for confirmation of signatures.
// It attempts to send the bundle, then continuously checks for the result of the bundle and validates
// the signatures of the transactions.
/*
func (c *SearcherClient) SendBundleWithConfirmation(
	ctx context.Context,
	transactions []*solana.Transaction,
	opts ...grpc.CallOption,
) (*BundleResponse, error) {
	// Send the bundle of transactions
	resp, err := c.SendBundle(transactions, opts...)
	if err != nil {
		return nil, err
	}

	// Retry checking the bundle result up to a configured number of times
	for i := 0; i < CheckBundleRetries; i++ {
		select {
		case <-c.AuthenticationService.GRPCCtx.Done():
			// If the GRPC context is done, return the error
			return nil, c.AuthenticationService.GRPCCtx.Err()
		default:
			// Wait for a configured delay before retrying
			time.Sleep(CheckBundleRetryDelay)

			// Attempt to receive the bundle result
			bundleResult, err := c.receiveBundleResult()
			if err != nil {
				//log.Println("error while receiving bundle result:", err)
			} else {
				// Handle the received bundle result
				if err = c.handleBundleResult(bundleResult); err != nil {
					if strings.Contains(err.Error(), "has already been processed") {
						return &BundleResponse{
							BundleResponse: resp,
							Signatures:     pkg.BatchExtractSigFromTx(transactions),
						}, nil
					}
					return nil, err
				}

				log.Println("Bundle was sent.")
			}

			// Wait for the statuses of the transaction signatures
			statuses, err := c.waitForSignatureStatuses(ctx, transactions)
			if err != nil {
				continue
			}

			// Validate the received signature statuses
			if err = pkg.ValidateSignatureStatuses(statuses); err != nil {
				continue
			}

			// Return the successful bundle response with extracted signatures
			return &BundleResponse{
				BundleResponse: resp,
				Signatures:     pkg.BatchExtractSigFromTx(transactions),
			}, nil
		}
	}

	// If the retries are exhausted, return an error
	return nil, fmt.Errorf("BroadcastBundleWithConfirmation error: max retries (%d) exceeded", CheckBundleRetries)
}
*/
func (c *SearcherClient) SendBundleWithConfirmation(
	ctx context.Context,
	transactions []*solana.Transaction,
	signature *solana.Signature,
	opts ...grpc.CallOption,
) (*BundleResponse, error) {
	var reachBlockChain bool = false
	// Send the bundle of transactions
	resp, err := c.SendBundle(transactions, opts...)
	if err != nil {
		return nil, err
	}

	fmt.Printf("tx (%s) sent...", signature)

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	// Create channels to handle the normal check and additional check
	normalCheckDone := make(chan *BundleResponse, 1)
	additionalCheckDone := make(chan *BundleResponse, 1)

	// Start the normal check in a separate goroutine
	go func() {
		for i := 0; i < CheckBundleRetries; i++ {
			select {
			case <-c.AuthenticationService.GRPCCtx.Done():
				// If the GRPC context is done, return the error
				normalCheckDone <- nil
				return
			default:
				// Wait for a configured delay before retrying
				time.Sleep(CheckBundleRetryDelay)

				// Attempt to receive the bundle result
				bundleResult, err := c.receiveBundleResult()
				if err != nil {
					log.Println("error while receiving bundle result:", err)
					normalCheckDone <- nil
				} else {
					// Handle the received bundle result
					if err = c.handleBundleResult(bundleResult); err != nil {
						if strings.Contains(err.Error(), "has already been processed") {
							normalCheckDone <- &BundleResponse{
								BundleResponse: resp,
								Signatures:     pkg.BatchExtractSigFromTx(transactions),
							}
							return
						}
						normalCheckDone <- nil
						return
					}
					log.Println("Bundle was sent.")
				}

				// Wait for the statuses of the transaction signatures
				statuses, err := c.waitForSignatureStatuses(ctx, transactions)
				if err != nil {
					continue
				}

				// Validate the received signature statuses
				if err = pkg.ValidateSignatureStatuses(statuses); err != nil {
					continue
				}

				// Return the successful bundle response with extracted signatures
				normalCheckDone <- &BundleResponse{
					BundleResponse: resp,
					Signatures:     pkg.BatchExtractSigFromTx(transactions),
				}
				return
			}
		}
		// If retries are exhausted, return nil
		normalCheckDone <- nil
	}()

	go func() {
		// Poll until confirmed
		for {
			// Check the signature status
			out, err := c.RPCConn.GetSignatureStatuses(
				ctx,
				false,
				*signature,
			)
			if err != nil {
				additionalCheckDone <- nil
				return
			}

			if out.Value != nil {
				reachBlockChain = true
				confirmed := false
				for _, status := range out.Value {
					if status != nil && status.ConfirmationStatus == "processed" && !reachBlockChain {
						reachBlockChain = true
					}

					if status != nil && status.ConfirmationStatus == "confirmed" {
						confirmed = true
						break
					}

					if status != nil && status.Err != nil {
						additionalCheckDone <- nil
						return
					}
				}

				if confirmed {
					additionalCheckDone <- &BundleResponse{
						BundleResponse: resp,
						Signatures:     pkg.BatchExtractSigFromTx(transactions),
					}
					return
				}
			} else if len(out.Value) == 0 || out.Value[0] == nil {
				additionalCheckDone <- nil
				return
			}

			// Wait for a short delay before retrying
			time.Sleep(CheckBundleRetryDelay)
		}
	}()

	// Wait for either normal check or additional check to complete
	select {
	case res := <-normalCheckDone:
		if res != nil {
			return res, nil
		}
	case res := <-additionalCheckDone:
		if res != nil {
			return res, nil
		}
		return nil, fmt.Errorf("tx (%s) didn't land", signature)
	case <-timeoutCtx.Done():
		// If the timeout is reached, return an error
		if !reachBlockChain {
			return nil, fmt.Errorf("BroadcastBundleWithConfirmation error: timeout exceeded")
		}
	}

	// If both checks fail, return an error
	return nil, fmt.Errorf("BroadcastBundleWithConfirmation error: max retries (%d) exceeded", CheckBundleRetries)
}

// SendBundle creates and sends a bundle of transactions to the Searcher service.
// It converts transactions to a protobuf packet and sends it using the SearcherService.
func (c *SearcherClient) SendBundle(
	transactions []*solana.Transaction,
	opts ...grpc.CallOption,
) (*searcher_pb.SendBundleResponse, error) {
	// Create a new bundle from the transactions
	bundle, err := c.NewBundle(transactions)
	if err != nil {
		return nil, err
	}

	// Send the bundle request to the Searcher service
	return c.SearcherService.SendBundle(
		c.AuthenticationService.GRPCCtx,
		&searcher_pb.SendBundleRequest{
			Bundle: bundle,
		},
		opts...,
	)
}

// NewBundle creates a new bundle protobuf object from a slice of transactions.
// It converts the transactions into protobuf packets and includes them in the bundle.
func (c *SearcherClient) NewBundle(transactions []*solana.Transaction) (*bundle_pb.Bundle, error) {
	// Convert the transactions to protobuf packets
	packets, err := pkg.ConvertBatchTransactionToProtobufPacket(transactions)
	if err != nil {
		return nil, err
	}

	// Create and return the bundle with the converted packets
	return &bundle_pb.Bundle{
		Packets: packets,
		Header:  nil,
	}, nil
}

// NewBundleSubscriptionResults subscribes to bundle result updates from the Searcher service.
// It uses the provided gRPC call options to set up the subscription.
func (c *SearcherClient) NewBundleSubscriptionResults(opts ...grpc.CallOption) (searcher_pb.SearcherService_SubscribeBundleResultsClient, error) {
	// Subscribe to bundle results from the Searcher service
	return c.SearcherService.SubscribeBundleResults(
		c.AuthenticationService.GRPCCtx,
		&searcher_pb.SubscribeBundleResultsRequest{},
		opts...,
	)
}
