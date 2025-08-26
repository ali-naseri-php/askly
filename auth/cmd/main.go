package main

import (
	"auth/internal/db"
	"auth/internal/delivery"
	"auth/internal/repository"
	"auth/internal/service"
	authpb "github.com/ali-naseri-php/Askly/proto/auth"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// --- تنظیمات دیتابیس از ENV ---
	dsn := os.Getenv("AUTH_DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=authdb port=5432 sslmode=disable"
	}

	// --- اتصال به دیتابیس ---
	gormDB := db.InitDB(dsn)

	// --- AutoMigrate اختیاری فقط در DEV ---
	if os.Getenv("DEV") == "true" {
		err := gormDB.AutoMigrate(&db.AuthDB{})
		if err != nil {
			log.Fatalf("AutoMigrate error: %v", err)
		}
		log.Println("✅ AutoMigrate finished")
	}

	// --- init لایه‌ها ---
	userRepo := repository.NewUserRepository(gormDB)
	authSvc := service.NewAuthService(userRepo)
	authHandler := delivery.NewAuthHandler(authSvc)

	// --- gRPC server ---
	port := os.Getenv("AUTH_SERVICE_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, authHandler)

	log.Printf("🚀 Auth Service running on :%s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
