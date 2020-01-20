package msg

//for snippet用于标准返回值的微服务接口

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/common/ytime"
	"github.com/vhaoran/vchat/lib/ykit"
	"net/http"
)

const (
	TriggerOnLine_HANDLER_PATH = "/TriggerOnLine"
)

type (
	TriggerOnLineService interface {
		Exec(in *TriggerOnLineIn) (*ykit.Result, error)
	}

	//input data
	TriggerOnLineIn struct {
		UID         int64      `json:"uid omitempty"`
		LastAckTime ytime.Date `json:"last_ack_time omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	TriggerOnLineHandler struct {
		base ykit.RootTran
	}
)

func (r *TriggerOnLineHandler) MakeLocalEndpoint(svc TriggerOnLineService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  TriggerOnLine ###########")
		spew.Dump(ctx)

		in := request.(*TriggerOnLineIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *TriggerOnLineHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(TriggerOnLineIn), ctx, req)
}

//个人实现,参数不能修改
func (r *TriggerOnLineHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *TriggerOnLineHandler) HandlerLocal(service TriggerOnLineService,
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
func (r *TriggerOnLineHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		TriggerOnLine_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *TriggerOnLineHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		TriggerOnLine_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
