package model

type Content struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Status    string `json:"status"` // pending, approved, rejected
	ReviewedBy string `json:"reviewed_by,omitempty"`
	Comment    string `json:"comment,omitempty"`
}
