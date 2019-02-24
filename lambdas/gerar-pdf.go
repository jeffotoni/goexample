// Go Api server
// @jeffotoni
// 2019-02-22

package main

import (
    "bytes"
    "context"
    "fmt"
    "io"
    "log"
    "os"
    "strings"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

var sess *session.Session

func init() {
    // Let wkhtmltopdf where to find our bin file (in the zip)
    os.Setenv("WKHTMLTOPDF_PATH", os.Getenv("LAMBDA_TASK_ROOT"))
    // Setup AWS S3 Session (build once use every function)
    sess = session.Must(session.NewSession(&aws.Config{
        Region: aws.String("eu-west-1"),
    }))
}

func main() {
    // Start() is when Lambda will connect and start running functions.
    // Everything before this is setup, and anything after won't be run.
    lambda.Start(LambdaHandler)
}

func LambdaHandler(ctx context.Context, s3Event events.S3Event) {
    for _, record := range s3Event.Records {
        fmt.Printf(
            "[%s - %s] Bucket = %s, Key = %s \n",
            record.EventSource,
            record.EventTime,
            record.S3.Bucket.Name,
            record.S3.Object.Key,
        )
        err := ProcessFile(record)
        if err != nil {
            log.Println(err)
            continue
        }
    }
}

// ProcessFile handles a single S3 event, this will do the conversion from HTML
// to PDF for a single file only and write it back to it's origin.
func ProcessFile(record events.S3EventRecord) error {
    s3Item := record.S3

    // Get the HTML file from S3
    obj, err := s3.New(sess).GetObject(&s3.GetObjectInput{
        Bucket: &s3Item.Bucket.Name,
        Key:    &s3Item.Object.Key,
    })
    defer obj.Body.Close()

    pdfBytes, err := GeneratePDF(obj.Body)
    if err != nil {
        return err
    }

    // Replace .html filename with .pdf
    newKey := strings.Replace(s3Item.Object.Key, ".html", ".pdf", -1)
    fmt.Println("Rename to: " + newKey)

    // Put the PDF back onto S3
    fmt.Println("Save File to: " + s3Item.Bucket.Name)
    _, err = s3.New(sess).PutObject(&s3.PutObjectInput{
        Bucket: &s3Item.Bucket.Name,
        Key:    &newKey,
        Body:   bytes.NewReader(pdfBytes),
    })
    return err
}

// GeneratePDF converts the file body from S3 into a PDF as a byte array to
// write back to S3.
func GeneratePDF(s3Obj io.Reader) ([]byte, error) {

    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        return nil, err
    }

    // Pass S3 Object body (as reader) directly into wkhtmltopdf
    pdfg.AddPage(wkhtmltopdf.NewPageReader(s3Obj))

    // Create PDF document in internal buffer
    if err := pdfg.Create(); err != nil {
        return nil, err
    }

    // Return PDF as bytes array
    return pdfg.Bytes(), nil
}
