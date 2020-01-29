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
	FriendReq_H_PATH = "/FriendReq"
)

type (
	FriendReqService interface {
		Exec(ctx context.Context, in *FriendReqIn) (*ykit.Result, error)
	}

	//input data
	FriendReqIn struct {
		From int64  `json:"from,omitempty"`
		To   int64  `json:"to,omitempty"`
		Msg  string `json:"msg,omitempty"`
	}

	//output data
	// handler implements
	FriendReqH struct {
		base ykit.RootTran
	}
)

func (r *FriendReqH) MakeLocalEndpoint(svc FriendReqService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  FriendReq ###########")
		spew.Dump(ctx)

		in := request.(*FriendReqIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *FriendReqH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(FriendReqIn), ctx, req)
}

//个人实现,参数不能修改
func (r *FriendReqH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *FriendReqH) HandlerLocal(service FriendReqService,
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
func (r *FriendReqH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		FriendReq_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *FriendReqH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		FriendReq_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
