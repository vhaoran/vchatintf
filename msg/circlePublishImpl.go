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
)

const (
	CirclePublish_H_PATH = "/CirclePublish"
)

type (
	CirclePublishService interface {
		Exec(in *CirclePublishIn) (*ykit.Result, error)
	}

	//input data
	CirclePublishIn struct {
		UID  int64  `json:"uid,omitempty"`
		Text string `json:"text"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	CirclePublishH struct {
		base ykit.RootTran
	}
)

func (r *CirclePublishH) MakeLocalEndpoint(svc CirclePublishService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  CirclePublish ###########")
		spew.Dump(ctx)

		in := request.(*CirclePublishIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *CirclePublishH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(CirclePublishIn), ctx, req)
}

//个人实现,参数不能修改
func (r *CirclePublishH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *CirclePublishH) HandlerLocal(service CirclePublishService,
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
func (r *CirclePublishH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		CirclePublish_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}
