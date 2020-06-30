FROM golang:1.13.4

WORKDIR /app

COPY ./csvReader/go.mod ./csvReader/go.sum ./
RUN go mod download

COPY ./csvReader ./

EXPOSE 4000
