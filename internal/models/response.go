package models

type SuccessResp struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ErrorResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
