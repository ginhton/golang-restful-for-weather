FROM golang:1.18.2
WORKDIR /opt
COPY . /opt
RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN go get -u github.com/gorilla/mux
RUN go mod download
RUN go build -o grfw ./grfw.go

EXPOSE 8081
CMD ["/opt/grfw"]

