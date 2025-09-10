package middlewares

import (
	"log"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(ctx *gin.Context) {
	// if ctx.Request.Method == http.MethodPost {
	// 	ctx.AbortWithStatus(http.StatusMethodNotAllowed)
	// 	return
	// }
	// memasangkan header-header CORS
	// setup whitelist origin
	whitelist := []string{"http://127.0.0.1:5500", "http://127.0.0.1:3001"}
	origin := ctx.GetHeader("Origin")
	if slices.Contains(whitelist, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
	} else {
		log.Printf("Origin is not in the whitelist: %s", origin)
	}
	// header untuk preflight cors
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	// tangani apabila bertemu preflight
	if ctx.Request.Method == http.MethodOptions {
		// ctx.Header("X-DEBUG", "preflight-handled")
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	// ctx.Header("X-DEBUG", "actual request")
	ctx.Next()
}
