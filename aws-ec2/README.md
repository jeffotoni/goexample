### Start EC2 and Stop EC2

A simple script to start and stop in EC2 instances

#### Config Aws

Environment Variables
When a Session is created several environment variables can be set to adjust how the SDK functions, and what configuration data it loads when creating Sessions.
Environment configuration values. If set both Access Key ID and Secret Access Key must be provided.

#### KEY AND SECRET KEY

```bash
AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxx
AWS_ACCESS_KEY_ID=xxxxxx
```

Just export them to your environment and lists3 will be able to read your buckets and generate the csv file successfully.

```bash
$ export AWS_SECRET_ACCESS_KEY=xxxxx
$ export AWS_ACCESS_KEY_ID=xxxxx
```

#### AWS ./aws/credentials

Another way to configure AWS is by creating the directory and ./aws/credentials file in your $ HOME and configuring it.

Check the contents of the file below:

```bash
[default]
aws_access_key_id = xxxxxx
aws_secret_access_key = xxxxxxxx
```

#### Install and Execute start.go

To execute just do the following:

```go
$ go run start.go --region = us-east-1 --instance=i-05bef719c14d68d10
$ go run stop.go --region = us-east-1 --instance=i-05bef719c14d68d10
```

or

```go
$ go build
$ ./start --region=us-east-1 --instance=i-05bef719c14d68d10
$ ./stop --region=us-east-1 --instance=i-05bef719c14d68d10
```

