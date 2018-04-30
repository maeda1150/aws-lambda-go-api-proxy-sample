FROM masaki1111/golang-vim-dev

RUN go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/golang/lint/golint

RUN go get -u github.com/aws/aws-lambda-go/events && \
    go get -u github.com/aws/aws-lambda-go/lambda && \
    go get -u github.com/awslabs/aws-lambda-go-api-proxy/...

WORKDIR /go/src/aws-lambda-go-api-proxy-sample/
ADD . /go/src/aws-lambda-go-api-proxy-sample/
