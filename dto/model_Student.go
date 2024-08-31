package dto

type Student struct {
	StudentId   	 string `json:"StudentId" `
	Name        	 string ` json:"Name"`
	Email            string `json:"Email" `
	Country          string `json:"Country" `
	Phone            string `json:"Phone" `
	Address          string `json:"Address" `
	Gender           string `json:"Gender" `
	DateOfBirth      string `json:"DateOfBirth" `
	Deleted          bool   `json:"deleted"`
}