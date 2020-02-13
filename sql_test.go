package main

import (
	"reflect"
	"testing"
)

func init() {
	cfg.DB.FileName = "./db/hardware.db"
}

// Test function addUser
func TestAddUser(t *testing.T) {
	u := user{
		Login:     "username",
		FirstName: "Vasya",
		LastName:  "Pupkov",
		Password:  "SecureKey",
		Group:     "1",
	}
	v, err := addUser("./db/hardware.db", u)
	if err != nil {
		t.Error(err.Error())
	}
	if v == 0 {
		t.Error("Failed to write user!")
	}
	u, err = readOneUserByID("./db/hardware.db", v)
	if err != nil {
		t.Error(err.Error())
	}
	if len(u.Login) == 0 {
		t.Error("Failed to check written user!")
	}
}

// Test function updateUser
func TestUpdateUser(t *testing.T) {
	u := user{
		Login:     "username",
		FirstName: "Vasyan",
		LastName:  "Pupkov",
		Password:  "SecureKey",
		Group:     "2",
	}
	su, err := readOneUserByLogin("./db/hardware.db", u.Login)
	if err != nil {
		t.Error(err)
	}
	u.ID = su.ID
	u.Password = su.Password
	r, err := updateUser("./db/hardware.db", u)
	if err != nil {
		t.Error(err)
	}
	if r == 0 {
		t.Error("No records were updated!")
	}
}

// Test function readUsers
func TestReadUsers(t *testing.T) {
	vl, err := readUsers("./db/hardware.db")
	if err != nil {
		t.Error(err.Error())
	}
	if len(vl) == 0 {
		t.Error("Failed to read users! Make sure there're records in user table")
	}
	if reflect.TypeOf(vl[0].Login).Kind() != reflect.String {
		t.Error("Login is not a string! Type of login is", reflect.TypeOf(vl[0].Login).Kind())
	}
	vl, err = readUsers("emptyTest.db")
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function readOneUserByID
func TestReadOneUserByID(t *testing.T) {
	u, err := readOneUserByLogin("./db/hardware.db", "username")
	if err != nil {
		t.Error(err.Error())
	}
	v, err := readOneUserByID("./db/hardware.db", u.ID)
	if err != nil {
		t.Error(err.Error())
	}
	if len(v.Login) == 0 {
		t.Error("Failed to read users! Make sure there're records in user table")
	}
	if reflect.TypeOf(v.Login).Kind() != reflect.String {
		t.Error("Login is not a string! Type of login is", reflect.TypeOf(v.Login).Kind())
	}
	v, err = readOneUserByID("emptyTest.db", 1)
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function readOneUserByLogin
func TestReadOneUserByLogin(t *testing.T) {
	v, err := readOneUserByLogin("./db/hardware.db", "username")
	if err != nil {
		t.Error(err.Error())
	}
	if len(v.Login) == 0 {
		t.Error("Failed to read users! Make sure there're records in user table")
	}
	if reflect.TypeOf(v.Login).Kind() != reflect.String {
		t.Error("Login is not a string! Type of login is", reflect.TypeOf(v.Login).Kind())
	}
	v, err = readOneUserByID("emptyTest.db", 1)
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function checkPasswordFromDB
func TestCheckPasswordFromDB(t *testing.T) {
	v, err := checkPasswordFromDB("./db/hardware.db", "username", "SecureKey")
	if err != nil {
		t.Error(err)
	}
	if !v {
		t.Error("Passwords doesn't match!")
	}
}

// Test function deleteUser
func TestDeleteUser(t *testing.T) {
	v, err := deleteUser("./db/hardware.db", "username")
	if err != nil {
		t.Error(err.Error())
	}
	if v == 0 {
		t.Error("No user records affected!")
	}
	v, err = deleteUser("./db/hardware.db", "wrongusername")
	if err == nil && v != 0 {
		t.Error("Wrong user, but no error!")
	}
}

// Test function wrieHardwareType
func TestWriteHardwareType(t *testing.T) {
	ht := hardwareType{Name: "laptop"}
	res, err := writeHardwareType(cfg.DB.FileName, ht)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No hardware types were added!")
	}
}

// Test function readHardwareTypes
func TestReadHardwareTypes(t *testing.T) {
	vl, err := readHardwareTypes("./db/hardware.db")
	if err != nil {
		t.Error(err.Error())
	}
	if len(vl) == 0 {
		t.Error("Failed to read users! Make sure there're records in user table")
	}
	if reflect.TypeOf(vl[0].ID).Kind() != reflect.Int {
		t.Error("ID is not an Integer! Type of ID is", reflect.TypeOf(vl[0].ID).Kind())
	}
	if reflect.TypeOf(vl[0].Name).Kind() != reflect.String {
		t.Error("Name is not a string! Type of name is", reflect.TypeOf(vl[0].Name).Kind())
	}
	vl, err = readHardwareTypes("emptyTest.db")
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function readOneHardwareTypeByName
func TestReadOneHardwareTypeByName(t *testing.T) {
	h, err := readOneHardwareTypeByName(cfg.DB.FileName, "laptop")
	if err != nil {
		t.Error(err.Error())
	}
	if h.ID == 0 {
		t.Error("Hardware type ID equals zero!")
	}
	if reflect.TypeOf(h.Name).Kind() != reflect.String {
		t.Error("Name is not a string! Type of name is", reflect.TypeOf(h.Name).Kind())
	}
}

// Test function updateModel
func TestUpdateHardwareType(t *testing.T) {
	h, err := readOneHardwareTypeByName(cfg.DB.FileName, "laptop")
	if err != nil {
		t.Error(err.Error())
	}
	h.Name = "NotALaptop"
	res, err := updateHardwareType(cfg.DB.FileName, h)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No hardware types were updated!")
	}
}

// Test function deleteHardwareType
func TestDeleteHardwareType(t *testing.T) {
	ht := hardwareType{Name: "NotALaptop"}
	res, err := deleteHardwareType(cfg.DB.FileName, ht)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No hardware types were deleted!")
	}
}

// Test function wrieManufacturer
func TestWriteManufacturer(t *testing.T) {
	m := manufacturer{Name: "NoName", Logo: "noname.png"}
	res, err := writeManufacturer(cfg.DB.FileName, m)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No manufacturers were added!")
	}
}

// Test function readManufacturers
func TestReadManufacturers(t *testing.T) {
	vl, err := readManufacturers("./db/hardware.db")
	if err != nil {
		t.Error(err.Error())
	}
	if len(vl) == 0 {
		t.Error("Failed to read users! Make sure there're records in user table")
	}
	if reflect.TypeOf(vl[0].ID).Kind() != reflect.Int {
		t.Error("ID is not an Integer! Type of ID is", reflect.TypeOf(vl[0].ID).Kind())
	}
	if reflect.TypeOf(vl[0].Name).Kind() != reflect.String {
		t.Error("Name is not a string! Type of name is", reflect.TypeOf(vl[0].Name).Kind())
	}
	vl, err = readManufacturers("emptyTest.db")
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function updateManufacturer
func TestUpdateManufacturer(t *testing.T) {
	m := manufacturer{Name: "NoName", Logo: "nonamelogo.png"}
	res, err := updateManufacturer(cfg.DB.FileName, m)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No manufacturers were updated!")
	}
}

// Test function deleteManufacturer
func TestDeleteManufacturer(t *testing.T) {
	m := manufacturer{Name: "NoName", Logo: "noname.png"}
	res, err := deleteManufacturer(cfg.DB.FileName, m)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No manufacturers were deleted!")
	}
}

// Test function wrieModel
func TestWriteModel(t *testing.T) {
	m := model{Manufacturer: 1, Name: "ComputerCase", Type: 1}
	res, err := writeModel(cfg.DB.FileName, m)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No models were added!")
	}
}

// Test function readModels
func TestReadModels(t *testing.T) {
	ml, err := readModels(cfg.DB.FileName)
	if err != nil {
		t.Error(err.Error())
	}
	if len(ml) == 0 {
		t.Error("Failed to read models! Make sure there're records in model table")
	}
	for _, m := range ml {
		if reflect.TypeOf(m.ID).Kind() != reflect.Int {
			t.Error("ID is not an Integer! Type of ID is", reflect.TypeOf(m.ID).Kind())
		}
	}
	for _, m := range ml {
		if reflect.TypeOf(m.Name).Kind() != reflect.String {
			t.Error("Name is not a string! Type of name is", reflect.TypeOf(m.Name).Kind())
		}
	}
	ml, err = readModels("emptyTest.db")
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function readOneModelByName
func TestReadOneModelByName(t *testing.T) {
	m, err := readOneModelByName(cfg.DB.FileName, "ComputerCase")
	if err != nil {
		t.Error(err.Error())
	}
	if m.ID == 0 {
		t.Error("Model ID equals zero!")
	}
	if reflect.TypeOf(m.Name).Kind() != reflect.String {
		t.Error("Name is not a string! Type of name is", reflect.TypeOf(m.Name).Kind())
	}
}

// Test function updateModel
func TestUpdateModel(t *testing.T) {
	om := model{Manufacturer: 1, Name: "ComputerCase", Type: 1}
	m, err := readOneModelByName(cfg.DB.FileName, om.Name)
	if err != nil {
		t.Error(err.Error())
	}
	m.Name = "NotAComputerCase"
	res, err := updateModel(cfg.DB.FileName, m)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No models were updated!")
	}
}

// Test function deleteModel
func TestDeleteModel(t *testing.T) {
	om := model{Manufacturer: 1, Name: "NotAComputerCase", Type: 1}
	m, err := readOneModelByName(cfg.DB.FileName, om.Name)
	if err != nil {
		t.Error(err.Error())
	}
	res, err := deleteModel(cfg.DB.FileName, m)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No models were deleted!")
	}
}

// Test function readHardwareFullList
func TestReadHardwareFullList(t *testing.T) {
	vl, err := readHardwareFullList("./db/hardware.db")
	if err != nil {
		t.Error(err.Error())
	}
	if len(vl) == 0 {
		t.Error("Failed to read users! Make sure there're records in user table")
	}
	if reflect.TypeOf(vl[0].ID).Kind() != reflect.Int {
		t.Error("ID is not an Integer! Type of ID is", reflect.TypeOf(vl[0].ID).Kind())
	}
	if reflect.TypeOf(vl[0].Name).Kind() != reflect.String {
		t.Error("Name is not a string! Type of name is", reflect.TypeOf(vl[0].Name).Kind())
	}
	vl, err = readHardwareFullList("emptyTest.db")
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function wrieModel
func TestWriteHardware(t *testing.T) {
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
	res, err := writeHardware(cfg.DB.FileName, h)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No hardware were added!")
	}
}

// Test function readHardware
func TestReadHardware(t *testing.T) {
	hl, err := readHardware(cfg.DB.FileName)
	if err != nil {
		t.Error(err.Error())
	}
	if len(hl) == 0 {
		t.Error("Failed to read hardware! Make sure there're records in model table")
	}
	for _, h := range hl {
		if reflect.TypeOf(h.ID).Kind() != reflect.Int {
			t.Error("ID is not an Integer! Type of ID is", reflect.TypeOf(h.ID).Kind())
		}
	}
	for _, h := range hl {
		if reflect.TypeOf(h.Name).Kind() != reflect.String {
			t.Error("Name is not a string! Type of name is", reflect.TypeOf(h.Name).Kind())
		}
	}
	hl, err = readHardware("emptyTest.db")
	if err == nil {
		t.Error(err.Error())
	}
}

// Test function readOneHardwareByName
func TestReadOneHardwareByName(t *testing.T) {
	h, err := readOneHardwareByName(cfg.DB.FileName, "test-pc")
	if err != nil {
		t.Error(err.Error())
	}
	if h.ID == 0 {
		t.Error("Hardware ID equals zero!")
	}
	if reflect.TypeOf(h.Name).Kind() != reflect.String {
		t.Error("Name is not a string! Type of name is", reflect.TypeOf(h.Name).Kind())
	}
}

// Test function updateHardware
func TestUpdateHardware(t *testing.T) {
	h, err := readOneHardwareByName(cfg.DB.FileName, "test-pc")
	if err != nil {
		t.Error(err.Error())
	}
	h.Name = "NotATest-pc"
	res, err := updateHardware(cfg.DB.FileName, h)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No hardware entries were updated!")
	}
}

// Test function deleteModel
func TestDeleteHardware(t *testing.T) {
	h, err := readOneHardwareByName(cfg.DB.FileName, "NotATest-pc")
	if err != nil {
		t.Error(err.Error())
	}
	res, err := deleteHardware(cfg.DB.FileName, h)
	if err != nil {
		t.Error(err.Error())
	}
	if res == 0 {
		t.Error("No hardware entries were deleted!")
	}
}
