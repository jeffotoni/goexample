### List Bucket S3

A simple program to list bucket and its size and store everything in a csv file.

## Config Aws

Environment Variables

When a Session is created several environment variables can be set to adjust how the SDK functions, and what configuration data it loads when creating Sessions. All environment values are optional, but some values like credentials require multiple of the values to set or the partial values will be ignored. All environment variable values are strings unless otherwise noted.

Environment configuration values. If set both Access Key ID and Secret Access Key must be provided. Session Token and optionally also be provided, but is not required. 

### Access Key ID
AWS_ACCESS_KEY_ID=AKID
AWS_ACCESS_KEY=AKID # only read if AWS_ACCESS_KEY_ID is not set.

### Secret Access Key
AWS_SECRET_ACCESS_KEY=SECRET
AWS_SECRET_KEY=SECRET=SECRET # only read if AWS_SECRET_ACCESS_KEY is not set.

### Session Token
AWS_SESSION_TOKEN=TOKEN


```bash
$ export AWS_SECRET_ACCESS_KEY=xxxxx
$ export AWS_ACCESS_KEY_ID=xxxxx
```

## AWS Creating Sessions

When creating Sessions optional aws.Config values can be passed in that will override the default, or loaded config values the Session is being created with. This allows you to provide additional, or case based, configuration as needed. 

By default NewSession will only load credentials from the shared credentials file (~/.aws/credentials). If the AWS_SDK_LOAD_CONFIG environment variable is set to a truthy value the Session will be created from the configuration values from the shared config (~/.aws/config) and shared credentials (~/.aws/credentials) files. See the section Sessions from Shared Config for more information. 


### Install and Execute lists3

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

