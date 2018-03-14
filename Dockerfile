FROM itchain/http-func

WORKDIR /go/src/hello-func

COPY ./main.go  /go/src/hello-func

RUN go build -o hellofunc .

FROM itchain/http-func

COPY --from=0 /go/src/hello-func/hellofunc /usr/local/bin

CMD [ "httpfunc" ]