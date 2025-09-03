package answer

import "time"

type AnswerDB struct {
	ID         string `gorm:"primaryKey"`
	QuestionID string
	Body       string
	UserID     string
	CreatedAt  int64
}

// NewAnswer helper برای ایجاد Answer
func NewAnswer(questionID, body, userID string) *AnswerDB {
	return &AnswerDB{
		ID:         generateID(),
		QuestionID: questionID,
		Body:       body,
		UserID:     userID,
		CreatedAt:  time.Now().Unix(),
	}
}

// simple uuid placeholder
func generateID() string {
	return "uuid-5678"
}
