FROM golang:latest
WORKDIR /home/user/Desktop/sacrif

COPY . /home/user/Desktop/sacrif
RUN go get -t github.com/gin-gonic/gin
RUN go get -t gopkg.in/mgo.v2

COPY . .
EXPOSE 8080



CMD ["go run", "main.go"]
