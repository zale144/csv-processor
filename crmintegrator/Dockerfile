FROM golang:1.13.4

WORKDIR /app

COPY ./crmIntegrator/go.mod ./crmIntegrator/go.sum ./
RUN go mod download

COPY ./crmIntegrator ./

EXPOSE 4001
