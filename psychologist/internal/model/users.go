package model

//Client presents client
type Client struct {
	ID         string `json:"id,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Name       string `json:"name,omitempty"`
	Patronomic string `json:"patronomic,omitempty"`
}