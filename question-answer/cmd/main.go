package main

import (
	"log"
	"net"
	"question-answer/pkg/migrations"
	"google.golang.org/grpc"

	"question-answer/pkg/db"
	"question-answer/config"

	"question-answer/internal/question"
	"question-answer/internal/answer"
	"question-answer/internal/category"

	questionAnswerpb "github.com/ali-naseri-php/Askly/proto/question-answer"
)

func main() {
	// --- Load config ---
	cfg := config.LoadConfig()

	// --- Init DB ---
	gormDB := db.InitDB()

	// --- AutoMigrate در حالت Dev ---
	if cfg.DevMode {
		if err := gormDB.AutoMigrate(
			&question.QuestionDB{},
			&answer.AnswerDB{},
			&category.CategoryDB{},
		); err != nil {
			log.Fatalf("AutoMigrate error: %v", err)
		}
		log.Println("✅ AutoMigrate finished")
	}

	// --- Init ماژول‌ها ---
	questionRepo := question.NewQuestionRepository(gormDB)
	questionSvc := question.NewQuestionService(questionRepo)
	questionHandler := question.NewQuestionHandler(questionSvc)

	answerRepo := answer.NewAnswerRepository(gormDB)
	answerSvc := answer.NewAnswerService(answerRepo)
	answerHandler := answer.NewAnswerHandler(answerSvc)

	categoryRepo := category.NewCategoryRepository(gormDB)
	categorySvc := category.NewCategoryService(categoryRepo)
	categoryHandler := category.NewCategoryHandler(categorySvc)

	// --- gRPC server ---
	lis, err := net.Listen("tcp", ":"+cfg.ServicePort)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register services
	questionAnswerpb.RegisterQuestionServiceServer(grpcServer, questionHandler)
	questionAnswerpb.RegisterAnswerServiceServer(grpcServer, answerHandler)
	// categoryHandler آماده ولی proto ندارد


	migrations.RunMigrations(gormDB)

	log.Printf("🚀 Question/Answer Service running on :%s", cfg.ServicePort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
