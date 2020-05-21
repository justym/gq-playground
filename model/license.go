package model

type Response struct {
	License struct {
		Name        string `json: "name"`
		Description string `json: "description"`
	} `json: "license"`
}
