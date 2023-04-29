package ssh

import "encoding/json"


func parseRequestJSON(rqst []byte)(Queue_Request, error){
	var request Queue_Request
	err := json.Unmarshal([]byte(rqst), &request)
	if err != nil {
		return Queue_Request{}, err
	}
	return request, nil
}

func SendOutConnection(rqst []byte) error{
	q_rqst, err := parseRequestJSON(rqst)
	if err != nil{
		return err
	}
	q_rqst.





	return nil
}