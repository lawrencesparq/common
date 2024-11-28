package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type PocketBaseResponse struct {
	CollectionID   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Created        string `json:"created"`
	ID             string `json:"id"`
	Invitation     string `json:"invitation"`
	Updated        string `json:"updated"`
}

func UploadFileToPocketBase(baseURL, collection, field, filename, filePath string) (*PocketBaseResponse, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(field, filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	request, err := http.NewRequest("POST", baseURL+"/api/collections/"+collection+"/records", body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	// Handle the server response...
	resp, err := io.ReadAll(response.Body) // response body is []byte
	if err != nil {
		log.Println(err)
	}
	var pocketBaseResponse PocketBaseResponse
	// Parse the JSON data into the struct
	err = json.Unmarshal(resp, &pocketBaseResponse)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	log.Printf("response: %v", pocketBaseResponse)

	return &pocketBaseResponse, nil
}

type DeletedRecordResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
	} `json:"data"`
}

func DeletePocketBaseRecord(baseURL, collection, recordID string) (*DeletedRecordResponse, error) {
	request, err := http.NewRequest("DELETE", baseURL+"/api/collections/"+collection+"/records/"+recordID, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	// Handle the server response...
	resp, err := io.ReadAll(response.Body) // response body is []byte
	if err != nil {
		log.Println(err)
	}

	log.Println(string(resp))
	var deletedRecordResponse DeletedRecordResponse
	// Parse the JSON data into the struct
	err = json.Unmarshal(resp, &deletedRecordResponse)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	log.Printf("response: %v", deletedRecordResponse)

	return &deletedRecordResponse, nil
}

func MakeFileUrl(baseURl, collectionId, id, filename string) string {
	return baseURl + "/api/files/" + collectionId + "/" + id + "/" + filename
}
