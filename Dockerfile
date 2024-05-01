FROM golang:1.22.2

COPY . /app
WORKDIR /app


RUN go build -o bin .

EXPOSE 8080

CMD [ "./bin" ]
