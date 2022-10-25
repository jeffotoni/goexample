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
    --function-name gping \
    --runtime go1.x \
    --zip-file fileb://function.zip \
    --handler index.handler \
    --role arn:aws:iam::123456789012:role/lambda-url-role
```

### Create URL endpoint
```bash
aws lambda create-function-url-config \
    --function-name gping \
    --auth-type NONE
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