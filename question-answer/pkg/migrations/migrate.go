package migrations

import (
	"log"

	"gorm.io/gorm"

	"question-answer/internal/question"
	"question-answer/internal/answer"
	"question-answer/internal/category"
)

// RunMigrations همه جداول را migrate می‌کند
func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&question.QuestionDB{},
		&answer.AnswerDB{},
		&category.CategoryDB{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("✅ Database migrations completed")
}
