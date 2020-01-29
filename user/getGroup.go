package intf

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
	GetGroup_H_PATH = "/GetGroup"
)

type (
	GetGroupService interface {
		Exec(ctx context.Context, in *GetGroupIn) (*ykit.Result, error)
	}

	//input data 
	GetGroupIn struct {
		GID int64 `json:",omitempty"`
	}

	//output data

	// handler implements
	GetGroupH struct {
		base ykit.RootTran
	}
)

func (r *GetGroupH) MakeLocalEndpoint(svc GetGroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetGroup ###########")
		spew.Dump(ctx)

		in := request.(*GetGroupIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *GetGroupH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetGroupIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetGroupH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetGroupH) HandlerLocal(service GetGroupService,
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
func (r *GetGroupH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetGroup_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetGroupH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetGroup_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
