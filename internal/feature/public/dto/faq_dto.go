package dto

import "time"

type FAQQueryRequest struct {
	Lang string `form:"lang"`
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
	ID            string        `json:"id"`
	SubTopics     []FAQSubTopic `json:"subTopics"`
	LastUpdatedAt *time.Time    `json:"last_updated_at"`
}
