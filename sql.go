package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Reads from db available hardware types
func readHardwareTypes(filename string) ([]hardwareType, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM hardware_type")
	if err != nil {
		return nil, err
	}
	var hl []hardwareType
	var h hardwareType
	for rows.Next() {
		rows.Scan(&h.ID, &h.Name)
		hl = append(hl, h)
	}
	return hl, nil
}

// Reads from db existing manufacturers
func readManufacturers(filename string) ([]manufacturer, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM manufacturer")
	if err != nil {
		return nil, err
	}
	var ml []manufacturer
	var m manufacturer
	for rows.Next() {
		rows.Scan(&m.ID, &m.Name, &m.Logo)
		ml = append(ml, m)
	}
	return ml, nil
}

// Reads from db existing hardware list
func readHardwareFullList(filename string) ([]hardware, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM hardware_full")
	if err != nil {
		return nil, err
	}
	var hl []hardware
	var h hardware
	for rows.Next() {
		rows.Scan(&h.ID, &h.Name, &h.UserName, &h.Manufacturer, &h.Model, &h.Type,
			&h.ProductionDate, &h.PurchaseDate, &h.Barcode, &h.CPUName,
			&h.CPUMaxClock, &h.CPUCores, &h.RAM, &h.Disk)
		hl = append(hl, h)
	}
	return hl, nil
}

// Write user into database
func addUser(filename string, u user) (id int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	// Make password hash to store it in database
	h, err := hashPassword(u.Password)
	if err != nil {
		return 0, err
	}
	u.Password = h
	res, err := db.Exec(
		fmt.Sprintf(
			`INSERT
			INTO user
				(login, password, first_name, last_name, group_id)
			VALUES
				("%s", "%s", "%s", "%s", "%s")`,
			u.Login, u.Password, u.FirstName, u.LastName, u.Group))
	if err != nil {
		return 0, err
	}
	ra, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(ra), nil
}

// Update existing user entry
func updateUser(filename string, _user user) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	u, err := readOneUserByLogin(filename, _user.Login)
	if err != nil {
		return 0, err
	}
	// Check if new password passed
	if len(u.Password) == 0 {
		// Make password hash to store it in database
		h, err := hashPassword(u.Password)
		if err != nil {
			return 0, err
		}
		u.Password = h
	}
	res, err := db.Exec(
		fmt.Sprintf("UPDATE user SET login='%s', password='%s', first_name='%s', last_name='%s' WHERE ID=%v",
			_user.Login, _user.Password, _user.FirstName, _user.LastName, u.ID))
	if err != nil {
		return 0, err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(ra), nil
}

// Delete existing user entry
func deleteUser(filename, login string) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	res, err := db.Exec(
		fmt.Sprintf("DELETE FROM user WHERE login = '%s'", login))
	if err != nil {
		return 0, err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(ra), nil
}

// Read all users from database
func readUsers(filename string) ([]user, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	var ul []user
	var u user
	for rows.Next() {
		rows.Scan(&u.ID, &u.Login, &u.Password, &u.FirstName, &u.LastName, &u.Group)
		ul = append(ul, u)
	}
	return ul, nil
}

// Read certain user from database by id
func readOneUserByID(filename string, id int) (user, error) {
	var u user
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return u, err
	}
	defer db.Close()
	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM user where id=%v", id))
	if err != nil {
		return u, err
	}
	for rows.Next() {
		rows.Scan(&u.ID, &u.Login, &u.Password, &u.FirstName, &u.LastName, &u.Group)
	}
	return u, nil
}

// Read certain user from database by login
func readOneUserByLogin(filename, login string) (user, error) {
	var u user
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return u, err
	}
	defer db.Close()
	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM user where login='%s'", login))
	if err != nil {
		return u, err
	}
	for rows.Next() {
		rows.Scan(&u.ID, &u.Login, &u.Password, &u.FirstName, &u.LastName, &u.Group)
	}
	return u, nil
}

// Check if entered password matches with saved in database
func checkPasswordFromDB(filename, login, password string) (bool, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return false, err
	}
	defer db.Close()
	rows, err := db.Query(
		fmt.Sprintf(`SELECT password FROM user WHERE login = "%s"`, login))
	if err != nil {
		return false, err
	}
	var p string
	for rows.Next() {
		rows.Scan(&p)
	}
	return checkHashMatch(password, p), nil
}
