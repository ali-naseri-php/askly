package question


import "time"

// QuestionDB مدل دیتابیس برای سوال
type QuestionDB struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Body      string
	UserID    string // استخراج شده از توکن
	CreatedAt int64
}

// helper برای تبدیل زمان
func NewQuestion(title, body, userID string) *QuestionDB {
	return &QuestionDB{
		ID:        generateID(),
		Title:     title,
		Body:      body,
		UserID:    userID,
		CreatedAt: time.Now().Unix(),
	}
}

// simple uuid placeholder
func generateID() string {
	return "uuid-1234"
}
