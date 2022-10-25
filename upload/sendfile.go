import (
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "time"
)

func main() {

    var client = &http.Client{
        Timeout: time.Second * 10,
    }

    r, err := UploadMultipartFile(client, "http://localhost:5001/upload", "file", "./jeffotoni.png")
}

func UploadMultipartFile(client *http.Client, uri, key, path string) (*http.Response, error) {
    body, writer := io.Pipe()

    req, err := http.NewRequest(http.MethodPost, uri, body)
    if err != nil {
        return nil, err
    }

    mwriter := multipart.NewWriter(writer)
    req.Header.Add("Content-Type", mwriter.FormDataContentType())

    errchan := make(chan error)

    go func() {
        defer close(errchan)
        defer writer.Close()
        defer mwriter.Close()

        w, err := mwriter.CreateFormFile(key, path)
        if err != nil {
            errchan <- err
            return
        }

        in, err := os.Open(path)
        if err != nil {
            errchan <- err
            return
        }
        defer in.Close()

        if written, err := io.Copy(w, in); err != nil {
            errchan <- fmt.Errorf("error copying %s (%d bytes written): %v", path, written, err)
            return
        }

        if err := mwriter.Close(); err != nil {
            errchan <- err
            return
        }
    }()

    resp, err := client.Do(req)
    merr := <-errchan

    if err != nil || merr != nil {
        return resp, fmt.Errorf("http error: %v, multipart error: %v", err, merr)
    }

    return resp, nil
}