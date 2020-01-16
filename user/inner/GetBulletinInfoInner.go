package inner

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

	"github.com/vhaoran/vchatintf/user/refuser"
)

const (
	GetBulletinInfoInner_HANDLER_PATH = "/GetBulletinInfoRef"
)

type (
	GetBulletinInfoInnerService interface {
		Exec(in *GetBulletinInfoInnerIn) (*refuser.BulletinInfoRef, error)
	}

	//input data
	GetBulletinInfoInnerIn struct {
		BID int64 `json:"bid,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	GetBulletinInfoInnerHandler struct {
		base ykit.RootTran
	}
)

func (r *GetBulletinInfoInnerHandler) MakeLocalEndpoint(svc GetBulletinInfoInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetBulletinInfoRef ###########")
		spew.Dump(ctx)

		in := request.(*GetBulletinInfoInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetBulletinInfoInnerHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetBulletinInfoInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetBulletinInfoInnerHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *refuser.BulletinInfoRef
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetBulletinInfoInnerHandler) HandlerLocal(service GetBulletinInfoInnerService,
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
func (r *GetBulletinInfoInnerHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBulletinInfoInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

//func (r *GetBulletinInfoInnerHandler) ProxySD() endpoint.Endpoint {
//	return r.base.ProxyEndpointSD(
//		context.Background(),
//		MSTAG,
//		"POST",
//		GetBulletinInfoInner_HANDLER_PATH,
//		r.DecodeRequest,
//		r.DecodeResponse)
//}

func (r *GetBulletinInfoInnerHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBulletinInfoInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
