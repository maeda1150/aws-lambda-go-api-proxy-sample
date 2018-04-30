## Require

* Docker
* [AWS SAM](https://github.com/awslabs/serverless-application-model)
* awscli

## Development

* provisioning

```
$ docker-compose build
```

* development on golang-vim-dev

```
$ docker run --rm -tiv `pwd`:/go/src/aws-lambda-go-api-proxy-sample aws-lambda-go-api-proxy-sample_lambdaapi

# fish shell
$ docker run --rm -tiv (pwd):/go/src/aws-lambda-go-api-proxy-sample aws-lambda-go-api-proxy-sample_lambdaapi
```

* build bin(in container)

```
$ GOOS=linux GOARCH=amd64 go build -o api .
```

* invoke lambda on local

```
$ echo '{"hoge": "fuga"}' | sam local invoke ApiPOST
```

or

```
$ sam local invoke ApiPOST -e event.json
```

* validate template.yml

```
$ sam validate
```

* start api

```
$ sam local start-api
```

* post localhost

get

```
$ curl 'http://127.0.0.1:3000/test?hoge=fuga'
```

post

```
$ curl -XPOST 'http://127.0.0.1:3000/test' -H "Content-Type: application/json" -d '{"hoge":"fuga"}'
```

* make s3 bucket

```
$ aws s3 mb s3://sum-lambda
```

* upload s3 bucket & create package.yml

```
$ sam package --template-file template.yaml --s3-bucket sum-lambda --output-template-file package.yaml
```

* deploy cloudformation & lambda

```
$ aws cloudformation deploy --template-file package.yaml --stack-name lambdaapi --capabilities CAPABILITY_IAM
```

* test api

```
$ curl -XPOST https://xxxxxxxx.ap-northeast-1.amazonaws.com/Prod/test
```

```
$ curl -XGET https://xxxxxxxxx.ap-northeast-1.amazonaws.com/Prod/test
```


REF: https://github.com/awslabs/aws-lambda-go-api-proxy
