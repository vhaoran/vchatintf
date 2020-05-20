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
	IsGroupMemberOfRedInner_H_PATH = "/IsGroupMemberOfRedInner"
)

type (
	IsGroupMemberOfRedInnerService interface {
		Exec(in *IsGroupMemberOfRedInnerIn) (*IsGroupMemberOfRedInnerOut, error)
	}

	//input data
	IsGroupMemberOfRedInnerIn struct {
		GID int64 `json:"gid,omitempty"`
		UID int64 `json:"uid,omitempty"`
	}

	//output data
	IsGroupMemberOfRedInnerOut struct {
		OK bool `json:"ok"`
	}

	// handler implements
	IsGroupMemberOfRedInnerH struct {
		base ykit.RootTran
	}
)

func (r *IsGroupMemberOfRedInnerH) MakeLocalEndpoint(svc IsGroupMemberOfRedInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  isGroupMemberOfRedInner ###########")
		spew.Dump(ctx)

		in := request.(*IsGroupMemberOfRedInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *IsGroupMemberOfRedInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(IsGroupMemberOfRedInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *IsGroupMemberOfRedInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *IsGroupMemberOfRedInnerOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *IsGroupMemberOfRedInnerH) HandlerLocal(service IsGroupMemberOfRedInnerService,
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
func (r *IsGroupMemberOfRedInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		IsGroupMemberOfRedInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *IsGroupMemberOfRedInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		IsGroupMemberOfRedInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_isGroupMemberOfRedInner sync.Once
var local_isGroupMemberOfRedInner_EP endpoint.Endpoint

func (r *IsGroupMemberOfRedInnerH) Call(in *IsGroupMemberOfRedInnerIn) (bool, error) {
	once_isGroupMemberOfRedInner.Do(func() {
		local_isGroupMemberOfRedInner_EP = new(IsGroupMemberOfRedInnerH).ProxySD()
	})
	//
	ep := local_isGroupMemberOfRedInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return false, err
	}

	return result.(*IsGroupMemberOfRedInnerOut).OK, nil
}
