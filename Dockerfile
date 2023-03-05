FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go install net/http
CMD ["go", "run", "main.go"]
