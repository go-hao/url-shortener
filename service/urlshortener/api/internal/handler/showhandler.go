package handler

import (
	"net/http"

	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/logic"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/svc"
	"github.com/go-hao/url-shortener/service/urlshortener/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/go-hao/zero/xerrors"
	"github.com/go-hao/zero/xhttp"
)

func ShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowReq
		if err := httpx.Parse(r, &req); err != nil {
			switch e := svcCtx.ErrBadReqest.(type) {
			case *xerrors.Error:
				xhttp.Json(r.Context(), w, e.Detail(err))
			case xerrors.Error:
				xhttp.Json(r.Context(), w, e.Detail(err))
			default:
				xhttp.Json(r.Context(), w, err)
			}
			return
		}

		l := logic.NewShowLogic(r.Context(), svcCtx)
		resp, err := l.Show(&req)
		if err != nil {
			xhttp.Json(r.Context(), w, err)
		} else {
			// xhttp.Json(r.Context(), w, resp)
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
		}
	}
}
