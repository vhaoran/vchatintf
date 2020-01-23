package msg

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"

	"github.com/vhaoran/vchatintf/msg/refmsg"
)

const (
	//todo
	MsgSend_H_PATH = "/SendMsg"
)

type (
	MsgSendService interface {
		//todo
		Exec(in *SendMsgIn) (*ykit.Result, error)
	}

	//input data
	//todo
	SendMsgIn struct {
		//发送后，是否显示传入的内容。
		SendBack bool           `json:"send_back,omitempty"`
		MsgType  refmsg.MsgType `json:"msg_type,omitempty"   bson:"msg_type,omitempty"`

		//发送方id
		FromUID int64 `json:"from_uid,omitempty"   bson:"from_uid,omitempty"`

		//目标用户ID
		ToUID int64 `json:"to_uid,omitempty"   bson:"to_uid,omitempty"`

		ToGID int64 `json:"to_gid,omitempty"   bson:"to_gid,omitempty"`
		//自定义的消息内容类型，可选
		BodyType int `json:"body_type,omitempty"   bson:"body_type,omitempty"`

		//消息体
		Body interface{} `json:"body,omitempty"   bson:"body,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	SendMsgH struct {
		base ykit.RootTran
	}
)

func (r *SendMsgH) MakeLocalEndpoint(svc MsgSendService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  SendMsg ###########")
		spew.Dump(ctx)

		//todo
		in := request.(*SendMsgIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *SendMsgH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(SendMsgIn), ctx, req)
}

//个人实现,参数不能修改
func (r *SendMsgH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *SendMsgH) HandlerLocal(service MsgSendService,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	ep := r.MakeLocalEndpoint(service)
	for _, f := range mid {
		ep = f(ep)
	}

	handler := tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
		options...)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *SendMsgH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		//todo
		"POST",
		MsgSend_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}
