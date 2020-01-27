package intf

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vhaoran/vchat/common/ypage"
	"net/http"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
)

const (
	AddrPage_H_PATH = "/AddrPage"
)

type (
	AddrPageService interface {
		Exec(ctx context.Context, in *AddrPageIn) (*AddrPageOut, error)
	}

	//input data
	AddrPageIn struct {
		ypage.PageBean
	}

	//output data
	AddrPageOut struct {
		ypage.PageResult
	}

	// handler implements
	AddrPageH struct {
		base ykit.RootTran
	}
)

func (r *AddrPageH) MakeLocalEndpoint(svc AddrPageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  AddrPage ###########")
		spew.Dump(ctx)

		in := request.(*AddrPageIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *AddrPageH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(AddrPageIn), ctx, req)
}

//个人实现,参数不能修改
func (r *AddrPageH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *AddrPageOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *AddrPageH) HandlerLocal(service AddrPageService,
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
func (r *AddrPageH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		AddrPage_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *AddrPageH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		AddrPage_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_AddrPage sync.Once
var local_AddrPage_EP endpoint.Endpoint

func (r *AddrPageH) Call(in AddrPageIn) (*AddrPageOut, error) {
	once_AddrPage.Do(func() {
		local_AddrPage_EP = new(AddrPageH).ProxySD()
	})
	//
	ep := local_AddrPage_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*AddrPageOut), nil
}
