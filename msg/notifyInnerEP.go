package msg

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/vhaoran/vchatintf/msg/refmsg"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
)

const (
	NotifyMsgInner_HANDLER_PATH = "/NotifyMsgInner"
)

type (
	NotifyMsgInnerService interface {
		Exec(in *NotifyMsgInnerIn) (*ykit.Result, error)
	}

	//input data
	NotifyMsgInnerIn struct {
		refmsg.MsgHisRef
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	NotifyMsgInnerH struct {
		base ykit.RootTran
	}
)

func (r *NotifyMsgInnerH) MakeLocalEndpoint(svc NotifyMsgInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  NotifyMsgInner ###########")
		spew.Dump(ctx)

		in := request.(*NotifyMsgInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *NotifyMsgInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(NotifyMsgInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *NotifyMsgInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *NotifyMsgInnerH) HandlerLocal(service NotifyMsgInnerService,
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
func (r *NotifyMsgInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		NotifyMsgInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *NotifyMsgInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		NotifyMsgInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从gateway调用
var once_NotifyMsgInner sync.Once
var local_NotifyMsgInner_EP endpoint.Endpoint

func (r *NotifyMsgInnerH) Call(in *NotifyMsgInnerIn) (*ykit.Result, error) {
	once_NotifyMsgInner.Do(func() {
		local_NotifyMsgInner_EP = new(NotifyMsgInnerH).ProxySD()
	})
	//
	ep := local_NotifyMsgInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return ykit.RErr(err), nil
	}

	return result.(*ykit.Result), nil
}
