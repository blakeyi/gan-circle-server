package login

import (
	"net/http"

	"github.com/blakeyi/gan-circle-server/api/internal/logic/sys/login"
	"github.com/blakeyi/gan-circle-server/api/internal/svc"
	"github.com/blakeyi/gan-circle-server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			var res types.Response
			res.SetErrorCode(types.ErrParam, err)
			httpx.OkJson(w, res)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			resp.RetContent = err.Error()
		}
		httpx.OkJson(w, resp)
	}
}
