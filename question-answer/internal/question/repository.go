package question

import (
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(q *QuestionDB) error {
	return r.db.Create(q).Error
}

func (r *QuestionRepository) GetByID(id string) (*QuestionDB, error) {
	var q QuestionDB
	err := r.db.First(&q, "id = ?", id).Error
	return &q, err
}

func (r *QuestionRepository) List() ([]*QuestionDB, error) {
	var questions []*QuestionDB
	err := r.db.Find(&questions).Error
	return questions, err
}
