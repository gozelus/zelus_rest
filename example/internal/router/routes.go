package router

import (
	"github.com/gozelus/zelus_rest"
	"net/http"
)

type ShoppingCartControllerInter interface {
	GoodsList(w http.ResponseWriter, req *http.Request)
	AddGoods(w http.ResponseWriter, req *http.Request)
	RemoveGoods(w http.ResponseWriter, req *http.Request)
}
type GoodsControllerInter interface {
	GetGoods(w http.ResponseWriter, req *http.Request)
}
type UserControllerInter interface {
	CreateUser(w http.ResponseWriter, req *http.Request)
	GetUser(w http.ResponseWriter, req *http.Request)
}
type Router struct {
	ShoppingCart ShoppingCartControllerInter
	goods        GoodsControllerInter
	user         UserControllerInter
}

func NewRouter(user UserControllerInter) *Router {
	return &Router{user: user}
}

var ShoppingCartGroup = func(r *Router) []rest.Route {
	return []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/shopping_cart/goods_list",
			Handler: r.ShoppingCart.GoodsList,
		},
		{
			Method:  http.MethodPost,
			Path:    "/shopping_cart/add_goods",
			Handler: r.ShoppingCart.AddGoods,
		},
		{
			Method:  http.MethodPost,
			Path:    "/shopping_cart/remove_goods",
			Handler: r.ShoppingCart.RemoveGoods,
		},
	}
}
var GoodsGroup = func(r *Router) []rest.Route {
	return []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/goods/get",
			Handler: r.goods.GetGoods,
		},
	}
}
var UserGroup = func(r *Router) []rest.Route {
	return []rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/user/get",
			Handler: r.user.GetUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/user/create",
			Handler: r.user.CreateUser,
		},
	}
}

func (r *Router) Routes() []rest.Route {
	return append(append(UserGroup(r), GoodsGroup(r)...), ShoppingCartGroup(r)...)
}
