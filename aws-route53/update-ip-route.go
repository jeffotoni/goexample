// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
    //"flag"
    "fmt"
    //"os"
    "encoding/json"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/route53"
    "net/http"
    "time"
)

var (
    ZONEID = "your zoneid route53"
    DOMAIN = "subdomain or your-domain.com"
    REGION = "us-east-1"
)

type IpApi struct {
    Country     string  `json:"country"`
    CountryCode string  `json:"countryCode"`
    Region      string  `json:"region"`
    RegionName  string  `json:"regionName"`
    City        string  `json:"city"`
    District    string  `json:"district"`
    Zip         string  `json:"zip"`
    Lat         float64 `json:"lat"`
    Lon         float64 `json:"lon"`
    Timezone    string  `json:"timezone"`
    Isp         string  `json:"isp"`
    Org         string  `json:"org"`
    As          string  `json:"as"`
    Mobile      bool    `json:"mobile"`
    Proxy       bool    `json:"proxy"`
    Query       string  `json:"query"`
}

func IpNow() string {

    var ipapi IpApi
    timeout := time.Duration(10 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }

    response, err := client.Get("http://ip-api.com/json")
    if err != nil {
        fmt.Printf("\nError client.Get %s", err)
        return ""
    } else {
        defer response.Body.Close()
        err := json.NewDecoder(response.Body).Decode(&ipapi)
        if err != nil {
            fmt.Printf("\nError json.NewDecoder: %s", err)
            return ""
        }

        return ipapi.Query
    }

    return ""
}

func main() {
    svc := route53.New(session.New(), aws.NewConfig().WithRegion(REGION))

    ipNow := IpNow()
    fmt.Println(ipNow)

    if len(ipNow) <= 0 {
        fmt.Println("Nao encontramos seu Ip externo para atualizar no route53 seu domain: " + DOMAIN)
        return
    }

    input := &route53.ChangeResourceRecordSetsInput{
        ChangeBatch: &route53.ChangeBatch{
            Changes: []*route53.Change{
                {
                    Action: aws.String("UPSERT"),
                    ResourceRecordSet: &route53.ResourceRecordSet{
                        Name: aws.String(DOMAIN),
                        ResourceRecords: []*route53.ResourceRecord{
                            {
                                Value: aws.String(ipNow),
                            },
                        },
                        TTL:  aws.Int64(60), // 1minuto
                        Type: aws.String("A"),
                    },
                },
            },
            Comment: aws.String("Atualizando IP da impressora em Rede!"),
        },
        HostedZoneId: aws.String(ZONEID),
    }

    result, err := svc.ChangeResourceRecordSets(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case route53.ErrCodeNoSuchHostedZone:
                fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
            case route53.ErrCodeNoSuchHealthCheck:
                fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
            case route53.ErrCodeInvalidChangeBatch:
                fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
            case route53.ErrCodeInvalidInput:
                fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
            case route53.ErrCodePriorRequestNotComplete:
                fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
            default:
                fmt.Println(aerr.Error())
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
        return
    }

    fmt.Println(result)
}
