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
	cfg.DebugMode = true
	// JWT authorization
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "inventory",
		// TO DO: Load this from config file
		Key:             []byte("ThisIsVerySecretKey"),
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
	// API endpoints
	hwAPI := r.Group("/api/v0.1/hw")
	userAPI := r.Group("/api/v0.1/user")
	// Use token auth to access api endpoints
	hwAPI.Use(authMiddleware.MiddlewareFunc())
	userAPI.Use(authMiddleware.MiddlewareFunc())
	// Hardware
	hwAPI.POST("add", postNewHardware)
	hwAPI.GET("all", getHardware)
	hwAPI.GET("full", getHardwareFullList)
	hwAPI.PUT("update", putUpdateHardware)
	hwAPI.DELETE("delete", postDeleteHardware)
	// Hardware types
	hwAPI.GET("type", getHardwareTypes)
	hwAPI.POST("type/add", postNewHardwareType)
	// Manufacturers
	hwAPI.GET("manufacturer", getManufacturers)
	hwAPI.POST("manufacturer/add", postNewManufacturer)
	hwAPI.PUT("manufacturer/update", putUpdateManufacturer)
	hwAPI.DELETE("manufacturer/delete", postDeleteManufacturer)
	// Models
	hwAPI.GET("model", getModels)
	hwAPI.POST("model/add", postNewModel)
	hwAPI.PUT("model/update", putUpdateModel)
	hwAPI.DELETE("model/delete", postDeleteModel)
	// Users
	userAPI.POST("user", postNewUser)
	userAPI.PUT("user_update", putUpdateUser)
	userAPI.DELETE("user_delete", postDeleteUser)
	// Run app
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

// Shortcut for debug logging
func logger(msg string) {
	if cfg.DebugMode {
		log.Println(msg)
	}
}

// shortcut for error handling
func check(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
