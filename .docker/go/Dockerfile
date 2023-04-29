FROM golang:1.20.3-alpine AS build

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /tmp/app

COPY . .

RUN GOOS=linux go build -o ./out/api .

#########################################################

FROM alpine:latest

RUN apk add ca-certificates

COPY --from=build /tmp/app/out/api /app/api

WORKDIR "/app"

ARG PORT
ENV PORT "$PORT"
EXPOSE ${PORT}

CMD ["./api"]
