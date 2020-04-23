package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

const defaultName = "world"

var (
	addr     = flag.String("addr", "localhost:50051", "Address of grpc server.")
	key      = flag.String("api-key", "", "API key.")
	token    = flag.String("token", "", "Authentication token.")
	keyfile  = flag.String("keyfile", "", "Path to a Google service account key file.")
	audience = flag.String("audience", "", "Audience.")
)
var keep = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func main() {
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile("../certs/server.crt", "")
	if err != nil {
		panic("could not load tls cert: %s" + err.Error())
	}

	opts := grpc.WaitForReady(false)

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr,
		//grpc.WithInsecure()
		grpc.WithTransportCredentials(creds),
		grpc.WithDefaultCallOptions(opts),
		grpc.WithKeepaliveParams(keep),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	if *keyfile != "" {
		log.Printf("Authenticating using Google service account key in %s", *keyfile)
		keyBytes, err := ioutil.ReadFile(*keyfile)
		if err != nil {
			log.Fatalf("Unable to read service account key file %s: %v", *keyfile, err)
		}

		tokenSource, err := google.JWTAccessTokenSourceFromJSON(keyBytes, *audience)
		if err != nil {
			log.Fatalf("Error building JWT access token source: %v", err)
		}
		jwt, err := tokenSource.Token()
		if err != nil {
			log.Fatalf("Unable to generate JWT token: %v", err)
		}
		*token = jwt.AccessToken
		// NOTE: the generated JWT token has a 1h TTL.
		// Make sure to refresh the token before it expires by calling TokenSource.Token() for each outgoing requests.
		// Calls to this particular implementation of TokenSource.Token() are cheap.
	}

	ctx := context.Background()
	if *key != "" {
		log.Printf("Using API key: %s", *key)
		ctx = metadata.AppendToOutgoingContext(ctx, "x-api-key", *key)
	}
	if *token != "" {
		log.Printf("Using authentication token: %s", *token)
		ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", fmt.Sprintf("Bearer %s", *token))
	}

	// Contact the server and print out its response.
	name := defaultName
	if len(flag.Args()) > 0 {
		name = flag.Arg(0)
	}
	r, err := c.CepSearch(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
