package answer

import (
	"question-answer/pkg/middleware"
)

type AnswerService struct {
	repo *AnswerRepository
}

func NewAnswerService(repo *AnswerRepository) *AnswerService {
	return &AnswerService{repo: repo}
}

// CreateAnswer با استفاده از middleware برای استخراج userID
func (s *AnswerService) CreateAnswer(questionID, body, token string) (*AnswerDB, error) {
	userID, err := middleware.ExtractUserIDFromToken(token)
	if err != nil {
		return nil, err
	}

	a := NewAnswer(questionID, body, userID)
	if err := s.repo.Create(a); err != nil {
		return nil, err
	}
	return a, nil
}
