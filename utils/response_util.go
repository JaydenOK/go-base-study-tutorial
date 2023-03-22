package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	StatusSuccess = 1
	StatusFail    = 2
)

type ResponseData struct {
	status  int
	message string
	data    interface{}
}

func JsonResponse(resp http.ResponseWriter, httpCode, status int, message string, data interface{}) {
	responseData := ResponseData{
		status:  status,
		message: message,
		data:    data,
	}
	bytes, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println("json.Marshal(user) err=", err)
	}
	//响应返回客户端json数据
	resp.WriteHeader(httpCode)
	resp.Header().Set("Content-Type", "application/json")
	if i, err := resp.Write(bytes); err != nil {
		fmt.Println("JsonResponse Fail:", i)
	}
}

func SuccessResponse(resp http.ResponseWriter, data interface{}) {
	JsonResponse(resp, http.StatusOK, StatusSuccess, "success", data)
}

func FailResponse(resp http.ResponseWriter, data interface{}) {
	JsonResponse(resp, http.StatusOK, StatusFail, "fail", data)
}

func MessageResponse(resp http.ResponseWriter, message string) {
	JsonResponse(resp, http.StatusOK, StatusSuccess, "success", message)
}
