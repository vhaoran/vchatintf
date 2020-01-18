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
	GetGroupInfoInner_HANDLER_PATH = "/GetGroupInfoInner"
)

type (
	GetGroupInfoInnerService interface {
		//
		Exec(in *GetGroupInfoInnerIn) (*refuser.GroupInfoRef, error)
	}

	//input data
	//
	GetGroupInfoInnerIn struct {
		GID int64 `json:"gid,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	GetGroupInfoInnerHandler struct {
		base ykit.RootTran
	}
)

func (r *GetGroupInfoInnerHandler) MakeLocalEndpoint(svc GetGroupInfoInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetGroupInfoInner ###########")
		spew.Dump(ctx)

		//
		in := request.(*GetGroupInfoInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetGroupInfoInnerHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetGroupInfoInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetGroupInfoInnerHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *refuser.GroupInfoRef
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetGroupInfoInnerHandler) HandlerLocal(service GetGroupInfoInnerService,
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
func (r *GetGroupInfoInnerHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		intf.MSTAG,
		//
		"POST",
		GetGroupInfoInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetGroupInfoInnerHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetGroupInfoInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部的其它微服务调用 ，不从网关层调用
var once_GetGroupInfoInner sync.Once
var local_GetGroupInfoInner_EP endpoint.Endpoint

//
func (r *GetGroupInfoInnerHandler) Call(in GetGroupInfoInnerIn) (*refuser.GroupInfoRef, error) {
	once_GetGroupInfoInner.Do(func() {
		local_GetGroupInfoInner_EP = new(GetGroupInfoInnerHandler).ProxySD()
	})
	//
	ep := local_GetGroupInfoInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*refuser.GroupInfoRef), nil
}
