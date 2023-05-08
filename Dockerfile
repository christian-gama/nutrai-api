FROM golang:1.20-alpine AS base
RUN apk update && apk add make

FROM base AS api
ARG WORKDIR
ARG AIRVERSION=1.43.0
RUN go install github.com/cosmtrek/air@v${AIRVERSION}
WORKDIR $WORKDIR
COPY . ./

FROM base AS test
ARG GOTESTSUMVERSION=1.9.0
RUN go install gotest.tools/gotestsum@v${GOTESTSUMVERSION}
ARG WORKDIR
WORKDIR $WORKDIR
COPY . ./