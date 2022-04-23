package user

import (
	"net/http"

	"github.com/blakeyi/gan-circle-server/api/internal/logic/sys/user"
	"github.com/blakeyi/gan-circle-server/api/internal/svc"
	"github.com/blakeyi/gan-circle-server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			var res types.Response
			res.SetErrorCode(types.ErrParam, err)
			httpx.OkJson(w, res)
			return
		}

		l := user.NewUserQueryLogic(r.Context(), svcCtx)
		resp, err := l.UserQuery(&req)
		if err != nil {
			resp.RetContent = err.Error()
		}
		httpx.OkJson(w, resp)
	}
}
