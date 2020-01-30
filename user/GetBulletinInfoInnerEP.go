package user

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

	"github.com/vhaoran/vchatintf/user/refuser"
)

const (
	GetBuInfoInner_H_PATH = "/GetBulletinInfoInner"
)

type (
	GetBuInfoInnerService interface {
		Exec(in *GetBuInfoInnerIn) (*refuser.BulletinInfoRef, error)
	}

	//input data
	GetBuInfoInnerIn struct {
		BID int64 `json:"bid,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	GetBuInfoInnerH struct {
		base ykit.RootTran
	}
)

func (r *GetBuInfoInnerH) MakeLocalEndpoint(svc GetBuInfoInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetBulletinInfoRef ###########")
		spew.Dump(ctx)

		in := request.(*GetBuInfoInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetBuInfoInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetBuInfoInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetBuInfoInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *refuser.BulletinInfoRef
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetBuInfoInnerH) HandlerLocal(service GetBuInfoInnerService,
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
func (r *GetBuInfoInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuInfoInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

//func (r *GetBuInfoInnerH) ProxySD() endpoint.Endpoint {
//	return r.base.ProxyEndpointSD(
//		context.Background(),
//		MSTAG,
//		"POST",
//		GetBuInfoInner_H_PATH,
//		r.DecodeRequest,
//		r.DecodeResponse)
//}

func (r *GetBuInfoInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuInfoInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_GetBulletinInfoInner sync.Once
var local_GetBulletinInfoInner_EP endpoint.Endpoint

func (r *GetBuInfoInnerH) Call(in *GetBuInfoInnerIn) (*refuser.BulletinInfoRef, error) {
	once_GetBulletinInfoInner.Do(func() {
		local_GetBulletinInfoInner_EP = new(GetBuInfoInnerH).ProxySD()
	})
	//
	ep := local_GetBulletinInfoInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*refuser.BulletinInfoRef), nil
}
