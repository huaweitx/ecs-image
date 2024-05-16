ARG GLOBAL_BUILD_BASEIMAGE=image.cestc.cn/xx/xx:xx

FROM --platform=$BUILDPLATFORM image.cestc.cn/xx/golang:xx
RUN mkdir xx
COPY . xx/

FROM image.cestc.cn/xxb/xx:20240223

FROM ${GLOBAL_BUILD_BASEIMAGE}
CMD ["./xx"]
