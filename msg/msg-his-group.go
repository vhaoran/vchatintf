package msg

//for snippet用于标准返回值的微服务接口

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/common/ypage"
	"github.com/vhaoran/vchat/lib/ykit"
	"net/http"
)

const (
	MsgHisGroup_HANDLER_PATH = "/MsgHisGroup"
)

type (
	MsgHisGroupService interface {
		Exec(in *MsgHisGroupIn) (*ypage.PageResult, error)
	}

	//input data
	MsgHisGroupIn struct {
		ypage.PageBean
		GID int64 `json:"gid omitempty"`
	}

	//output data

	// handler implements
	MsgHisGroupHandler struct {
		base ykit.RootTran
	}
)

func (r *MsgHisGroupHandler) MakeLocalEndpoint(svc MsgHisGroupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  MsgHisGroup ###########")
		spew.Dump(ctx)

		in := request.(*MsgHisGroupIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *MsgHisGroupHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(MsgHisGroupIn), ctx, req)
}

//个人实现,参数不能修改
func (r *MsgHisGroupHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *ypage.PageResult
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *MsgHisGroupHandler) HandlerLocal(service MsgHisGroupService,
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
func (r *MsgHisGroupHandler) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		MsgHisGroup_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *MsgHisGroupHandler) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		MsgHisGroup_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}
