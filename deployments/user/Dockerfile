# 1 шаг - сборки
FROM golang:1.19-alpine AS build_stage
RUN apk add git
WORKDIR /go/src/
RUN git clone https://github.com/go-park-mail-ru/2023_1_BKS.git
WORKDIR /go/src/2023_1_BKS
RUN git checkout dev-user/crud
WORKDIR ./cmd/user
RUN go env -w GO111MODULE=auto
RUN go install .

# 2 шаг
FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/bin/user /app_binary/
RUN chmod +x ./user
EXPOSE 8080
ENTRYPOINT ./user
