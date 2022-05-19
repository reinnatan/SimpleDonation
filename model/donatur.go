package model

type Donatur struct {
	Id         string `json:"id" xml:"id" form:"id"`
	Name       string `json:"name" xml:"name" form:"name"`
	DueDate    int64  `json:"dueDate" xml:"dueDate" form:"dueDate"`
	IdDonation string `json:"idDonation" xml:"idDonation" form:"idDonation"`
	Type       string `json:"type" xml:"type" form:"type"`
}
