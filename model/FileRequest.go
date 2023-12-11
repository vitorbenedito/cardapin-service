package model

type FileRequest struct {
	Extension   string `json:"extension"`
	CompanyName string `json:"companyName"`
	FileName    string `json:"fileName"`
}
