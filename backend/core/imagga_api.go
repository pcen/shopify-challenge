package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Imagga API endpoints
const (
	uploadURL = "https://api.imagga.com/v2/uploads"
	tagURL = "https://api.imagga.com/v2/tags"
)

// uploadID
// The unique ID used to specify an image when sending requests to the '/tag'
// endpoint.
type uploadID struct {
	ID string `json:"upload_id"`
}

// ImaggaUploadResponse
// Imagga API response for '/upload' requests
type uploadResponse struct {
	Result uploadID    `json:"result"`
	Status interface{} `json:"status"`
}

// UploadImage
// Uploads an image to imagga and returns the corresponding imagga ID for the
// uploaded image.
// Modified example code from https://docs.imagga.com/?go#uploads
func uploadImage(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	fi, err := file.Stat()
	if err != nil {
		return "", err
	}

	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", fi.Name())
	if err != nil {
		return "", err
	}

	part.Write(fileContents)

	client := &http.Client{}
	apiKey := os.Getenv("IMAGGA_API_KEY")
	apiSecret := os.Getenv("IMAGGA_API_SECRET")

	req, _ := http.NewRequest("POST", uploadURL, body)
	req.SetBasicAuth(apiKey, apiSecret)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New("Error when sending request to imagga API")
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	var result uploadResponse
	json.Unmarshal(respBody, &result)
	return result.Result.ID, nil
}

// imageTagLexeme JSON Model
// The lexical value of an image tag
type imageTagLexeme struct {
	Value string `json:"en"`
}

// ImageTag JSON Model
// Tag used to classify images
type ImageTag struct {
	Confidence float32        `json:"confidence"`
	Tag        imageTagLexeme `json:"tag"`
}

// tagList
// List of tags returned from the imagga '/tag' endpoint response
type tagList struct {
	Tags []ImageTag `json:"tags"`
}

// tagResponse
// Imagga API response from '/tag' endpoint requests
type tagResponse struct {
	Result tagList     `json:"result"`
	Status interface{} `json:"status"`
}

// TagImage
// Gets a list of tags for the specified image from imagga API.
// Modified example code from https://docs.imagga.com/?go#tags
func tagImage(uploadID string) ([]ImageTag, error) {
	client := &http.Client{}
	apiKey := os.Getenv("IMAGGA_API_KEY")
	apiSecret := os.Getenv("IMAGGA_API_SECRET")

	queryURL := fmt.Sprintf("%s?image_upload_id=%s", tagURL, uploadID)
	req, _ := http.NewRequest("GET", queryURL, nil)
	req.SetBasicAuth(apiKey, apiSecret)

	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Error when sending request to imagga API")
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	var result tagResponse
	err = json.Unmarshal(respBody, &result)
	if len(result.Result.Tags) <= 10 {
		return result.Result.Tags, err
	}
	return result.Result.Tags[:10], err
}

// GetImageTags gets the tags for the image at the given filepath and returns
// a list of tags and the confidence score for each tag. If an error is
// encountered, returns an empty ImageTag slice.
func GetImageTags(filepath string) ([]ImageTag, error) {
	uploadID, err := uploadImage(filepath)
	if err != nil {
		return []ImageTag{}, err
	}
	tags, err := tagImage(uploadID)
	if err != nil {
		return []ImageTag{}, err
	}
	return tags, nil
}
