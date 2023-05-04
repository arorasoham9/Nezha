package basic

import (
	// "bufio"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"

	// c "goshelly-client/cmd"
	t "goshelly-client/template"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/mail"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)
var HTTPSCLIENT *http.Client

func InitRest(){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},	
	}
	HTTPSCLIENT = &http.Client{Transport: tr}
}

func sendGET(GETURL string) (*http.Response, error){
	resp, err := HTTPSCLIENT.Get(GETURL)
	if err != nil {
		fmt.Println("Service offline.")
		return &http.Response{}, err
	}
	return resp, nil
}
func SendPOST(POSTURL string, user interface{}) (*http.Response, error) {
	body, _ := json.Marshal(user)
	resp, err := HTTPSCLIENT.Post(POSTURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Service offline.")
		return &http.Response{}, err
	}
	return resp, nil
}

func SendDEL(DELURL string, user interface{})(*http.Response, error){
	body, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodDelete, DELURL, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("HTTP Error.")
		return &http.Response{}, err
	}
	
	resp, err := HTTPSCLIENT.Do(req)
	if err != nil {
		fmt.Println("Service offline.")
		return &http.Response{}, err
	}

	defer resp.Body.Close()

	return resp, nil
}

