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
	IsGroupSuperUserInner_H_PATH = "/IsGroupSuperUserInner"
)

type (
	IsGroupSuperUserInnerService interface {
		Exec(in *IsGroupSuperUserInnerIn) (*IsGroupSuperUserInnerOut, error)
	}

	//input data
	IsGroupSuperUserInnerIn struct {
		GID int64 `json:"gid"`
		UID int64 `json:"uid"`
	}

	//output data
	IsGroupSuperUserInnerOut struct {
		IsSuperUser bool
	}

	// handler implements
	IsGroupSuperUserInnerH struct {
		base ykit.RootTran
	}
)

func (r *IsGroupSuperUserInnerH) MakeLocalEndpoint(svc IsGroupSuperUserInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  IsGroupSuperUser ###########")
		spew.Dump(ctx)

		in := request.(*IsGroupSuperUserInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *IsGroupSuperUserInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(IsGroupSuperUserInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *IsGroupSuperUserInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *IsGroupSuperUserInnerOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *IsGroupSuperUserInnerH) HandlerLocal(service IsGroupSuperUserInnerService,
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
func (r *IsGroupSuperUserInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		IsGroupSuperUserInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *IsGroupSuperUserInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		IsGroupSuperUserInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_IsGroupSuperUser sync.Once
var local_IsGroupSuperUser_EP endpoint.Endpoint

func (r *IsGroupSuperUserInnerH) Call(in *IsGroupSuperUserInnerIn) (bool, error) {
	once_IsGroupSuperUser.Do(func() {
		local_IsGroupSuperUser_EP = new(IsGroupSuperUserInnerH).ProxySD()
	})
	//
	ep := local_IsGroupSuperUser_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return false, nil
	}

	return result.(*IsGroupSuperUserInnerOut).IsSuperUser, nil
}
