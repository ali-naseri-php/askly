package routes

import (
	"context"
	"net/http"
	"time"

	questionAnswerpb "github.com/ali-naseri-php/Askly/proto/question-answer"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func RegisterQuestionAnswerRoutes(e *echo.Echo, conn *grpc.ClientConn) {
	questionClient := questionAnswerpb.NewQuestionServiceClient(conn)
	answerClient := questionAnswerpb.NewAnswerServiceClient(conn)

	// ایجاد سوال
	e.POST("/questions", func(c echo.Context) error {
		var req struct {
			Title string `json:"title"`
			Body  string `json:"body"`
			Token string `json:"token"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		resp, err := questionClient.CreateQuestion(ctx, &questionAnswerpb.CreateQuestionRequest{
			Title: req.Title,
			Body:  req.Body,
			Token: req.Token,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp.Question)
	})

	// گرفتن یک سوال خاص
	e.GET("/questions/:id", func(c echo.Context) error {
		id := c.Param("id")
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		resp, err := questionClient.GetQuestion(ctx, &questionAnswerpb.GetQuestionRequest{
			Id: id,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp.Question)
	})

	// لیست سوالات
	e.GET("/questions", func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		resp, err := questionClient.ListQuestions(ctx, &questionAnswerpb.ListQuestionsRequest{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp.Questions)
	})

	// ایجاد جواب برای سوال
	e.POST("/answers", func(c echo.Context) error {
		var req struct {
			QuestionID string `json:"question_id"`
			Body       string `json:"body"`
			Token      string `json:"token"`
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		resp, err := answerClient.CreateAnswer(ctx, &questionAnswerpb.CreateAnswerRequest{
			QuestionId: req.QuestionID,
			Body:       req.Body,
			Token:      req.Token,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp.Answer)
	})
}
