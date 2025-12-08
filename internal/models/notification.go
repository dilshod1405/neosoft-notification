package models

type Notification struct {
	ID        int64       `json:"id"`
	UserID    int64       `json:"user_id"`
	Type      string      `json:"type"`
	Title     string      `json:"title"`
	Message   string      `json:"message"`
	Metadata  interface{} `json:"metadata"`
	ActionURL string      `json:"action_url"`
	CreatedAt string      `json:"created_at"`
}
