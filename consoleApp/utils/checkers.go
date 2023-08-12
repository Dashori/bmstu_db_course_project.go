package utils

import (
	errors "consoleApp/errors"
	models "consoleApp/models"
	"encoding/json"
	basicErrors "errors"
	"io"
	"net/http"
)

func CheckErrorInBody(response *http.Response) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.ErrorReadBody
	}

	var result models.ErrorBody
	if err := json.Unmarshal(body, &result); err != nil {
		return errors.ErrorParseBody
	}

	return basicErrors.New(result.Err)
}

func DoAndCheckRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return &http.Response{}, CheckErrorInBody(response)
	}

	return response, nil
}

func ParseClientBody(response *http.Response) (models.Client, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Client{}, errors.ErrorReadBody
	}

	var result models.Client
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Client{}, errors.ErrorParseBody
	}

	return result, nil
}

func ParsePetsBody(response *http.Response) (models.Pets, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Pets{}, errors.ErrorReadBody
	}

	var result models.Pets
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Pets{}, errors.ErrorParseBody
	}

	return result, nil
}

func ParseDoctorBody(response *http.Response) (models.Doctor, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Doctor{}, errors.ErrorParseBody
	}

	var result models.Doctor
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Doctor{}, errors.ErrorParseBody
	}

	return result, nil
}

func ParseDoctorsBody(response *http.Response) (models.Doctors, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Doctors{}, errors.ErrorParseBody
	}

	var result models.Doctors
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Doctors{}, errors.ErrorParseBody
	}

	return result, nil
}

func ParseRecordsBody(response *http.Response) (models.Records, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Records{}, errors.ErrorParseBody
	}

	var result models.Records
	if err := json.Unmarshal(body, &result); err != nil {
		return models.Records{}, errors.ErrorParseBody
	}

	return result, nil
}
