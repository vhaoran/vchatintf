package msg

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

	"github.com/vhaoran/vchatintf/msg/refmsg"
)

const (
	//todo
	BulletinPublish_HANDLER_PATH = "/BulletinPublish"
)

type (
	BulletinPublishService interface {
		//todo
		Exec(in *BulletinPublishRequest) (*ykit.Result, error)
	}

	//input data
	//todo
	BulletinPublishRequest struct {
		refmsg.BulletinContentRef
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	BulletinPublishHandler struct {
		base ykit.RootTran
	}
)

func (r *BulletinPublishHandler) MakeLocalEndpoint(svc BulletinPublishService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  BulletinPublish ###########")
		spew.Dump(ctx)

		//todo
		in := request.(*BulletinPublishRequest)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *BulletinPublishHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(BulletinPublishRequest), ctx, req)
}

//个人实现,参数不能修改
func (r *BulletinPublishHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *BulletinPublishHandler) HandlerLocal(service BulletinPublishService,
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
func (r *BulletinPublishHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		//todo
		"POST",
		BulletinPublish_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}