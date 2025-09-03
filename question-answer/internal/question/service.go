package question

import (
	"question-answer/pkg/middleware"
)

type QuestionService struct {
	repo *QuestionRepository
}

func NewQuestionService(repo *QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

// CreateQuestion با استفاده از middleware برای استخراج userID
func (s *QuestionService) CreateQuestion(title, body, token string) (*QuestionDB, error) {
	userID, err := middleware.ExtractUserIDFromToken(token)
	if err != nil {
		return nil, err
	}

	q := NewQuestion(title, body, userID)
	if err := s.repo.Create(q); err != nil {
		return nil, err
	}
	return q, nil
}

func (s *QuestionService) GetQuestion(id string) (*QuestionDB, error) {
	return s.repo.GetByID(id)
}

func (s *QuestionService) ListQuestions() ([]*QuestionDB, error) {
	return s.repo.List()
}
