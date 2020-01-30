package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	cfg.DB.FileName = "./db/hardware.db"
	gin.SetMode(gin.TestMode)
}

// Test getHardwareTypes function
func TestGetHardwareType(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	getHardwareTypes(c)
	//c.Params = []gin.Param{gin.Param{Key: "k", Value: "v"}}
	if w.Code != 200 {
		t.Error("getHardwareTypes returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var htl hardwareTypeList
	err = json.Unmarshal(b, &htl)
	if err != nil {
		t.Error(err.Error())
	}
	if len(htl.Result) == 0 {
		t.Error("API returned empty hardware type list! Make sure the database is not empty")
	}
	if htl.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test getManufacturers function
func TestGetManufacturers(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	getManufacturers(c)
	//c.Params = []gin.Param{gin.Param{Key: "k", Value: "v"}}
	if w.Code != 200 {
		t.Error("getManufacturers returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var ml manufacturersList
	err = json.Unmarshal(b, &ml)
	if err != nil {
		t.Error(err.Error())
	}
	if len(ml.Result) == 0 {
		t.Error("API returned empty manufacturers list! Make sure the database is not empty")
	}
	if ml.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test getHardwareFullList function
func TestGetHardwareFullList(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	getHardwareFullList(c)
	//c.Params = []gin.Param{gin.Param{Key: "k", Value: "v"}}
	if w.Code != 200 {
		t.Error("getHardwareFullList returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var hl hardwareList
	err = json.Unmarshal(b, &hl)
	if err != nil {
		t.Error(err.Error())
	}
	if len(hl.Result) == 0 {
		t.Error("API returned empty manufacturers list! Make sure the database is not empty")
	}
	if hl.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test postUser function
func TestPostUser(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	u := user{
		Login:     "username",
		FirstName: "Vasyan",
		LastName:  "Pupkiv",
		Password:  "SecureKey",
		Group:     "2",
	}
	ju, err := json.Marshal(u)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("POST", "/", bytes.NewReader(ju))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	// c.Params = []gin.Param{
	// 	gin.Param{Key: "login", Value: "Conanthebarbarian"},
	// 	gin.Param{Key: "password", Value: "qwerty"},
	// 	gin.Param{Key: "first_name", Value: "Conan"},
	// 	gin.Param{Key: "last_name", Value: "Barbarian"}}
	//c.Params = []gin.Param{gin.Param{Key: "k", Value: "v"}}
	postUser(c)
	if w.Code != 200 {
		t.Error("postUsser returned", w.Code)
	}
	// b, err := ioutil.ReadAll(w.Body)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
	// var hl hardwareList
	// err = json.Unmarshal(b, &hl)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
	// if len(hl.Result) == 0 {
	// 	t.Error("API returned empty manufacturers list! Make sure the database is not empty")
	// }
	// if hl.Status != "ok" {
	// 	t.Error("API returned bad status")
	// }
}
