# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine AS build-release-stage


WORKDIR /bin
# install bash, Java11 and Nodejs for Swagger Codegen CLI and AsyncAPI CLI
RUN apk update && apk add bash openjdk11 curl nodejs npm && \
    wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.50/swagger-codegen-cli-3.0.50.jar -O swagger-codegen-cli.jar && \
    echo '#!/bin/bash' > /bin/swagger-codegen && echo 'java -jar /bin/swagger-codegen-cli.jar "$@"' >> /bin/swagger-codegen && chmod +x /bin/swagger-codegen && \
    npm install -g @asyncapi/cli


WORKDIR /

COPY --from=build-stage /api /api


EXPOSE 8080


ENTRYPOINT ["/api"]