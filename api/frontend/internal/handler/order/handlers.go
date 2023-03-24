package handler

import (
	"net/http"

	"dcs/api/frontend/internal/logic/order"
	"dcs/api/frontend/internal/svc"
	"dcs/api/frontend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewOrderLogic(r.Context(), ctx)
		resp, err := l.Order(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}












