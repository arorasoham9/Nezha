package basic


import (
	
)

const (
	DOMAIN  = ""
	ROUTE_LOGIN = "/login"
	ROUTE_LIST = "/list"
	ROUTE_CONNECT = "/connect"
	LOGIN_INFO_MOUNT = "/Nezha/config.json"
)

type LoginResponse struct {
	TOKEN   string `json:"token"`
	MESSAGE string `json:"message"`
}

type LoggedUser struct {
	USERNAME string `json:"username"`
	LOGGED_AT 		string `json:"logged_at"`
	TOKEN string `json:"token"`
}

type LoginUser struct {
	USERNAME    string `json:"username"`
	CREATED_AT 		string `json:"created_at"`
	PASSWORD []byte `json:"pwd"`
}