FROM registry.cn-shanghai.aliyuncs.com/pigeonligh/build:golang-1.16-alpine AS build

WORKDIR /build

ADD cmd ./cmd
ADD pkg ./pkg
ADD go.* ./
ADD Makefile ./

RUN make
