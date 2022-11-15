package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esmira23/go-postgresql-docker/config"
	"github.com/esmira23/go-postgresql-docker/csvparser"
	"github.com/esmira23/go-postgresql-docker/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostAll(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	req, err := http.NewRequest("POST", "/api/post_data", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetAll(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	csvdata := csvparser.GetCSVData()
	byteData, _ := json.Marshal(csvdata)

	req, err := http.NewRequest("GET", "/api/all", bytes.NewBuffer(byteData))

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	json.Unmarshal(w.Body.Bytes(), &csvdata)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, csvdata)

}

func TestGetByTransactionID(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	req, err := http.NewRequest("GET", "/api/transaction/3", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetByTerminalID(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	req, err := http.NewRequest("GET", "/api/terminal", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	q := req.URL.Query()
	q.Add("id", "3507,3512")
	req.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetByStatus(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	req, err := http.NewRequest("GET", "/api/status/accepted", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetByPaymentType(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	req, err := http.NewRequest("GET", "/api/payment_type/card", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetByDate(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	req, err := http.NewRequest("GET", "/api/date", nil)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	q := req.URL.Query()
	q.Add("from", "2022-08-17")
	q.Add("to", "2022-08-23")
	req.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetPaymentNarrative(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := routers.SetUpRouter()
	w := httptest.NewRecorder()
	config.ConnectDB()

	req, err := http.NewRequest("GET", "/api/payment_narrative/27122", nil)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
