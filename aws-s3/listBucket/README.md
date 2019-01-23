### List Bucket S3

This program lists bucket s3 and generates a comma-separated .csv file.
id line, name, size.

To execute just do the following:

```go
$ go run lists3.go --region = us-east-1 --bucket = mybucket
```

or

```go
$ go build
$ ./lists3 --region = us-east-1 --bucket = mybucket
```

