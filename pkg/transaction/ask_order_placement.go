package transaction

import (
	pb "github.com/ac0v/aspera/pkg/api/p2p"
	"github.com/ac0v/aspera/pkg/encoding"
)

const (
	AskOrderPlacementType    = 2
	AskOrderPlacementSubType = 2
)

type AskOrderPlacement struct {
	*pb.AskOrderPlacement
}

func (tx *AskOrderPlacement) WriteAttachmentBytes(e encoding.Encoder) {
	e.WriteUint64(tx.Attachment.Asset)
	e.WriteUint64(tx.Attachment.Quantity)
	e.WriteUint64(tx.Attachment.Price)
}

func (tx *AskOrderPlacement) AttachmentSizeInBytes() int {
	return 8 + 8 + 8
}

func (tx *AskOrderPlacement) GetType() uint16 {
	return AskOrderPlacementSubType<<8 | AskOrderPlacementType
}
