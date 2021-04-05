package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arh0329/supermarket-api/models"
	"github.com/stretchr/testify/assert"
)

// health
func Test_HealthEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r := CreateRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, string(body), `{"message":"ok"}`)
}

// getAllProduce
func Test_getAllProduce_Success(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/produce", nil)
	r := CreateRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := []models.Item{}
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		t.Errorf("An error occurred: %v", err.Error())
	}
	if !models.Equal(res, models.Produce) {
		t.Errorf("Actual does not equal expected.\nActual:%v\nExpected:%v", models.Produce, res)
	}
}

//addProduce
func Test_addProduce_Success(t *testing.T) {
	w := httptest.NewRecorder()
	newItem := []models.Item{
		{
			Name:        "pear",
			ProduceCode: "E5T6-9UI3-TH15-QR88",
			UnitPrice:   2.99,
		},
	}
	marshaled, _ := json.Marshal(newItem)

	req, _ := http.NewRequest("POST", "/produce", bytes.NewBuffer(marshaled))
	r := CreateRouter()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, `{"added":["pear"],"errors":[],"message":"Item(s) added"}`, string(body))
}

func Test_addProduce_BadRequest(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/produce", bytes.NewBuffer([]byte("not json")))
	r := CreateRouter()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, `{"message":"invalid character 'o' in literal null (expecting 'u')"}`, string(body))
}

func Test_addProduce_InvalidItem(t *testing.T) {
	w := httptest.NewRecorder()
	newItem := []models.Item{
		{
			Name:        "pe$ar",
			ProduceCode: "E5T6-9UI3-TH15-QR88",
			UnitPrice:   2.99,
		},
	}
	marshaled, _ := json.Marshal(newItem)

	req, _ := http.NewRequest("POST", "/produce", bytes.NewBuffer(marshaled))
	r := CreateRouter()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, `{"added":[],"errors":["Key: 'Item.Name' Error:Field validation for 'Name' failed on the 'alphanum' tag"],"message":"No items added"}`, string(body))
}

// getOneItem
var pc = "E5T6-9UI3-TH15-QR88"

func Test_getOneItem_Success(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/produce/"+pc, nil)
	r := CreateRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := models.Item{}
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		t.Errorf("An error occurred: %v", err.Error())
	}

	assert.Equal(t, models.Produce[1], res)
}

func Test_getOneItem_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/produce/"+"loeh-fdsa-jfda-fdsa", nil)
	r := CreateRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, `{"message":"Item not found"}`, string(body))
}

// deleteItem
func Test_deleteItem_Success(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/produce/"+pc, nil)
	r := CreateRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, `{"message":"Item deleted"}`, string(body))
}

func Test_deleteItem_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/produce/"+"loeh-fdsa-jfda-fdsa", nil)
	r := CreateRouter()
	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
	body, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, `{"message":"Item not found"}`, string(body))
}
