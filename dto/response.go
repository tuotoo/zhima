package dto

type Response struct {
	Encrypted bool   `json:"encrypted"`
	Sign      string `json:"biz_response_sign"`
	Data      string `json:"biz_response"`
}
