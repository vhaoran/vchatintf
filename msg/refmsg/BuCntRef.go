package refmsg

import (
	"github.com/vhaoran/vchat/common/ytime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BuCntRef struct {
	ID        primitive.ObjectID `json:"id,omitempty"   bson:"_id,omitempty"`
	CreatedAt ytime.Date         `json:"created_at,omitempty"   bson:"created_at,omitempty"`
	Created   time.Time `json:"created,omitempty"   bson:"created,omitempty"`

	// 公众号id
	BID   int64  `json:"bid,omitempty"   bson:"bid,omitempty"`
	BName string `json:"b_name,omitempty"   bson:"b_name,omitempty"`

	//图标
	BIcon string `json:"b_icon,omitempty"   bson:"b_icon,omitempty"`

	// 页眉 可以是有格式内容
	Header interface{} `json:"header,omitempty"   bson:"header,omitempty"`

	//标题    可以是有格式内容
	Title interface{} `json:"title,omitempty"   bson:"title,omitempty"`
	//副标题	可以是有格式内容
	Title1 interface{} `json:"title_1,omitempty"   bson:"title_1,omitempty"`
	//  摘要(显示在推送区域)		可以是有格式内容
	Brief interface{} `json:"brief,omitempty"   bson:"brief,omitempty"`
	//  内容
	Content interface{} `json:"content,omitempty"   bson:"content,omitempty"`
	// 回复及点赞
	Reply interface{} `json:"reply,omitempty"   bson:"reply,omitempty"`
	// 页脚
	Footer interface{} `json:"footer,omitempty"   bson:"footer,omitempty"`
	//total push count
	TotalPushCount int64 `json:"total_push_count,omitempty"   bson:"total_push_count,omitempty"`

	//from
	From int64 `json:"from,omitempty"   bson:"from,omitempty"`
	FUserCode string `json:"f_user_code,omitempty"   bson:"f_user_code,omitempty"`
	FIcon string `json:"f_icon,omitempty"   bson:"f_icon,omitempty"`
	FNick string `json:"f_nick,omitempty"   bson:"f_nick,omitempty"`

}
