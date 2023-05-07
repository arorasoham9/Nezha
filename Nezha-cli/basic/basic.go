package basic

import (
	// "bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"

	"golang.org/x/term"
)
var HTTPSCLIENT *http.Client

func InitRest(){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},	
	}
	HTTPSCLIENT = &http.Client{Transport: tr}
}

func SendGET(GETURL string) (*http.Response, error){
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
func SendLoginRequest( newUserLoginRequest LoginUser, HOSTPORT string) (LoginResponse, bool){
	var loginResp LoginResponse
	resp, err := SendPOST(ROUTE_LOGIN, newUserLoginRequest)
	if err != nil {
		loginResp.MESSAGE = "Request could not be sent." 
		return loginResp, false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loginResp.MESSAGE = "Could not read response."
		return loginResp, false
	}
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		loginResp.MESSAGE = "Could not parse response."
		return loginResp, false
	}
	if !storeLoginInfo(newUserLoginRequest, loginResp){
		loginResp.MESSAGE = "Could not store login info."
		return loginResp, false
	}
	return loginResp, true
}
func SaveLoginResult(Request LoginUser, response LoginResponse, resp *http.Response, email string) {
	// var obj t.LogSuccess
	
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(resp.StatusCode, "Could not read response.")
	// 	os.Exit(1)
	// }
	
	// json.Unmarshal(body, &obj)
	// // fmt.Println(obj)
	// if resp.StatusCode != http.StatusOK{
	// 	fmt.Println(obj.MESSAGE)
	// 	return
	// }
	// switch obj.TOKEN {
	// case "":
	// 	return
	// default:
	// 	os.MkdirAll("./config/", os.ModePerm)
	// 	fo, err := os.Create("./config/token-config.json")
	// 	if err != nil {
	// 		fmt.Println("Service unavailable.")
			
	// 	}
	// 	fo.Close()
	// 	file, _ := json.MarshalIndent(t.LoggedUser{
	// 		TOKEN: obj.TOKEN,
	// 		EMAIL: base64.StdEncoding.EncodeToString([]byte(email)),
	// 	}, "", " ")
	// 	_ = ioutil.WriteFile("./config/token-config.json", file, 0644)
	// }
}
func storeLoginInfo(Request LoginUser, response LoginResponse) bool{
 return true
}

func readLoginInfo(){

}
func LoginStatus() (LoggedUser, bool){
	var user LoggedUser
	// read

	return user, true
}

func GetCred() string{
	var  email string
	fmt.Printf("Enter username: ")
	fmt.Scanf("%s\n", &email)
	return email
}
func GetCredentials() (string, []byte) {
	
	EMAIL := GetCred()
	fmt.Printf("Enter password: ")
	tmpPass, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Cannot read from terminal. Try again later.")
		os.Exit(1)
	}
	fmt.Printf("\n.....\n")
	return EMAIL, tmpPass

}