package transaction

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderFromBytes(t *testing.T) {
	for _, parseTest := range ParseTransactionTests {
		bs, _ := hex.DecodeString(parseTest.txByteStr)
		header, err := headerFromBytes(bs)
		if assert.Nil(t, err) {
			assert.Equal(t, parseTest.header.Type, header.Type)
			assert.Equal(t, parseTest.header.GetSubtype(), header.GetSubtype())
			assert.Equal(t, parseTest.header.Timestamp, header.Timestamp)
			assert.Equal(t, parseTest.header.Deadline, header.Deadline)
			assert.Equal(t, parseTest.header.SenderPublicKey, header.SenderPublicKey)
			assert.Equal(t, parseTest.header.RecipientID, header.RecipientID)
			assert.Equal(t, parseTest.header.AmountNQT, header.AmountNQT)
			assert.Equal(t, parseTest.header.FeeNQT, header.FeeNQT)
			assert.Equal(t, parseTest.header.ReferencedTransactionFullHash,
				header.ReferencedTransactionFullHash)
			assert.Equal(t, parseTest.header.Signature, header.Signature)
			assert.Equal(t, parseTest.header.GetVersion(), header.GetVersion())
		}
	}
}

func BenchmarkHeaderFromBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs, _ := hex.DecodeString(ParseTransactionTests[i%len(ParseTransactionTests)].txByteStr)
		_, _ = headerFromBytes(bs)
	}
}

func TestTransactionFromAndToBytes(t *testing.T) {
	for _, parseTest := range ParseTransactionTests {
		bsExp, _ := hex.DecodeString(parseTest.txByteStr)
		tx, err := FromBytes(bsExp)
		if assert.Nil(t, err) && assert.NotNil(t, tx) {
			bs, err := tx.ToBytes()
			if assert.Nil(t, err) {
				assert.Equal(t, bsExp, bs)
			}
		}
	}
}

func TestVerifySignature(t *testing.T) {
	for _, parseTest := range ParseTransactionTests {
		// TODO: escrow result does not need to be check
		if parseTest.header.Type == 21 && parseTest.header.SubtypeAndVersion == 18 {
			continue
		}
		bsExp, _ := hex.DecodeString(parseTest.txByteStr)
		tx, _ := FromBytes(bsExp)
		assert.Nil(t, tx.VerifySignature())
	}
}
