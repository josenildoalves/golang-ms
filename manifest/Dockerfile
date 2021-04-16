FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get ./ && go build -o main . && rm -Rf main.go && mv main /
EXPOSE 9002
CMD [ "/main" ]

