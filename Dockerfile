FROM golang:1.18

WORKDIR /usr/src/app

EXPOSE 9999

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go get -d -v ./...
RUN go install
RUN go build -v -o /usr/local/bin/mainApp main.go
#
#CMD ["go test -v ./..."]
#CMD ["mainApp"]