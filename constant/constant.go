package constant

import (
	"encoding/json"
	"net/http"
)

var (
	HOST      string
	PORT      string
	USER      string
	PASS      string
	DBNAME    string
	JwtSecret []byte
)

func init() {
	HOST = "localhost"
	PORT = "5432"
	USER = "postgres"
	PASS = "rufus123"
	DBNAME = "school"

	JwtSecret = []byte("monkey")
}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}
