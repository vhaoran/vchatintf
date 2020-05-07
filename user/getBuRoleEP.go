package user

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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	GetBuRole_H_PATH = "/GetBuRole"
)

type (
	GetBuRoleService interface {
		Exec(in *GetBuRoleIn) (*GetBuRoleOut, error)
	}

	//input data
	GetBuRoleIn struct {
		ID primitive.ObjectID `json:"id"   bson:"_id"`
	}

	//output data
	GetBuRoleOut struct {
		RoleName string `json:"role_name"`
	}

	// handler implements
	GetBuRoleH struct {
		base ykit.RootTran
	}
)

func (r *GetBuRoleH) MakeLocalEndpoint(svc GetBuRoleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetBuRole ###########")
		spew.Dump(ctx)

		in := request.(*GetBuRoleIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetBuRoleH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetBuRoleIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetBuRoleH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response *GetBuRoleOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetBuRoleH) HandlerLocal(service GetBuRoleService,
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
func (r *GetBuRoleH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuRole_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetBuRoleH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetBuRole_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_GetBuRole sync.Once
var local_GetBuRole_EP endpoint.Endpoint

func (r *GetBuRoleH) Call(in *GetBuRoleIn) (*GetBuRoleOut, error) {
	once_GetBuRole.Do(func() {
		local_GetBuRole_EP = new(GetBuRoleH).ProxySD()
	})
	//
	ep := local_GetBuRole_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.(*GetBuRoleOut), nil
}
