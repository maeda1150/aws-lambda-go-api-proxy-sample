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

* build binary(in container)

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

  * get

  ```
  $ curl 'http://127.0.0.1:3000/test?hoge=fuga'
  ```

  * post

  ```
  $ curl -XPOST 'http://127.0.0.1:3000/test' -H "Content-Type: application/json" -d '{"hoge":"fuga"}'
  ```

## Deployment

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

## Test deployed api

* get apigateway id

```
$ aws apigateway get-rest-apis --output json --query 'items[?name==`lambdaapi`].id'
```

* post api

  * get

  ```
  $ curl 'https://{{id}}.execute-api.ap-northeast-1.amazonaws.com/Prod/test?hoge=fuga'
  ```

  * post

  ```
  $ curl -XPOST 'https://{{id}}.execute-api.ap-northeast-1.amazonaws.com/Prod/test' -H "Content-Type: application/json" -d '{"hoge":"fuga"}'
  ```

REF: https://github.com/awslabs/aws-lambda-go-api-proxy
