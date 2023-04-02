# syntax=docker/dockerfile:1

FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go get github.com/Varma1506/user-account-management/useraccountmanagement/dbconfig
RUN go get github.com/Varma1506/user-account-management/model
RUN go get github.com/Varma1506/user-account-management/services
RUN go get github.com/Varma1506/user-account-management/token
RUN go mod download
COPY ./**/*.go ./
RUN go build -o /useraccount-management
CMD [ "/useraccount-management" ]