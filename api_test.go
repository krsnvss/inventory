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
	cfg.DebugMode = true
	gin.SetMode(gin.TestMode)
}

// Test posNewHardware function
func TestPostNewHardware(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	h := hardware{
		Name:           "test-pc",
		SerialNumber:   "TPC009003",
		Manufacturer:   1,
		Model:          1,
		Type:           1,
		ProductionDate: "2010-12-22",
		PurchaseDate:   "2020-12-30",
		Barcode:        "01009",
		CPUName:        "Intel Pentium D",
		CPUCores:       2,
		CPUMaxClock:    2.6,
		RAM:            2,
		Disk:           250,
		UserName:       "John Doe",
	}
	jh, err := json.Marshal(h)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("POST", "/", bytes.NewReader(jh))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	postNewHardware(c)
	if w.Code != 200 {
		t.Error("postManufacturer returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No hardware items were added")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test getHardwarefunction
func TestGetHardware(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	getHardware(c)
	if w.Code != 200 {
		t.Error("getHardware returned", w.Code)
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
		t.Error("API returned empty hardware list! Make sure the database is not empty")
	}
	if hl.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test getOneHardware function
func TestGetOneHardware(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	param := c.Request.URL.Query()
	param.Add("name", "test-pc")
	c.Request.URL.RawQuery = param.Encode()
	getOneHardware(c)
	if w.Code != 200 {
		t.Error("getOneHardware returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var h hardwareItem
	err = json.Unmarshal(b, &h)
	if err != nil {
		t.Error(err.Error())
	}
	if h.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test postUpdateHardware function
func TestPutUpdateHardware(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	h, err := readOneHardwareByName(cfg.DB.FileName, "test-pc")
	if err != nil {
		t.Error(err.Error())
	}
	h.Name = "Another-test-pc"
	jh, err := json.Marshal(h)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("PUT", "/", bytes.NewReader(jh))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	putUpdateHardware(c)
	if w.Code != 200 {
		t.Error("updateModel returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were updated")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test postDeleteHardware function
func TestPostDeleteHardware(t *testing.T) {
	h, err := readOneHardwareByName(cfg.DB.FileName, "Another-test-pc")
	if err != nil {
		t.Error(err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jh, err := json.Marshal(h)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("DELETE", "/", bytes.NewReader(jh))
	if err != nil {
		t.Error(err.Error())
	}
	c.Request.Header.Set("Content-Type", "application/json")
	postDeleteHardware(c)
	if w.Code != 200 {
		t.Error("postDeleteModel returned ", w.Code)
	}
}

// Test postNewHardwareType function
func TestPostNewHardwareType(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	h := hardwareType{Name: "laptop"}
	jh, err := json.Marshal(h)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("POST", "/", bytes.NewReader(jh))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	postNewHardwareType(c)
	if w.Code != 200 {
		t.Error("postNewHardwareType returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were updated")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
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

// Test putUpdateHardwareType function
func TestPutUpdateHardwareType(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	ht, err := readOneHardwareTypeByName(cfg.DB.FileName, "laptop")
	if err != nil {
		t.Error(err.Error())
	}
	ht.Name = "NotALaptop"
	jht, err := json.Marshal(ht)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("PUT", "/", bytes.NewReader(jht))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	putUpdateHardwareType(c)
	if w.Code != 200 {
		t.Error("updateHardwareType returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No hardware types were updated")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test postDeleteHardwareType function
func TestPostDeleteHardwareType(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	h := hardwareType{Name: "laptop"}
	jh, err := json.Marshal(h)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("POST", "/", bytes.NewReader(jh))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	postDeleteHardwareType(c)
	if w.Code != 200 {
		t.Error("postDeleteHardwareType returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No hardware types were deleted")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test posNewManufacturer function
func TestPostNewManufacturer(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	m := manufacturer{Name: "NoName", Logo: "noname.png"}
	jm, err := json.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("POST", "/", bytes.NewReader(jm))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	postNewManufacturer(c)
	if w.Code != 200 {
		t.Error("postManufacturer returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were added")
	}
	if sr.Status != "ok" {
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

// Test putUpdateManufacturer function
func TestPutUpdateManufacturer(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	m := manufacturer{Name: "NoName", Logo: "nonameOne.png"}
	jm, err := json.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("PUT", "/", bytes.NewReader(jm))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	putUpdateManufacturer(c)
	if w.Code != 200 {
		t.Error("updateManufacturer returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were updated")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test postDeleteManufacturer function
func TestPostDeleteManufacturer(t *testing.T) {
	m := manufacturer{Name: "NoName", Logo: "noname.png"}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jm, err := json.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("DELETE", "/", bytes.NewReader(jm))
	if err != nil {
		t.Error(err.Error())
	}
	c.Request.Header.Set("Content-Type", "application/json")
	postDeleteManufacturer(c)
	if w.Code != 200 {
		t.Error("postDeleteManufacturer returned ", w.Code)
	}
}

// Test postNewManufacturer function
func TestPostNewModel(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	m := model{Manufacturer: 1, Name: "ComputerCase", Type: 1}
	jm, err := json.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("POST", "/", bytes.NewReader(jm))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	postNewModel(c)
	if w.Code != 200 {
		t.Error("postModel returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were added")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test getModels function
func TestGetModels(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	getModels(c)
	if w.Code != 200 {
		t.Error("getModels returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var ml modelsList
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

// Test putUpdateModel function
func TestPutUpdateModel(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	om := model{Manufacturer: 1, Name: "ComputerCase", Type: 1}
	m, err := readOneModelByName(cfg.DB.FileName, om.Name)
	if err != nil {
		t.Error(err.Error())
	}
	m.Name = "NotAComputerCase"
	jm, err := json.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("PUT", "/", bytes.NewReader(jm))
	c.Request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err.Error())
	}
	putUpdateModel(c)
	if w.Code != 200 {
		t.Error("updateModel returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were updated")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test postDeleteModelfunction
func TestPostDeleteModel(t *testing.T) {
	om := model{Manufacturer: 1, Name: "NotAComputerCase", Type: 1}
	m, err := readOneModelByName(cfg.DB.FileName, om.Name)
	if err != nil {
		t.Error(err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jm, err := json.Marshal(m)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("DELETE", "/", bytes.NewReader(jm))
	if err != nil {
		t.Error(err.Error())
	}
	c.Request.Header.Set("Content-Type", "application/json")
	postDeleteModel(c)
	if w.Code != 200 {
		t.Error("postDeleteModel returned ", w.Code)
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
	var hl hardwareFullList
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

// Test posNewtUser function
func TestPostNewUser(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	u := user{
		Login:     "vasyan2003",
		FirstName: "Vasyan",
		LastName:  "Pupkov",
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
	postNewUser(c)
	if w.Code != 200 {
		t.Error("postUser returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were added")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test putUpdateUser function
func TestPutUpdateUser(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	u := user{
		Login:     "vasyan2003",
		FirstName: "Vasyan",
		LastName:  "Pupkov-Gromyko",
		Password:  "NewSecureKey",
		Group:     "1",
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
	putUpdateUser(c)
	if w.Code != 200 {
		t.Error("updateUser returned", w.Code)
	}
	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err.Error())
	}
	var sr serverResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		t.Error(err.Error())
	}
	if sr.Result == 0 {
		t.Error("API returned 0! No users were updated")
	}
	if sr.Status != "ok" {
		t.Error("API returned bad status")
	}
}

// Test postDeleteUser function
func TestPostDeleteUser(t *testing.T) {
	var u user
	u.Login = "vasyan2003"
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	ju, err := json.Marshal(u)
	if err != nil {
		t.Error(err.Error())
	}
	c.Request, err = http.NewRequest("DELETE", "/", bytes.NewReader(ju))
	if err != nil {
		t.Error(err.Error())
	}
	c.Request.Header.Set("Content-Type", "application/json")
	postDeleteUser(c)
	if w.Code != 200 {
		t.Error("postDeleteUser returned ", w.Code)
	}
}
