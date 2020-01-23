package inner

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
	"github.com/vhaoran/vchat/lib/ylog"
)

const (
	GetBuSubsInner_H_PATH = "/GetBulletinSubsInner"
)

type (
	GetBuSubsInnerService interface {
		Exec(in *GetBuSubsInnerIn) ([]*GetBuSubsInnerOut, error)
	}

	//input data
	GetBuSubsInnerIn struct {
		BID int64 `json:"bid,omitempty"`
	}

	//output data
	GetBuSubsInnerOut struct {
		UID int64 `json:"uid,omitempty"`
		//冗余字段，提升性能,用户帐号
		UserCodeRef string `json:"user_code_ref,omitempty"`
		//冗余字段，提升性能,眤称
		NickRef string `json:"nick_ref,omitempty"`
		//冗余字段，提升性能,图标
		IconRef string `json:"icon_ref,omitempty"`
	}

	// handler implements
	GetBuSubsInnerH struct {
		base ykit.RootTran
	}
)

func (r *GetBuSubsInnerH) MakeLocalEndpoint(svc GetBuSubsInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ylog.Debug("#############  GetBulletinSubsInner ###########")
		spew.Dump(ctx)

		in := request.(*GetBuSubsInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetBuSubsInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetBuSubsInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetBuSubsInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response []*GetBuSubsInnerOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetBuSubsInnerH) HandlerLocal(service GetBuSubsInnerService,
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
func (r *GetBuSubsInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuSubsInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetBuSubsInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuSubsInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_GetBulletinSubsInner sync.Once
var local_GetBulletinSubsInner_EP endpoint.Endpoint

func (r *GetBuSubsInnerH) Call(in *GetBuSubsInnerIn) ([]*GetBuSubsInnerOut, error) {
	once_GetBulletinSubsInner.Do(func() {
		local_GetBulletinSubsInner_EP = new(GetBuSubsInnerH).ProxySD()
	})
	//
	ep := local_GetBulletinSubsInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.([]*GetBuSubsInnerOut), nil
}
