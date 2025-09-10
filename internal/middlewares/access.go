package middlewares

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/utils"
	"github.com/kodacampmain/koda3_gin/pkg"
)

func Access(roles ...string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// ambil data claim
		claims, isExist := ctx.Get("claims")
		if !isExist {
			utils.HandleMiddlewareError(ctx, http.StatusUnauthorized, "silahkan login kembali", "Unauthorized Access")
			return
		}
		user, ok := claims.(pkg.Claims)
		if !ok {
			utils.HandleMiddlewareError(ctx, http.StatusInternalServerError, "Internal Server Error", "cannot cast into pkg.claims")
			return
		}
		if !slices.Contains(roles, user.Role) {
			utils.HandleMiddlewareError(ctx, http.StatusForbidden, "Anda tidak punya hak akses untuk resource ini", "Forbidden Access")
			return
		}
		ctx.Next()
	}
}
