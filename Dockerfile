

FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go build -o /build/cmd/main.go

EXPOSE 8000

CMD [ "/build" ]
