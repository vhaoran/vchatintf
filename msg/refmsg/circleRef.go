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
	UserCodeRef string `json:"user_code ,omitempty"   bson:"user_code ,omitempty"`
	NickRef     string `json:"nick ,omitempty"   bson:"nick ,omitempty"`
	IconRef     string `json:"icon ,omitempty"   bson:"icon ,omitempty"`

	Content interface{} `json:"content,omitempty"   bson:"content,omitempty"`
	//回复及点赞
	Reply interface{} `json:"reply,omitempty"   bson:"reply,omitempty"`
}