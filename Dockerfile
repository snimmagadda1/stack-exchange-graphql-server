FROM golang:1.15-alpine3.12

EXPOSE 8080
ENV SERVER test
ENV UNAME testu
ENV PASS testpass
ENV SCHEMA stacke

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main cmd/server.go
CMD ["/app/main"]