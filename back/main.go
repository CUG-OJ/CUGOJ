package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

//获取新闻

func main() {
	//连接到MYSQL
	//usr := "root:MYcs2050@tcp(127.0.0.1:3306)/geo?charset=utf8"
	//db, _ := sql.Open("mysql", usr)
	////r, _ := db.Exec("select * from geo")

	//r := gin.Default()
	r := gin.New()
	r.Use(Cors())
	r.POST("/api/login", func(c *gin.Context) {
		//	打印登录信息
		username := c.PostForm("username")
		password := c.PostForm("password")
		//生成token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"password": password,
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(500, gin.H{
				"message": "生成token失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"token": tokenString,
		})
	})
	err := r.Run("127.0.0.1:5174")
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
