package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API handler. Returns all hardware items
func getHardware(c *gin.Context) {
	h, err := readHardware(cfg.DB.FileName)
	if err != nil {
		logger(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "result": h})
}

// API handler. Adds new hardware into database
func postNewHardware(c *gin.Context) {
	var h hardware
	err := c.ShouldBindJSON(&h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
	}
	res, err := writeHardware(cfg.DB.FileName, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
}

// API handler. Updates an existing manufacturer entry in database
func putUpdateHardware(c *gin.Context) {
	var h hardware
	err := c.ShouldBindJSON(&h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	res, err := updateHardware(cfg.DB.FileName, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
}

// API handler. Deletes an existing hardware entry
func postDeleteHardware(c *gin.Context) {
	var h hardware
	err := c.ShouldBindJSON(&h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	del, err := deleteHardware(cfg.DB.FileName, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": del})
	}
}

// API handler. Returns available types of hardware
func getHardwareTypes(c *gin.Context) {
	hw, err := readHardwareTypes(cfg.DB.FileName)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "result": hw})
}

// API handler. Adds new hardware type
func postNewHardwareType(c *gin.Context) {
	var ht hardwareType
	err := c.ShouldBindJSON(&ht)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad", "error": err.Error()})
	}
	res, err := writeHardwareType(cfg.DB.FileName, ht)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "bad", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
}

// API handler. Deletes an existing hardware type
func postDeleteHardwareType(c *gin.Context) {
	var ht hardwareType
	err := c.ShouldBindJSON(&ht)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad", "error": err.Error()})
	}
	res, err := deleteHardwareType(cfg.DB.FileName, ht)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "bad", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
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

// API handler. Adds new manufacturer into database
func postNewManufacturer(c *gin.Context) {
	var m manufacturer
	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
	}
	res, err := writeManufacturer(cfg.DB.FileName, m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
}

// API handler. Updates an existing manufacturer entry in database
func putUpdateManufacturer(c *gin.Context) {
	var m manufacturer
	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	res, err := updateManufacturer(cfg.DB.FileName, m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
}

// API handler. Deletes user entry
func postDeleteManufacturer(c *gin.Context) {
	var m manufacturer
	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	del, err := deleteManufacturer(cfg.DB.FileName, m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": del})
	}
}

// API handler. Returns available hardware models
func getModels(c *gin.Context) {
	m, err := readModels(cfg.DB.FileName)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "result": m})
}

// API handler. Adds new model into database
func postNewModel(c *gin.Context) {
	var m model
	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
	}
	res, err := writeModel(cfg.DB.FileName, m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
}

// API handler. Updates an existing model entry in database
func putUpdateModel(c *gin.Context) {
	var m model
	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	res, err := updateModel(cfg.DB.FileName, m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": res})
	}
}

// API handler. Deletes model entry
func postDeleteModel(c *gin.Context) {
	var m model
	err := c.ShouldBindJSON(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	del, err := deleteModel(cfg.DB.FileName, m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": del})
	}
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
func postNewUser(c *gin.Context) {
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
	}
	add, err := addUser(cfg.DB.FileName, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": add})
	}
}

// API handler. Updates an existing user entry in database
func putUpdateUser(c *gin.Context) {
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	up, err := updateUser(cfg.DB.FileName, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": up})
	}
}

// API handler. Deletes user entry
func postDeleteUser(c *gin.Context) {
	var u user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	}
	del, err := deleteUser(cfg.DB.FileName, u.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		logger(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "result": del})
	}
}
