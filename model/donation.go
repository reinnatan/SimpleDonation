package model

type Donation struct {
	Id          string
	Description string `json:"description" xml:"description" form:"description"`
	Total       int64  `json:"total" xml:"total" form:"total"`
	DueDate     int64  `json:"dueDate" xml:"dueDate" form:"dueDate"`
	Type        string `json:"type" xml:"type" form:"type"`
}

type GeneralResponse struct {
	Message      string
	ResponseCode int64
}
