package model

// Response is a type for http response.
type Response struct {
	Config  Config `json:"Config"`
	Command uint64 `json:"Command"`
}
