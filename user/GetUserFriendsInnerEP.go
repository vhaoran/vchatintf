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
)

const (
	GetUserFriendsInner_H_PATH = "/GetUserFriendsInner"
)

//获取用户所有好友
type (
	GetUserFriendsInnerService interface {
		Exec(in *GetUserFriendsInnerIn) ([]*GetUserFriendsInnerOut, error)
	}

	//input data
	GetUserFriendsInnerIn struct {
		UID int64 `json:"uid,omitempty"`
	}

	//output data
	GetUserFriendsInnerOut struct {
		FriendID    int64  `json:"friend_id,omitempty" gorm:"unique_index:uf_uk_uid_friend_id;"`
		UserCodeRef string `json:"user_code_ref,omitempty" gorm:"size:50"`
		//冗余字段，提升性能,眤称
		NickRef string `json:"nick_ref,omitempty" gorm:"type:varchar(50);unique_index"`
		//冗余字段，提升性能,图标
		IconRef string `json:"icon_ref,omitempty" gorm:"type:varchar(100);unique_index"` //	好友备注(显示在我的朋友圈的名称) //varchar(50)	用户显示用好友的备 注主表修改时注意同步
	}

	// handler implements
	GetUserFriendsInnerH struct {
		base ykit.RootTran
	}
)

func (r *GetUserFriendsInnerH) MakeLocalEndpoint(svc GetUserFriendsInnerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("#############  GetUserFriendsInner ###########")
		spew.Dump(ctx)

		in := request.(*GetUserFriendsInnerIn)
		return svc.Exec(in)
	}
}

//个人实现,参数不能修改
func (r *GetUserFriendsInnerH) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(GetUserFriendsInnerIn), ctx, req)
}

//个人实现,参数不能修改
func (r *GetUserFriendsInnerH) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response []*GetUserFriendsInnerOut
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *GetUserFriendsInnerH) HandlerLocal(service GetUserFriendsInnerService,
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
func (r *GetUserFriendsInnerH) HandlerSD(mid []endpoint.Middleware,
	options ...tran.ServerOption) *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		MSTAG,
		"POST",
		GetUserFriendsInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse,
		mid,
		options...)
}

func (r *GetUserFriendsInnerH) ProxySD() endpoint.Endpoint {
	return r.base.ProxyEndpointSD(
		context.Background(),
		MSTAG,
		"POST",
		GetUserFriendsInner_H_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

//只用于内部调用 ，不从风头调用
var once_GetUserFriendsInner sync.Once
var local_GetUserFriendsInner_EP endpoint.Endpoint

func (r *GetUserFriendsInnerH) Call(in *GetUserFriendsInnerIn) ([]*GetUserFriendsInnerOut, error) {
	once_GetUserFriendsInner.Do(func() {
		local_GetUserFriendsInner_EP = new(GetUserFriendsInnerH).ProxySD()
	})
	//
	ep := local_GetUserFriendsInner_EP
	//
	result, err := ep(context.Background(), in)

	if err != nil {
		return nil, err
	}

	return result.([]*GetUserFriendsInnerOut), nil
}
