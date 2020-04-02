package refmsg

import (
	"github.com/vhaoran/vchat/common/ytime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BuCntRef struct {
	ID        primitive.ObjectID `json:"id"   bson:"_id"`
	CreatedAt ytime.Date         `json:"created_at"   bson:"created_at"`
	Created   time.Time `json:"created"   bson:"created"`

	// 公众号id
	BID   int64  `json:"bid"   bson:"bid"`
	BName string `json:"b_name"   bson:"b_name"`

	//图标
	BIcon string `json:"b_icon"   bson:"b_icon"`

	// 页眉 可以是有格式内容
	Header interface{} `json:"header"   bson:"header"`

	//标题    可以是有格式内容
	Title interface{} `json:"title"   bson:"title"`
	//副标题	可以是有格式内容
	Title1 interface{} `json:"title_1"   bson:"title_1"`
	//  摘要(显示在推送区域)		可以是有格式内容
	Brief interface{} `json:"brief"   bson:"brief"`
	//  内容
	Content interface{} `json:"content"   bson:"content"`
	// 回复及点赞
	Reply interface{} `json:"reply"   bson:"reply"`
	// 页脚
	Footer interface{} `json:"footer"   bson:"footer"`
	//total push count
	TotalPushCount int64 `json:"total_push_count"   bson:"total_push_count"`

	//from
	From int64 `json:"from"   bson:"from"`
	FUserCode string `json:"f_user_code"   bson:"f_user_code"`
	FIcon string `json:"f_icon"   bson:"f_icon"`
	FNick string `json:"f_nick"   bson:"f_nick"`

}
