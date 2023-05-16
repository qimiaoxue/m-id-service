FROM golang:1.18 as base

WORKDIR /opt/app/api
COPY . .

RUN go get -d -v ./...
RUN go build -o /usr/bin/m-id-server ./main.go

EXPOSE 8080 8888

CMD ["m-id-server"]