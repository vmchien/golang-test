FROM golang:1.19-alpine

WORKDIR /app/zoro

COPY . .

RUN go mod tidy

RUN cd /app/zoro/cmd/ && go build -o /api

EXPOSE 5678

CMD ["/api"]