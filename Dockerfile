FROM golang

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /tails-social-go

CMD ["/tails-social-go"]