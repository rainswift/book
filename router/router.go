package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var MySecret = []byte("secret")

//func ParseToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenString := c.GetHeader("Authorization")
//		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
//			responose.FailedToken("token过期", c)
//			c.Abort()
//			return
//		}
//		tokenString = tokenString[7:]
//		// 解析token
//		token, err := jwt.ParseWithClaims(tokenString, &models.BlogUser{},
//			func(token *jwt.Token) (i interface{}, err error) {
//				return MySecret, nil
//			})
//		if err != nil {
//			responose.FailedToken("token过期", c)
//			c.Abort()
//			return
//		}
//		if claims, ok := token.Claims.(*models.BlogUser); ok && token.Valid { // 校验token
//			user, _ := dao.Mgr.GetLoadUser(claims.Username)
//			c.Set("user", user)
//			c.Next()
//			return
//		}
//		response.FailedToken("token过期", c)
//		c.Abort()
//	}
//}

func Start() {
	e := gin.Default()
	// 实现跨域访问
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	})
	e.Use(mwCORS)
	//e.POST("/add", api.AddUser)
	//e.POST("/login", api.Login)
	//e.POST("/userDelect", ParseToken(), api.UserDelect)

	e.Run()
}
