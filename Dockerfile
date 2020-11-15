#################
# Build Stage
#################
FROM golang:1.15-alpine as builder

RUN apk add make git

RUN mkdir -p /go/src/github.com/Bak3y/darkwind_degreaser
WORKDIR /go/src/github.com/Bak3y/darkwind_degreaser

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o darkwind_degreaser cmd/main.go


###################
# Package Stage
###################
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# copy our static linked library from the previous stage
COPY --from=builder /go/src/github.com/Bak3y/darkwind_degreaser/darkwind_degreaser /usr/local/bin/darkwind_degreaser

ENTRYPOINT [ "darkwind_degreaser" ]