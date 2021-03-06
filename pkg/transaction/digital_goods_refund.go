package transaction

import (
	pb "github.com/ac0v/aspera/pkg/api/p2p"
	"github.com/ac0v/aspera/pkg/encoding"
)

const (
	DigitalGoodsRefundType    = 3
	DigitalGoodsRefundSubType = 7
)

type DigitalGoodsRefund struct {
	*pb.DigitalGoodsRefund
}

func (tx *DigitalGoodsRefund) WriteAttachmentBytes(e encoding.Encoder) {
	e.WriteUint64(tx.Attachment.Purchase)
	e.WriteUint64(tx.Attachment.Refund)
}

func (tx *DigitalGoodsRefund) AttachmentSizeInBytes() int {
	return 8 + 8
}

func (tx *DigitalGoodsRefund) GetType() uint16 {
	return DigitalGoodsRefundSubType<<8 | DigitalGoodsRefundType
}
