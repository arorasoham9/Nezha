package ssh

type Queue_Request struct {
	NAME       string `json:"name"`
	EMAIL      string `json:"email"`
	CURRENT_IP string `json:"current_ip"`
	LOCATION   string `json:"location"`
	CREATED_AT string `json:"created_at"`
	LASTSEEN   string `json:"last_seen"`
	PORT       string `json:"port"`
	KEY        string `json:"key"`
}
