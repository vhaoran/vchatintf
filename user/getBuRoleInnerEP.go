package user

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
	GetBuRoleInner_H_PATH = "/GetBuRoleInner"
)

type (
	GetBuRoleInnerService interface {
		Exec(in *GetBuRoleInnerIn) (*GetBuRoleInnerOut, error)
	}

	//input data
	GetBuRoleInnerIn struct {
		BID int64 `json:"bid"`
		UID int64 `json:"uid"`
	}

	//output data
	GetBuRoleInnerOut struct {
		IsManager bool `json:"is_manager"`
		IsOwner   bool `json:"is_owner"`
	}

	// handler implements
	GetBuRoleInnerH struct {
		base ykit.RootTran
	}
)

func (r *GetBuRoleInnerH) MakeLocalEndpoint(svc GetBuRoleInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetBuRoleInner ###########")
		spew.Dump(ctx)

		in := request.(*GetBuRoleInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetBuRoleInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetBuRoleInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetBuRoleInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *GetBuRoleInnerOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetBuRoleInnerH) HandlerLocal(service GetBuRoleInnerService,
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
func (r *GetBuRoleInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuRoleInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetBuRoleInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuRoleInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_GetBuRoleInner sync.Once
var local_GetBuRoleInner_EP endpoint.Endpoint

func (r *GetBuRoleInnerH) Call(in *GetBuRoleInnerIn) (*GetBuRoleInnerOut, error) {
	once_GetBuRoleInner.Do(func() {
		local_GetBuRoleInner_EP = new(GetBuRoleInnerH).ProxySD()
	})
	//
	ep := local_GetBuRoleInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*GetBuRoleInnerOut), nil
}
