package main

import (
    "flag"
    "fmt"
    "os"

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
    Transactional Type = "Transactional"
)

// Defaults.
var (
    DefaultMaxPrice = 0.01
    DefaultType     = Transactional
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
    sms := SMS{Service: service}
    return sms.Send(message, number)
}

func main() {

    flagFone := flag.String("fone", "", "example: 319809876543")
    flagMsg := flag.String("msg", "", "example: Send sms teste!")
    flag.Parse()

    if *flagFone == "" {
        flag.PrintDefaults()
        os.Exit(0)
    }

    if *flagMsg == "" {
        flag.PrintDefaults()
        os.Exit(0)
    }

    fone := *flagFone
    msg := *flagMsg
    err := Send(msg, "+55"+fone)
    if err == nil {
        fmt.Println("Enviando com sucesso!")
    } else {
        fmt.Println("Error ao enviar: ", err.Error())
    }
}
