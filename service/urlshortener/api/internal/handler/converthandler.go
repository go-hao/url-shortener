package handler

import (
	"net/http"

	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/logic"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/svc"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/go-hao/zero/xhttp"
)

func ConvertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ConvertReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.Json(r.Context(), w, err)
			return
		}

		l := logic.NewConvertLogic(r.Context(), svcCtx)
		resp, err := l.Convert(&req)
		if err != nil {
			xhttp.Json(r.Context(), w, err)
		} else {
			xhttp.Json(r.Context(), w, resp)
		}
	}
}
