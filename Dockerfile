FROM golang:1.20-alpine AS base

FROM base AS api
RUN go install github.com/codegangsta/gin@latest
ARG WORKDIR
WORKDIR $WORKDIR
COPY . ./

FROM base AS test
ARG GOTESTSUMVERSION=1.9.0
RUN go install gotest.tools/gotestsum@v${GOTESTSUMVERSION}
ARG WORKDIR
WORKDIR $WORKDIR
COPY . ./