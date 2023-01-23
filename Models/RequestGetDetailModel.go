package Models

type RequestGetDetails struct {
	CIF           string `json:"cif"`
	TransactionID string `json:"transactionID"`
}

type GetDataFromDB struct {
	CIF string
}

type InputPostman struct {
	MobileNumber string `json:"mobileNumber"`
}
