package dto

type FAQQueryReq struct {
	Lang     string `form:"lang"`
	Category string `form:"category"`
	RoleType string `form:"role_type"`
}

type FAQItem struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type FAQSubTopic struct {
	Title string    `json:"title"`
	Items []FAQItem `json:"items"`
}

type FAQCategoryResponse struct {
	ID          string        `json:"id"`
	Label       string        `json:"label"`
	Description string        `json:"description"`
	SubTopics   []FAQSubTopic `json:"subTopics"`
}
