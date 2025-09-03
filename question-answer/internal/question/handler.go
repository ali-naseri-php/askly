package question

import (
	"context"

	pb "github.com/ali-naseri-php/Askly/proto/question-answer"
)

type QuestionHandler struct {
	pb.UnimplementedQuestionServiceServer
	service *QuestionService
}

func NewQuestionHandler(svc *QuestionService) *QuestionHandler {
	return &QuestionHandler{service: svc}
}

func (h *QuestionHandler) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.CreateQuestionResponse, error) {
	q, err := h.service.CreateQuestion(req.Title, req.Body, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.CreateQuestionResponse{
		Question: mapToProtoQuestion(q),
	}, nil
}

func (h *QuestionHandler) GetQuestion(ctx context.Context, req *pb.GetQuestionRequest) (*pb.GetQuestionResponse, error) {
	q, err := h.service.GetQuestion(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetQuestionResponse{
		Question: mapToProtoQuestion(q),
	}, nil
}

func (h *QuestionHandler) ListQuestions(ctx context.Context, req *pb.ListQuestionsRequest) (*pb.ListQuestionsResponse, error) {
	list, err := h.service.ListQuestions()
	if err != nil {
		return nil, err
	}
	resp := &pb.ListQuestionsResponse{}
	for _, q := range list {
		resp.Questions = append(resp.Questions, mapToProtoQuestion(q))
	}
	return resp, nil
}

// helper برای تبدیل به proto
func mapToProtoQuestion(q *QuestionDB) *pb.Question {
	return &pb.Question{
		Id:        q.ID,
		Title:     q.Title,
		Body:      q.Body,
		Token:     q.UserID, // برای client فقط UserID را توکن می‌کنیم
		CreatedAt: q.CreatedAt,
	}
}
