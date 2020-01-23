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
	"net/http"
)

const (
	OnLineNotify_H_PATH = "/OnLineNotify"
)

type (
	OnLineNotifyService interface {
		Exec(ctx context.Context, in *OnLineNotifyIn) (*ykit.Result, error)
	}

	//input data
	OnLineNotifyIn struct {
		UID int64 `json:"uid omitempty"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	OnLineNotifyH struct {
		base ykit.RootTran
	}
)

func (r *OnLineNotifyH) MakeLocalEndpoint(svc OnLineNotifyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  OnLineNotify ###########")
		spew.Dump(ctx)

		in := request.(*OnLineNotifyIn)
		return svc.Exec(ctx, in)
	}
}

//个人实现,参数不能修改
func (r *OnLineNotifyH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	uid := ykit.GetUIDOfContext(ctx)
	bean := &OnLineNotifyIn{
		UID: uid,
	}

	return bean, nil
}

//个人实现,参数不能修改
func (r *OnLineNotifyH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *OnLineNotifyH) HandlerLocal(service OnLineNotifyService,
	mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {

	ep := r.MakeLocalEndpoint(service)
	for _, f := range mid {
		ep = f(ep)
	}

	before := tran.ServerBefore(ykit.Jwt2ctx())

	opts := make([]tran.ServerOption, 0)
	opts = append(opts, before)
	opts = append(opts, options...)

	handler := tran.NewServer(
		ep,
		r.DecodeRequest,
		r.base.EncodeResponse,
		opts...)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *OnLineNotifyH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		OnLineNotify_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *OnLineNotifyH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		OnLineNotify_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
