package main

//ResponseSuccess struct
type ResponseSuccess struct {
	Data        []Data `json:"data"`
	ErrorString string `json:"error_string"`
	Timestamp   int    `json:"timestamp"`
}

//Data struct Status,Error,SmsLog,Queue,To
type Data struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	SmsLog string `json:"smslog_id"`
	Queue  string `json:"queue"`
	To     string `json:"to"`
}

//ResponseError struct
type ResponseError struct {
	Status      string `json:"status"`
	Error       string `json:"error"`
	ErrorString string `json:"error_string"`
	Timestamp   int    `json:"timestamp"`
}

//ResponseSuccessProfil struct
type ResponseSuccessProfil struct {
	Data        []DataProfil `json:"data"`
	ErrorString string       `json:"error_string"`
	Timestamp   int          `json:"timestamp"`
}

//DataProfil struct Status,Error,SmsLog,Queue,To
type DataProfil struct {
	SmsLog      string `json:"smslog_id"`
	Source      string `json:"src"`
	Destination string `json:"dst"`
	Message     string `json:"msg"`
	Date        string `json:"dt"`
	Update      string `json:"update"`
	Status      string `json:"status"`
}
