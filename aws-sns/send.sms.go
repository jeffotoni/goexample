package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sns"
    "github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// Type of SMS delivery mode.
type Type string

const (
    // Promotional are non-critical messages, such as marketing messages.
    // Amazon SNS optimizes the message delivery to incur the lowest cost.
    Promotional Type = "Promotional"

    // Transactional messages are critical messages that support
    // customer transactions, such as one-time passcodes for multi-factor authentication.
    // Amazon SNS optimizes the message delivery to achieve the highest reliability.
    Transactional = "Transactional"
)

// Defaults.
var (
    DefaultMaxPrice = 0.01
    DefaultType     = Promotional
)

// SMS configures an SNS SMS client.
type SMS struct {
    Service  snsiface.SNSAPI // Service implementation
    SenderID string          // SenderID (optional)
    Type     Type            // Type of SMS delivery mode
    MaxPrice float64         // MaxPrice (defaults to $0.01)
}

// Send `message` to `number`.
func (s *SMS) Send(message, number string) error {

    attrs := map[string]*sns.MessageAttributeValue{}

    if s.SenderID != "" {
        attrs["AWS.SNS.SMS.SenderID"] = &sns.MessageAttributeValue{
            DataType:    aws.String("String"),
            StringValue: &s.SenderID,
        }
    }

    // maxPrice := s.MaxPrice
    // if maxPrice == 0 {
    //     maxPrice = DefaultMaxPrice
    // }

    // attrs["AWS.SNS.SMS.MaxPrice"] = &sns.MessageAttributeValue{
    //     DataType:    aws.String("String"),
    //     StringValue: aws.String(fmt.Sprintf("%0.5f", maxPrice)),
    // }

    kind := s.Type
    if kind == "" {
        kind = DefaultType
    }

    attrs["AWS.SNS.SMS.SMSType"] = &sns.MessageAttributeValue{
        DataType:    aws.String("String"),
        StringValue: aws.String(string(kind)),
    }

    params := &sns.PublishInput{
        Message:           &message,
        PhoneNumber:       &number,
        MessageAttributes: attrs,
    }

    _, err := s.Service.Publish(params)
    return err
}

// Send `message` to `number` using defaults.
func Send(message, number string) error {
    service := sns.New(session.New(aws.NewConfig().WithRegion("us-east-1")))
    //service := sns.New(session.New(aws.NewConfig()))
    sms := SMS{Service: service}
    return sms.Send(message, number)
}

func main() {

    err := Send("Send SMS jeff..", "+15531987087256")
    fmt.Println(err)
}
