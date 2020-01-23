package inner

//for snippet用于标准返回值的微服务接口

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

	"github.com/vhaoran/vchatintf/user"
	"github.com/vhaoran/vchatintf/user/refuser"
)

const (
	//
	GetUserInfoInner_H_PATH = "/GetUserInfoInner"
)

type (
	GetUserInfoInnerService interface {
		Exec(in *GetUserInfoInnerIn) (*refuser.UserInfoRef, error)
	}

	//input data
	//
	GetUserInfoInnerIn struct {
		UID int64 `json:"uid,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	GetUserInfoInnerH struct {
		base ykit.RootTran
	}
)

func (r *GetUserInfoInnerH) MakeLocalEndpoint(svc GetUserInfoInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetUserInfoInner ###########")
		spew.Dump(ctx)

		//
		in := request.(*GetUserInfoInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetUserInfoInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetUserInfoInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetUserInfoInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *refuser.UserInfoRef
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetUserInfoInnerH) HandlerLocal(service GetUserInfoInnerService,
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
func (r *GetUserInfoInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		intf.MSTAG,
		//
		"POST",
		GetUserInfoInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetUserInfoInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetUserInfoInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_GetUserInfoInner sync.Once
var local_GetUserInfoInner_EP endpoint.Endpoint

func (r *GetUserInfoInnerH) Call(in GetUserInfoInnerIn) (*refuser.UserInfoRef, error) {
	once_GetUserInfoInner.Do(func() {
		local_GetUserInfoInner_EP = new(GetUserInfoInnerH).ProxySD()
	})
	//
	ep := local_GetUserInfoInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*refuser.UserInfoRef), nil
}
