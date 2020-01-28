package main

import (
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var cfg appConfig

func main() {
	log.Println("Started!")
	// TO DO: load configuration from file
	cfg.DB.FileName = "./db/hardware.db"
	cfg.Token.IdentityKey = "id"
	// JWT authorization
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "inventory",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     cfg.Token.IdentityKey,
		PayloadFunc:     middlewarePayload,
		IdentityHandler: middlewareIdentityHandler,
		Authenticator:   middlewareAuthenticator,
		Authorizator:    middlewareAuthorizator,
		Unauthorized:    middlewareUnauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Authentication endpoints
	auth := r.Group("auth")
	r.POST("login", authMiddleware.LoginHandler)
	auth.GET("refresh", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.GET("hello", helloHandler)
	// API endpoints
	api := r.Group("/api/v0.1")
	// Use token auth to access api endpoints
	api.Use(authMiddleware.MiddlewareFunc())
	api.GET("htypes", getHardwareTypes)
	api.GET("manufacturers", getManufacturers)
	api.GET("hfull", getHardwareFullList)

	// Run app
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	u, _ := c.Get(cfg.Token.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[cfg.Token.IdentityKey],
		"userName": u.(*user).Login,
		"text":     "Hello World.",
	})
}
