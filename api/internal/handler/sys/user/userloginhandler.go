package user

import (
	"net/http"

	"github.com/blakeyi/gan-circle-server/api/internal/common/errorx"
	"github.com/blakeyi/gan-circle-server/api/internal/logic/sys/user"
	"github.com/blakeyi/gan-circle-server/api/internal/svc"
	"github.com/blakeyi/gan-circle-server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		ret := errorx.Response{}
		if err := httpx.Parse(r, &req); err != nil {
			ret.SetErrorCode(errorx.ErrParam, err)
			httpx.OkJson(w, ret)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			ret.SetErrorCode(errorx.ErrLoginFail, err)
			httpx.OkJson(w, ret)
			return
		}
		ret.SetErrorCode(errorx.Succeed)
		ret.RetContent = resp
		httpx.OkJson(w, ret)
	}
}
