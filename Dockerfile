FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build ./main.go
EXPOSE 9002
CMD ["./main"]