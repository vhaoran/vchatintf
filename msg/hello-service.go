package msg

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/ykit"
	"github.com/vhaoran/vchat/lib/ylog"
	"net/http"
)

const (
	Hello_H_PATH = "/Hello"
)

type (
	HelloService interface {
		Exec(in *HelloIn) (*ykit.Result, error)
	}

	//input data
	HelloIn struct {
		UID int64  `json:"uid,omitempty"`
		S   string `json:"s"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	HelloH struct {
		base ykit.RootTran
	}
)

func (r *HelloH) MakeLocalEndpoint(svc HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  Hello ###########")
		spew.Dump(ctx)

		in := request.(*HelloIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *HelloH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	ylog.Debug(" #### enter DecodeRequest ")
	ylog.DebugDump("ctx:", ctx)

	return r.base.DecodeRequest(new(HelloIn), ctx, req)
}

//个人实现,参数不能修改
func (r *HelloH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *HelloH) HandlerLocal(service HelloService,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	// 访问 request内容,丁当于Java中的拦截器
	before := tran.ServerBefore(func(ctx context.Context, req *http.Request) context.Context {
		uid := req.Header.Get("uid")
		if len(uid) == 0 {
			uid = "123"
		}

		fmt.Println("------------before host:", req.Host)
		return context.WithValue(ctx, "uid", uid)
	})
	l := make([]tran.ServerOption, 0)
	l = append(l, before)
	l = append(l, options...)

	ep := r.MakeLocalEndpoint(service)
	for _, f := range mid {
		ep = f(ep)
	}

	handler := tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
		l...)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *HelloH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		Hello_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *HelloH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		Hello_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
