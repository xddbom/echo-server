FROM golang:1.22-alpine
WORKDIR /app
ADD . .
RUN go build -o server .
EXPOSE 8080
CMD ["./server"]