package cors

import "github.com/gin-gonic/gin"

func AllowCors(gc *gin.Context) {
	gc.Writer.Header().Set("Access-Control-Allow-Origin", "*")
}
