package main

import (
	"encoding/json"
	"io/ioutil"
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
		t.Error("API returned empty hardware type list! Make sure the database not empty")
	}
	if htl.Status != "ok" {
		t.Error("API returned bad status")
	}
}
