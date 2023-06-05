FROM golang:1.20-alpine AS base
RUN apk update && apk add make curl

FROM base AS dev
ARG WORKDIR
ARG AIRVERSION
RUN go install github.com/cosmtrek/air@${AIRVERSION}
WORKDIR $WORKDIR
COPY . ./

FROM base AS prod
ARG WORKDIR
ARG AIRVERSION
WORKDIR $WORKDIR
COPY . ./

FROM base AS test
ARG GOTESTSUMVERSION=v1.9.0
RUN go install gotest.tools/gotestsum@${GOTESTSUMVERSION}
ARG WORKDIR
WORKDIR $WORKDIR
COPY . ./