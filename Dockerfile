FROM golang:1.19
EXPOSE 10501:10501
WORKDIR /usr/src/app
RUN git clone https://github.com/wintermute767/VulnScanner.git .
RUN go mod download && go mod verify
RUN apt-get update && apt-get upgrade -y && apt-get install curl -y && apt-get install nmap -y
RUN go build -v -o /usr/bin/app ./main.go
ENV LOGLEVEL trace
ENV PORT 10501
CMD ["app"]
