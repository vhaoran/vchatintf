package msg

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
)

const (
	TriggerOffLine_HANDLER_PATH = "/TriggerOffLine"
)

type (
	TriggerOffLineService interface {
		Exec(in *TriggerOffLineIn) (*ykit.Result, error)
	}

	//input data
	TriggerOffLineIn struct {
		UID int64 `json:"uid omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	TriggerOffLineHandler struct {
		base ykit.RootTran
	}
)

func (r *TriggerOffLineHandler) MakeLocalEndpoint(svc TriggerOffLineService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  TriggerOffLine ###########")
		spew.Dump(ctx)

		in := request.(*TriggerOffLineIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *TriggerOffLineHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(TriggerOffLineIn), ctx, req)
}

//个人实现,参数不能修改
func (r *TriggerOffLineHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *TriggerOffLineHandler) HandlerLocal(service TriggerOffLineService,
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
func (r *TriggerOffLineHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		TriggerOffLine_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *TriggerOffLineHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		TriggerOffLine_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_TriggerOffLine sync.Once
var local_TriggerOffLine_EP endpoint.Endpoint

func (r *TriggerOffLineHandler) Call(in TriggerOffLineIn) (*ykit.Result, error) {
	once_TriggerOffLine.Do(func() {
		local_TriggerOffLine_EP = new(TriggerOffLineHandler).ProxySD()
	})
	//
	ep := local_TriggerOffLine_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*ykit.Result), nil
}
