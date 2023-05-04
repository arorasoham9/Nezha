FROM golang:latest AS builder

ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o /main .

FROM ubuntu:latest
# Needed to avoid `x509: a certificate signed by unknown authority`
COPY --from=builder /main ./

ENTRYPOINT ["./main"]
EXPOSE 8000


#dockerfile from API is below

# FROM golang:latest AS builder

# ADD . /app
# WORKDIR /app

# RUN go mod download
# RUN go build -o /main .

# FROM ubuntu:latest
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /main ./
# ENTRYPOINT ["./main"]
# EXPOSE 8000