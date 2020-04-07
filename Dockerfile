FROM golang:latest

WORKDIR $GOPATH/src/yisany.top/milu/dslConvert
COPY . $GOPATH/src/yisany.top/milu/dslConvert
RUN GOPROXY="https://goproxy.cn" go build .

EXPOSE 3232
ENTRYPOINT ["./dslConvert"]