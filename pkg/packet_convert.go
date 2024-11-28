package pkg

import (
	packet "github.com/Prophet-Solutions/jito-sdk/pb/packet"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

// ConvertProtobufPacketToTransaction converts a protobuf packet to a Solana transaction.
// It unmarshals the binary data in the packet into a Solana transaction.
func ConvertProtobufPacketToTransaction(packet *packet.Packet) (*solana.Transaction, error) {
	tx := &solana.Transaction{}
	err := tx.UnmarshalWithDecoder(bin.NewBorshDecoder(packet.Data))
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ConvertBatchProtobufPacketToTransaction converts a slice of protobuf packets to a slice of Solana transactions.
// It iterates over the provided packets, converts each one to a transaction, and accumulates the results in a slice.
func ConvertBatchProtobufPacketToTransaction(packets []*packet.Packet) ([]*solana.Transaction, error) {
	txs := make([]*solana.Transaction, 0, len(packets))

	for _, packet := range packets {
		tx, err := ConvertProtobufPacketToTransaction(packet)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}

	return txs, nil
}
