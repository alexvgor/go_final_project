FROM golang:1.23-alpine

WORKDIR /app/src

COPY go.mod go.sum ./

RUN go mod download

COPY internal/ ./internal
COPY cmd/ ./cmd

RUN CGO_ENABLED=0 go build -o ../go_final_project ./cmd/app

WORKDIR /app

RUN rm -r ./src

COPY web/ ./web

EXPOSE ${TODO_PORT}

CMD ["./go_final_project"]
