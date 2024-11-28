package pkg

import (
	packet "github.com/Prophet-Solutions/jito-sdk/pb/packet"
	"github.com/gagliardetto/solana-go"
)

// ConvertTransactionToProtobufPacket converts a Solana transaction to a protobuf packet.
// It marshals the transaction to binary data and creates a jito_pb.Packet with this data and associated metadata.
func ConvertTransactionToProtobufPacket(transaction *solana.Transaction) (packet.Packet, error) {
	data, err := transaction.MarshalBinary()
	if err != nil {
		return packet.Packet{}, err
	}

	return packet.Packet{
		Data: data,
		Meta: &packet.Meta{
			Size:        uint64(len(data)),
			Addr:        "", // Address not provided, set to empty string
			Port:        0,  // Port not provided, set to 0
			Flags:       nil,
			SenderStake: 0, // SenderStake not provided, set to 0
		},
	}, nil
}

// ConvertBatchTransactionToProtobufPacket converts a batch of Solana transactions to a slice of protobuf packets.
// It iterates over each transaction, converts it to a protobuf packet, and accumulates the results in a slice.
func ConvertBatchTransactionToProtobufPacket(transactions []*solana.Transaction) ([]*packet.Packet, error) {
	packets := make([]*packet.Packet, 0, len(transactions))

	for _, tx := range transactions {
		packet, err := ConvertTransactionToProtobufPacket(tx)
		if err != nil {
			return nil, err
		}

		packets = append(packets, &packet)
	}

	return packets, nil
}
