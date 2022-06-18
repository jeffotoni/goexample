#gping AWS LAMBDA

Lambda Function

#### #equirements

    - awscli
    - Go
    - curl
    - Zip

### Zip

```bash
$ zip function.zip gping
```


### Build GO

```bash
$ CGO_ENABLED=0 GOOS=linux go build -o gping gping.go
```

### Create Lambda Function
```bash
$ aws lambda create-function \
  --code "ImageUri=${AWS_REPOSITORY_URI}:0.0.1" \
  --function-name gping \
  --package-type Image \
  --region us-us-east-1 \
  --role ${LAMBDA_ROLE_ARN}
```

### Send .zip to AWS LAMBDA

```bash
$ aws lambda update-function-code \
--function-name gping \
--zip-file fileb://function.zip
```

### Test Lambda function

```bash
$ aws lambda invoke \
  --function-name gping \
  --payload '{"input": "ping"}' \
  response.json
```

```bash
$ cat response.json
```
### Makefile

```bash
$ make build
```