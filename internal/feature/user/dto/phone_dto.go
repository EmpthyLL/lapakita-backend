package dto

type PhoneNumberItem struct {
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}

type GetPhoneNumbersResponse struct {
	PhoneNumbers []PhoneNumberItem `json:"phone_numbers"`
}

type AddPhoneNumberRequest struct {
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}

type UpdatePhoneNumberRequest struct {
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}
