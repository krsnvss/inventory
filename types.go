package main

// Application configuration
type appConfig struct {
	DB struct {
		FileName string `json:"file_name"`
	} `json:"db"`
	Token struct {
		IdentityKey string `json:"identity_key"`
	} `json:"token"`
}

// Hardware type
type hardwareType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Hardware types list. API response
type hardwareTypeList struct {
	Status string         `json:"status"`
	Result []hardwareType `json:"result,omitempty"`
	Error  string         `json:"error,omitempty"`
}

// Hardwware item
type hardware struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	SerialNumber   string  `json:"serial_number"`
	Manufacturer   string  `json:"manufacturer"`
	Model          string  `json:"model"`
	Type           string  `json:"type"`
	ProductionDate string  `json:"production_date,omitempty"`
	PurchaseDate   string  `json:"purchase_date,omitempty"`
	Barcode        string  `json:"barcode,omitempty"`
	CPUName        string  `json:"cpu_name,omitempty"`
	CPUMaxClock    float64 `json:"cpu_max_clock,omitempty"`
	CPUCores       int     `json:"cpu_cores,omitempty"`
	RAM            int     `json:"ram,omitempty"`
	Disk           int     `json:"disk,omitempty"`
	UserName       string  `json:"user_name,omitempty"`
}

// Manufacturer
type manufacturer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo,omitempty"`
}

// Model
type model struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
}

// Repair
type repair struct {
	ID         int    `json:"id"`
	Hardware   int    `json:"hardware"`
	RepairDate string `json:"repair_date"`
	Comment    string `json:"comment,omitempty"`
}

// Decomission
type decomission struct {
	ID              int    `json:"id"`
	Hardware        int    `json:"hardware"`
	DecomissionDate string `json:"decomission_date"`
	Comment         string `json:"comment,omitempty"`
}
