package models

type TestRequest struct {
	ID       string `json:"int"`
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
}
