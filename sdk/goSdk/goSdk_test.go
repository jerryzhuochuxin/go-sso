package goSdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetResponseUsePingApi(t *testing.T) {
	SetRegisterIp("http://127.0.0.1:8081")
	res, _ := GetResponse("GET", "jwtService/api/ping", nil, nil)
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
}
func TestGetResponseUseGetTokenApi(t *testing.T) {
	SetRegisterIp("http://127.0.0.1:8081")
	res, _ := GetResponse("POST", "jwtService/api/getToken", strings.NewReader(`{"name":"jerryzhuo"}`), nil)
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
}
func TestGetResponseUseAuthTokenApi(t *testing.T) {
	SetRegisterIp("http://127.0.0.1:8081")
	res, _ := GetResponse("POST", "jwtService/api/authToken", nil, func(req *http.Request) {
		req.Header.Set("X-TokenValue", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjoiZXlKdVlXMWxJam9pYW1WeWNubDZhSFZ2SW4wPSIsImV4cCI6MTU5NDYzNDEzNn0.IT9MTbEcwe8h3fbmRXrjT-yyFwmfLUx5lnhE6pegcWo")
	})

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
}
