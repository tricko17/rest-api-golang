package main

import (
	"net/http"
	"os"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

  "dibo-go-api/core"
  "dibo-go-api/library"
  "dibo-go-api/models"

	"dibo-go-api/controllers/clients"
	"dibo-go-api/controllers/consultants"
	"dibo-go-api/controllers/brands"
	"dibo-go-api/controllers/projects"
	// "fmt"
)

func userFinder(userId string)(map[string]string){
		var consultant models.Users

	  db:= database.ConnectionPostgres()
	  // db.AutoMigrate(models.Consultant{})

		db.Find(&consultant, "username = ?", userId)

		person := make(map[string]string)
		person["username"] = consultant.Username
		person["password"] = consultant.Password
		return person
}

func main() {

	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zon",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			pwd := encryptor.GetMD5Hash(password)

			cons := userFinder(userId)

			if (cons["username"] == userId && pwd== cons["password"]) {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {

			cons := userFinder(userId)
			if (cons["username"] == userId) {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	if port == "" {
		port = "8000"
	}

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	consul := r.Group("/api/v1/consultants")
	consul.Use(authMiddleware.MiddlewareFunc())
	{
		consul.POST("/", consultant.CreateConsultant)
		consul.GET("/", consultant.FetchAllConsultant)
		consul.GET("/:id/clients", client.FetchAllClientBaseConsultant)
		consul.GET("/:id", consultant.FetchSingleConsultant)
		consul.PUT("/:id", consultant.UpdateConsultant)
		consul.DELETE("/:id", consultant.DeleteConsultant)
	}

	cli := r.Group("/api/v1/clients")
	cli.Use(authMiddleware.MiddlewareFunc())
	{
		cli.POST("/", client.CreateClient)
		cli.GET("/", client.FetchAllClient)
		cli.GET("/:id", consultant.FetchSingleConsultant)
		cli.GET("/:id/brands", brand.FetchBrandBaseClient)
		cli.PUT("/:id", consultant.UpdateConsultant)
		cli.DELETE("/:id", consultant.DeleteConsultant)
	}

	brn := r.Group("/api/v1/brands")
	brn.Use(authMiddleware.MiddlewareFunc())
	{
		brn.POST("/", brand.CreateBrand)
		brn.GET("/", brand.FetchAllBrand)
		brn.GET("/:id", brand.FetchSingleBrand)
		brn.PUT("/:id", brand.UpdateBrand)
		brn.DELETE("/:id", brand.DeleteBrand)
		brn.GET("/:id/projects", project.FetchProjectBaseBrand)
	}

	pro := r.Group("/api/v1/projects")
	pro.Use(authMiddleware.MiddlewareFunc())
	{
		pro.POST("/", project.CreateProject)
		pro.GET("/", project.FetchAllProject)
		pro.GET("/:id", project.FetchSingleProject)
		pro.PUT("/:id", project.UpdateProject)
		pro.DELETE("/:id", project.DeleteProject)
	}


	http.ListenAndServe(":"+port, r)
}
