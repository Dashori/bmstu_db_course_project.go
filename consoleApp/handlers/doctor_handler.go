package handlers

import (
	"bytes"
	errors "consoleApp/errors"
	models "consoleApp/models"
	"fmt"
	"net/http"
)

const port = "8080"
const adress = "localhost"

func DoRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.ErrorExecuteRequest
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return response, errors.ErrorResponseStatus
	}

	return response, nil
}

func GetDoctors(client *http.Client) (*http.Response, error) {
	url := "http://" + adress + ":" + port + "/api/doctors"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func LoginDoctor(client *http.Client, doctor *models.Doctor) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/doctor/login"
	params := fmt.Sprintf("{\"Login\": \"%s\", \"Password\": \"%s\"}", doctor.Login, doctor.Password)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorHTTP
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func GetInfo(client *http.Client, token string) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/doctor/info"

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func UpdateShedule(client *http.Client, token string, shedule models.NewShedule) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/doctor/shedule"
	params := fmt.Sprintf("{\"StartTime\": %d, \"EndTime\": %d}", shedule.Start, shedule.End)

	var jsonStr = []byte(params)
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func GetRecordsDoctor(client *http.Client, token string) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/doctor/records"

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}
