package handlers

import (
	"bytes"
	errors "consoleApp/errors"
	models "consoleApp/models"
	"fmt"
	"net/http"
)

func SetRole(client *http.Client, role string) (*http.Response, error) {
	url := "http://" + adress + ":" + port + "/api/setRole"
	params := fmt.Sprintf("{\"Role\": \"%s\"}", role)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func LoginClient(client *http.Client, newClient *models.Client) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/client/login"
	params := fmt.Sprintf("{\"Login\": \"%s\", \"Password\": \"%s\"}", newClient.Login, newClient.Password)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func CreateClient(client *http.Client, newClient *models.Client) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/client/create"
	params := fmt.Sprintf("{\"Login\": \"%s\", \"Password\": \"%s\"}", newClient.Login, newClient.Password)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func GetClientPets(client *http.Client, token string) (*http.Response, error) {
	url := "http://" + adress + ":" + port + "/api/client/pets"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func GetClientInfo(client *http.Client, token string) (*http.Response, error) {
	url := "http://" + adress + ":" + port + "/api/client/info"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func GetClientRecords(client *http.Client, token string) (*http.Response, error) {
	url := "http://" + adress + ":" + port + "/api/client/records"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func AddPet(client *http.Client, token string, pet models.Pet) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/client/pet"
	params := fmt.Sprintf("{\"Name\": \"%s\", \"Type\": \"%s\", \"Age\": %d,\"Health\": %d}",
		pet.Name, pet.Type, pet.Age, pet.Health)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func DeletePet(client *http.Client, token string, id uint64) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/client/pet"
	params := fmt.Sprintf("{\"PetId\": %d}", id)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}

func AddRecord(client *http.Client, token string, record models.Record) (*http.Response, error) {

	url := "http://" + adress + ":" + port + "/api/client/record"
	params := fmt.Sprintf("{\"PetId\": %d, \"DoctorId\": %d, \"Year\": %d,\"Month\": %d,\"Day\": %d, \"StartTime\": %d}",
		record.PetId, record.DoctorId, record.Year, record.Month, record.Day, record.DatetimeStart)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := DoRequest(client, request)

	return response, err
}
