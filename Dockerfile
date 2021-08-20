FROM registry.cn-shanghai.aliyuncs.com/pigeonligh/build:golang-1.16-alpine AS build

WORKDIR /build

ADD cmd ./cmd

ADD pkg ./pkg

ADD go.* ./

ADD Makefile ./

RUN make

FROM registry.cn-shanghai.aliyuncs.com/pigeonligh/runtime:alpine-3.14

WORKDIR /app

ADD ffxiv-datamining-cn ./

COPY --from=build /build/_output/bin/server ./

ENTRYPOINT ["/app/server"]
