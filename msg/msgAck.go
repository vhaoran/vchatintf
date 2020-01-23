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
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/vhaoran/vchatintf/msg/refmsg"
)

const (
	//todo
	AckMsg_H_PATH = "/AckMsg"
)

type (
	MsgAckService interface {
		//todo
		Exec(in *AckMsgIn) (*ykit.Result, error)
	}

	AckMsgIDType struct {
		MsgType refmsg.MsgType     `json:"msg_type,omitempty"   bson:"msg_type,omitempty"`
		MsgID   primitive.ObjectID `json:"id,omitempty"   bson:"_id,omitempty"`
	}

	//input data
	//todo
	AckMsgIn struct {
		UID    int64          `json:"uid"`
		MsgIDs []AckMsgIDType `json:"msg_ids,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	AckMsgH struct {
		base ykit.RootTran
	}
)

func (r *AckMsgH) MakeLocalEndpoint(svc MsgAckService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  AckMsg ###########")
		spew.Dump(ctx)

		//todo
		in := request.(*AckMsgIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *AckMsgH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(AckMsgIn), ctx, req)
}

//个人实现,参数不能修改
func (r *AckMsgH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *AckMsgH) HandlerLocal(service MsgAckService,
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
func (r *AckMsgH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		//todo
		"POST",
		AckMsg_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}
