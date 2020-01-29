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
	SendUserCard_H_PATH = "/SendUserCard"
)

type (
	SendUserCardService interface {
		Exec(ctx context.Context, in *SendUserCardIn) (*ykit.Result, error)
	}

	//input data
	SendUserCardIn struct {
		From int64 `json:"from,omitempty"`
		To   int64 `json:"to,omitempty"`
	}

	//output data

	// handler implements
	SendUserCardH struct {
		base ykit.RootTran
	}
)

func (r *SendUserCardH) MakeLocalEndpoint(svc SendUserCardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  SendUserCard ###########")
		spew.Dump(ctx)

		in := request.(*SendUserCardIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *SendUserCardH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(SendUserCardIn), ctx, req)
}

//个人实现,参数不能修改
func (r *SendUserCardH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *SendUserCardH) HandlerLocal(service SendUserCardService,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	ep := r.MakeLocalEndpoint(service)
	for _, f := range mid {
		ep = f(ep)
	}

	before := tran.ServerBefore(ykit.Jwt2ctx())
	opts := make([]tran.ServerOption, 0)
	opts = append(opts, before)
	opts = append(opts, options...)

	handler := tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
		opts...)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *SendUserCardH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		SendUserCard_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *SendUserCardH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		SendUserCard_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
