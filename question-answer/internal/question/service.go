package question

import (
	"errors"
	"strings"
)

// QuestionService منطق اصلی سوال‌ها
type QuestionService struct {
	repo *QuestionRepository
}

func NewQuestionService(repo *QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

// CreateQuestion با استخراج userID از توکن
func (s *QuestionService) CreateQuestion(title, body, token string) (*QuestionDB, error) {
	userID, err := extractUserIDFromToken(token)
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

// --- helper برای استخراج userID از توکن ---
func extractUserIDFromToken(token string) (string, error) {
	// placeholder: در عمل باید JWT decode بشه
	if token == "" {
		return "", errors.New("invalid token")
	}
	// فرض می‌کنیم userID بعد از ":" در token قرار داره
	// مثلا token = "userID:12345"
	parts := strings.Split(token, ":")
	if len(parts) != 2 {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
