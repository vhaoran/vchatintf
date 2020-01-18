package inner

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
	"github.com/vhaoran/vchat/lib/ylog"
)

const (
	GetBulletinSubsInner_HANDLER_PATH = "/GetBulletinSubsInner"
)

type (
	GetBulletinSubsInnerService interface {
		Exec(in *GetBulletinSubsInnerIn) ([]*GetBulletinSubsInnerOut, error)
	}

	//input data
	GetBulletinSubsInnerIn struct {
		BID int64 `json:"bid,omitempty"`
	}

	//output data
	GetBulletinSubsInnerOut struct {
		UID int64 `json:"uid,omitempty"`
		//冗余字段，提升性能,用户帐号
		UserCodeRef string `json:"user_code_ref,omitempty"`
		//冗余字段，提升性能,眤称
		NickRef string `json:"nick_ref,omitempty"`
		//冗余字段，提升性能,图标
		IconRef string `json:"icon_ref,omitempty"`
	}

	// handler implements
	GetBulletinSubsInnerHandler struct {
		base ykit.RootTran
	}
)

func (r *GetBulletinSubsInnerHandler) MakeLocalEndpoint(svc GetBulletinSubsInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ylog.Debug("#############  GetBulletinSubsInner ###########")
		spew.Dump(ctx)

		in := request.(*GetBulletinSubsInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetBulletinSubsInnerHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetBulletinSubsInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetBulletinSubsInnerHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response []*GetBulletinSubsInnerOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetBulletinSubsInnerHandler) HandlerLocal(service GetBulletinSubsInnerService,
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
func (r *GetBulletinSubsInnerHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBulletinSubsInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetBulletinSubsInnerHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBulletinSubsInner_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
