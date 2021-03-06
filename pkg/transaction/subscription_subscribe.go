package transaction

import (
	pb "github.com/ac0v/aspera/pkg/api/p2p"
	"github.com/ac0v/aspera/pkg/encoding"
)

const (
	SubscriptionSubscribeType    = 21
	SubscriptionSubscribeSubType = 3
)

type SubscriptionSubscribe struct {
	*pb.SubscriptionSubscribe
}

func (tx *SubscriptionSubscribe) WriteAttachmentBytes(e encoding.Encoder) {
	e.WriteUint32(tx.Attachment.Frequency)
}
func (tx *SubscriptionSubscribe) AttachmentSizeInBytes() int {
	return 4
}

func (tx *SubscriptionSubscribe) GetType() uint16 {
	return SubscriptionSubscribeSubType<<8 | SubscriptionSubscribeType
}
