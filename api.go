package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API handler. Returns available types of hardware
func getHardwareTypes(c *gin.Context) {
	hw, err := readHardwareTypes(cfg.DB.FileName)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "result": hw})
}

// API handler. Returns available hardware manufacturers
func getManufacturers(c *gin.Context) {
	m, err := readManufacturers(cfg.DB.FileName)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "result": m})
}

// API handler. Returns full listed of hardware
func getHardwareFullList(c *gin.Context) {
	h, err := readHardwareFullList(cfg.DB.FileName)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "result": h})
}

// API handler. Adds new user into database
func postUser(c *gin.Context) {
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
	}
	add, err := addUser(cfg.DB.FileName, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": strconv.Itoa(add)})
	}
}
