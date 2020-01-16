package refmsg

import (
	"github.com/vhaoran/vchat/common/ytime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CircleRef struct {
	ID primitive.ObjectID `json:"id,omitempty"   bson:"_id,omitempty"`
	// 发布时间
	CreatedAt ytime.Date `json:"created_at,omitempty"   bson:"created_at,omitempty"`

	UID         int64  `json:"uid,omitempty"   bson:"uid,omitempty"`
	UserCodeRef string `json:"user_code_ref,omitempty"   bson:"user_code_ref,omitempty"`
	NickRef     string `json:"nick_ref,omitempty"   bson:"nick_ref,omitempty"`
	IconRef     string `json:"icon_ref,omitempty"   bson:"icon_ref,omitempty"`

	Content interface{} `json:"content,omitempty"   bson:"content,omitempty"`
	//回复及点赞
	Reply interface{} `json:"reply,omitempty"   bson:"reply,omitempty"`
}
