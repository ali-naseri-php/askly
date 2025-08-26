# مرحله build
FROM golang:1.24.6-alpine AS builder

WORKDIR /app

# کپی go.mod و go.sum برای کش بهتر
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../ .
COPY . .
COPY ../proto ../proto   

# کپی کل سورس پروژه


# بیلد برنامه (فایل main داخل cmd)
RUN go build -o auth-service ./cmd

# مرحله نهایی: فقط باینری سبک
FROM alpine:latest

WORKDIR /app

# کتابخونه‌های ضروری برای اجرای Go
RUN apk add --no-cache ca-certificates tzdata

# کپی باینری از مرحله build
COPY --from=builder /app/auth-service .

EXPOSE 8001

CMD ["./auth-service"]
