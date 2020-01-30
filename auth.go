package main

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login type for token request
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User
type user struct {
	ID        int    `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Group     string `json:"group,omitempty"`
}

// Payload function for authorization middleware
func middlewarePayload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*user); ok {
		return jwt.MapClaims{
			cfg.Token.IdentityKey: v.Login,
		}
	}
	return jwt.MapClaims{}
}

// Identity function for authorization middleware
func middlewareIdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &user{
		Login: claims[cfg.Token.IdentityKey].(string),
	}
}

// Autorization function for authorization middleware
// Checks if user is registered
func middlewareAuthenticator(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	// if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
	valid, err := checkPasswordFromDB(cfg.DB.FileName, userID, password)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	if valid {
		u, err := readOneUserByLogin(cfg.DB.FileName, userID)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}
		log.Println(u)
		return &user{
			Login:     userID,
			LastName:  u.LastName,
			FirstName: u.FirstName,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

// Authorizator function for authorization middleware
// Grants user access
func middlewareAuthorizator(data interface{}, c *gin.Context) bool {
	u, err := readOneUserByLogin(cfg.DB.FileName, data.(*user).Login)
	if err != nil {
		return false
	}
	if v, ok := data.(*user); ok && v.Login == u.Login {
		return true
	}
	return false
}

// Unauthorized function for authorization middleware
// acts if user was not authorized
func middlewareUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// Generates hash of user password to store it in datatabase
func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(b), err
}

// Checks if password hashes are match
func checkHashMatch(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
