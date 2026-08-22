package dto

import "encoding/json"

type LegalDocumentQueryReq struct {
	DocType string `form:"doc_type" binding:"required"`
	Lang    string `form:"lang"`
}

type LegalDocumentResponse struct {
	DocType  string          `json:"doc_type"`
	Label    string          `json:"label"`
	Title    string          `json:"title"`
	Intro    string          `json:"intro"`
	Sections json.RawMessage `json:"sections"`
}
