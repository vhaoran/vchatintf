package refmsg

import (
	"time"

	"github.com/vhaoran/vchat/common/ytime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CircleRef struct {
	ID primitive.ObjectID `json:"id,omitempty"   bson:"_id,omitempty"`

	// 发布时间
	CreatedAt ytime.Date `json:"created_at"   bson:"created_at"`
	Created   time.Time  `json:"created"   bson:"created"`

	UID         int64  `json:"uid"   bson:"uid"`
	UserCodeRef string `json:"user_code "   bson:"user_code "`
	NickRef     string `json:"nick "   bson:"nick "`
	IconRef     string `json:"icon "   bson:"icon "`

	Content interface{} `json:"content"   bson:"content"`
	//回复及点赞
	//	Prize   []CirclePrizeRef   `json:"prize"   bson:"prize"`
	Comment []CircleCommentRef `json:"comment"   bson:"comment"`
}

//type CirclePrizeRef struct {
//	CreatedAt ytime.Date `json:"created_at"   bson:"created_at"`
//
//	OfUID int64 `json:"Of_uid"   bson:"Of_uid"`
//
//	UserCodeRef string `json:"user_code "   bson:"user_code "`
//	NickRef     string `json:"nick "   bson:"nick "`
//	IconRef     string `json:"icon "   bson:"icon "`
//}

type CircleCommentRef struct {
	//0 : prize 1: comment 2:reply

	Action    int        `json:"action"   bson:"action"`
	CreatedAt ytime.Date `json:"created_at"   bson:"created_at"`

	//评论人
	OfUID int64 `json:"Of_uid"   bson:"Of_uid"`

	UserCodeRef string `json:"user_code"   bson:"user_code "`
	NickRef     string `json:"nick"   bson:"nick "`
	IconRef     string `json:"icon"   bson:"icon "`

	Text string `json:"content"   bson:"content"`
}
