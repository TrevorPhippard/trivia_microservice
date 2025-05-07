package model

type Quiz struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
	// CreatedAt   string `json:"createdAt"`
	// UpdatedAt   string `json:"updatedAt"`
	// DeletedAt   *string `json:"deletedAt"`
}

type Question struct {
	ID       string   `json:"id"`
	QuizID   string   `json:"quizId"`
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Options  []string `json:"options"`
	// CreatedAt   string `json:"createdAt"`
	// UpdatedAt   string `json:"updatedAt"`
	// DeletedAt   *string `json:"deletedAt"`
}
