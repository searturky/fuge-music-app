# FROM golang:1.21.7-alpine
# ENV CGO_ENABLED=0 \
#     GOOS=linux \
#     GOPROXY=https://proxy.golang.com.cn/,direct \
#     GO111MODULE=on
# WORKDIR /app
# COPY go.mod go.sum ./
# COPY app ./
# RUN go mod download & go build -o fuge
# # RUN go build -o fuge
# EXPOSE 8080
# CMD ["./fuge"]

FROM alpine:3.19.1
WORKDIR /app
COPY ./dist/fuge ./
COPY ./config/settings.yaml ./
EXPOSE 8080
CMD ["./fuge"]
