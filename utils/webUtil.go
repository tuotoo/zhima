package utils

import (
	"bytes"
	"encoding/json"
	"github.com/delostik/go-zmxy/dto"
	"github.com/delostik/go-zmxy/interfaces"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func BuildQueryString(params map[string]string, needEncode bool) string {
	paramsList := []string{}
	for key, value := range params {
		if !IsEmpty(value) {
			if needEncode {
				paramsList = append(paramsList, key+"="+url.QueryEscape(value))
			} else {
				paramsList = append(paramsList, key+"="+value)
			}
		}
	}
	return strings.Join(paramsList, "&")
}

func PostRequest(url string, request *interfaces.ZhimaRequest, httpParams *map[string]string) (*dto.Response, error) {
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)

	// Put parameters into post form
	for key, val := range *httpParams {
		_ = writer.WriteField(key, val)
	}

	// Put files into post form (if exist)
	files := (*request).GetFileParams()
	if len(files) > 0 {
		for key, value := range files {
			part, err := writer.CreateFormFile(key, path.Base(value))
			if err != nil {
				return nil, err
			}
			file, err := os.Open(value)
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return nil, err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	// Do http request
	req, err := http.NewRequest("POST", url, reqBody)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var respData dto.Response
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}
