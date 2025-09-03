package answer

import (
	"context"

	pb "github.com/ali-naseri-php/Askly/proto/question-answer"
)

type AnswerHandler struct {
	pb.UnimplementedAnswerServiceServer
	service *AnswerService
}

func NewAnswerHandler(svc *AnswerService) *AnswerHandler {
	return &AnswerHandler{service: svc}
}

func (h *AnswerHandler) CreateAnswer(ctx context.Context, req *pb.CreateAnswerRequest) (*pb.CreateAnswerResponse, error) {
	a, err := h.service.CreateAnswer(req.QuestionId, req.Body, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAnswerResponse{
		Answer: mapToProtoAnswer(a),
	}, nil
}

// helper برای تبدیل به proto
func mapToProtoAnswer(a *AnswerDB) *pb.Answer {
	return &pb.Answer{
		Id:         a.ID,
		QuestionId: a.QuestionID,
		Body:       a.Body,
		Token:      a.UserID, // برای client فقط UserID را توکن می‌کنیم
		CreatedAt:  a.CreatedAt,
	}
}
