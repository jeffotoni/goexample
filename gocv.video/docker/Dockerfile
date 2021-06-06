FROM gbaeke/gocv-4.0.0-build as build       
RUN go get -u -d gocv.io/x/gocv       
RUN go get -u -d github.com/disintegration/imaging       
RUN go get -u -d github.com/gbaeke/emotion       
RUN cd $GOPATH/src/github.com/gbaeke/emotion && go build -o $GOPATH/bin/emo ./main.go       
                
FROM gbaeke/gocv-4.0.0-run       
COPY --from=build /go/bin/emo /emo       
ADD haarcascade_frontalface_default.xml /       

ENTRYPOINT ["/emo"]
