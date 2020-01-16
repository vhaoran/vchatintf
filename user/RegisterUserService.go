package intf

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
)

const (
	//todo
	RegUser_HANDLER_PATH = "/RegUser"
)

type (
	RegUserService interface {
		//todo
		Exec(in *RegUserIn) (*ykit.Result, error)
	}

	//input data
	//todo
	RegUserIn struct {
		UserCode   string `json:"user_code"`
		Mobile     string `json:"mobile omitempty"`
		Pwd        string `json:"pwd omitempty"`
		ChartKey   string `json:"chart_key omitempty"`
		ChartValue string `json:"chart_value omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	RegUserHandler struct {
		base ykit.RootTran
	}
)

func (r *RegUserHandler) MakeLocalEndpoint(svc RegUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  RegUser ###########")
		spew.Dump(ctx)

		//todo
		in := request.(*RegUserIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *RegUserHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(RegUserIn), ctx, req)
}

//个人实现,参数不能修改
func (r *RegUserHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *RegUserHandler) HandlerLocal(service RegUserService,
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
func (r *RegUserHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		//todo
		"POST",
		RegUser_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}
