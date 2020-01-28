package main

import (
	"reflect"
	"testing"
)

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
		LastName:  "Pupkiv",
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
