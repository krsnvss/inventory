package main

// Application configuration
type appConfig struct {
	DB struct {
		FileName string `json:"file_name"`
	} `json:"db"`
	Token struct {
		IdentityKey string `json:"identity_key"`
	} `json:"token"`
	DebugMode bool `json:"debug_mode"`
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

// Hardware item short representation
type hardware struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	SerialNumber   string  `json:"serial_number"`
	Manufacturer   int     `json:"manufacturer"`
	Model          int     `json:"model"`
	Type           int     `json:"type"`
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

// Hardware item full representation
type hardwareFull struct {
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

// Hardware list. API response
type hardwareList struct {
	Status string     `json:"status"`
	Result []hardware `json:"result,omitempty"`
	Error  string     `json:"error,omitempty"`
}

// Hardware full list. API response
type hardwareFullList struct {
	Status string         `json:"status"`
	Result []hardwareFull `json:"result,omitempty"`
	Error  string         `json:"error,omitempty"`
}

// Manufacturer
type manufacturer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo,omitempty"`
}

// Manufacturers list. API response
type manufacturersList struct {
	Status string         `json:"status"`
	Result []manufacturer `json:"result,omitempty"`
	Error  string         `json:"error,omitempty"`
}

// Model
type model struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Manufacturer int    `json:"manufacturer"`
	Type         int    `json:"type"`
}

// Models list. API response
type modelsList struct {
	Status string  `json:"status"`
	Result []model `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
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

// Server response for API
type serverResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Result int    `json:"result,omitempty"`
}
