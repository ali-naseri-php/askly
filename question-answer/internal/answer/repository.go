package answer

import "gorm.io/gorm"

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

func (r *AnswerRepository) Create(a *AnswerDB) error {
	return r.db.Create(a).Error
}

func (r *AnswerRepository) ListByQuestionID(questionID string) ([]*AnswerDB, error) {
	var answers []*AnswerDB
	err := r.db.Where("question_id = ?", questionID).Find(&answers).Error
	return answers, err
}
