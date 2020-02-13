package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Read from db all hardware items
func readHardware(filename string) ([]hardware, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM hardware")
	if err != nil {
		return nil, err
	}
	var hl []hardware
	var h hardware
	for rows.Next() {
		rows.Scan(&h.ID, &h.Name, &h.SerialNumber, &h.Manufacturer, &h.Model,
			&h.Type, &h.ProductionDate, &h.PurchaseDate, &h.Barcode,
			&h.CPUName, &h.CPUMaxClock, &h.CPUCores, &h.RAM, &h.Disk, &h.UserName)
		hl = append(hl, h)
	}
	return hl, nil
}

// Read certain hardware from database by name
func readOneHardwareByName(filename, name string) (hardware, error) {
	var h hardware
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return h, err
	}
	defer db.Close()
	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM hardware where name='%s'", name))
	if err != nil {
		return h, err
	}
	for rows.Next() {
		rows.Scan(&h.ID, &h.Name, &h.SerialNumber, &h.Manufacturer, &h.Model,
			&h.Type, &h.ProductionDate, &h.PurchaseDate, &h.Barcode,
			&h.CPUName, &h.CPUMaxClock, &h.CPUCores, &h.RAM, &h.Disk, &h.UserName)
	}
	return h, nil
}

// Write new hardware into database
func writeHardware(filename string, h hardware) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf(`INSERT
		INTO hardware
		(name,serial_number,manufacturer,model,type,production_date,purchase_date,barcode,cpu_name,cpu_max_clock,cpu_cores,ram,disk,user_name)
		VALUES('%s','%s',%v,%v,%v,'%s','%s','%s','%s',%v,%v,%v,%v,'%s')`,
			h.Name, h.SerialNumber, h.Manufacturer, h.Model, h.Type, h.ProductionDate,
			h.PurchaseDate, h.Barcode, h.CPUName, h.CPUMaxClock, h.CPUCores, h.RAM, h.Disk, h.UserName))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Delete an existing hardware entry from database
func deleteHardware(filename string, h hardware) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("DELETE FROM hardware WHERE id='%v'", h.ID))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Update an existing hardware entry from database
func updateHardware(filename string, h hardware) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf(`UPDATE hardware
		SET
		name = '%s', serial_number = '%s', manufacturer = %v, model = %v,
		type = %v, production_date = '%s', purchase_date = '%s', barcode = '%s',
		cpu_name = '%s', cpu_max_clock = %v, cpu_cores = %v, ram = %v, disk = %v,
		user_name = '%s'
		WHERE id = %v`,
			h.Name, h.SerialNumber, h.Manufacturer, h.Model, h.Type, h.ProductionDate,
			h.PurchaseDate, h.Barcode, h.CPUName, h.CPUMaxClock, h.CPUCores, h.RAM, h.Disk,
			h.UserName, h.ID))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Read from db available hardware types
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

// Read certain hardware type from database by name
func readOneHardwareTypeByName(filename, name string) (hardwareType, error) {
	var h hardwareType
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return h, err
	}
	defer db.Close()
	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM hardware_type WHERE name='%s'", name))
	if err != nil {
		return h, err
	}
	for rows.Next() {
		rows.Scan(&h.ID, &h.Name)
	}
	return h, nil
}

// Write new hardware type into database
func writeHardwareType(filename string, ht hardwareType) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("INSERT INTO hardware_type (name) VALUES('%s')", ht.Name))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Update an existing hardware type entry from database
func updateHardwareType(filename string, h hardwareType) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf(`UPDATE hardware_type SET name = '%s' WHERE id = %v`, h.Name, h.ID))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Delete an existing hardware type entry from database
func deleteHardwareType(filename string, ht hardwareType) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("DELETE FROM hardware_type WHERE name='%s'", ht.Name))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Read from db existing manufacturers
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

// Write new manufacturer into database
func writeManufacturer(filename string, m manufacturer) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("INSERT INTO manufacturer (name, logo) VALUES('%s','%s')", m.Name, m.Logo))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Update an existing manufacturer entry from database
func updateManufacturer(filename string, m manufacturer) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("UPDATE manufacturer SET name='%s', logo='%s' WHERE name='%s'", m.Name, m.Logo, m.Name))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Delete an existing manufacturer entry from database
func deleteManufacturer(filename string, m manufacturer) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("DELETE FROM manufacturer WHERE name='%s'", m.Name))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Write new model into database
func writeModel(filename string, m model) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("INSERT INTO model (manufacturer, name, type) VALUES(%v,'%s',%v )", m.Manufacturer, m.Name, m.Type))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Read certain model from database by name
func readOneModelByName(filename, modelname string) (model, error) {
	var m model
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return m, err
	}
	defer db.Close()
	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM model where name='%s'", modelname))
	if err != nil {
		return m, err
	}
	for rows.Next() {
		rows.Scan(&m.ID, &m.Manufacturer, &m.Name, &m.Type)
	}
	return m, nil
}

// Read from db existing models
func readModels(filename string) ([]model, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM model")
	if err != nil {
		return nil, err
	}
	var ml []model
	var m model
	for rows.Next() {
		rows.Scan(&m.ID, &m.Manufacturer, &m.Name, &m.Type)
		ml = append(ml, m)
	}
	return ml, nil
}

// Update an existing model entry from database
func updateModel(filename string, m model) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("UPDATE model SET manufacturer=%v, name='%s', type=%v WHERE id=%v", m.Manufacturer, m.Name, m.Type, m.ID))
	if err != nil {
		return 0, err
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(ra), nil
}

// Delete an existing model entry from database
func deleteModel(filename string, m model) (rowsAffected int, err error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return 0, nil
	}
	defer db.Close()
	req, err := db.Exec(
		fmt.Sprintf("DELETE FROM model WHERE ID='%v'", m.ID))
	if err != nil {
		return 0, nil
	}
	ra, err := req.RowsAffected()
	if err != nil {
		return 0, nil
	}
	return int(ra), nil
}

// Reads from db existing hardware list
func readHardwareFullList(filename string) ([]hardwareFull, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM hardware_full")
	if err != nil {
		return nil, err
	}
	var hl []hardwareFull
	var h hardwareFull
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
	if len(_user.Password) != 0 {
		// Make password hash to store it in database
		h, err := hashPassword(u.Password)
		if err != nil {
			return 0, err
		}
		_user.Password = h
	} else {
		_user.Password = u.Password
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
