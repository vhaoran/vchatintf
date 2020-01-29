package msg

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
)

const (
	//todo
	CirclePrize_H_PATH = "/CirclePrize"
)

type (
	CirclePrizeService interface {
		//todo
		Exec(in *CirclePrizeIn) (*ykit.Result, error)
	}

	//input data
	//todo
	CirclePrizeIn struct {
		//0 prize 1 comment
		Action int                `json:"action,omitempty"`
		ID     primitive.ObjectID `json:"id,omitempty"   bson:"_id,omitempty"`
		Text   string             `json:"text,omitempty"`
		UID    int64              `json:"uid,omitempty"   bson:"uid,omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	CirclePrizeH struct {
		base ykit.RootTran
	}
)

func (r *CirclePrizeH) MakeLocalEndpoint(svc CirclePrizeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  CirclePrize ###########")
		spew.Dump(ctx)

		//todo
		in := request.(*CirclePrizeIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *CirclePrizeH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(CirclePrizeIn), ctx, req)
}

//个人实现,参数不能修改
func (r *CirclePrizeH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *CirclePrizeH) HandlerLocal(service CirclePrizeService,
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
func (r *CirclePrizeH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		//todo
		"POST",
		CirclePrize_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}
