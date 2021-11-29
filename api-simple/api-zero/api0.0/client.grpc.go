package main

// import (
//   "log"
//   "golang.org/x/net/context"
//   "google.golang.org/grpc"
// )

// func main() {

//   var conn *grpc.ClientConn
//   conn, err := grpc.Dial("0.0.0.0:9001", grpc.WithInsecure())
//   if err != nil {
//     log.Fatalf("did not connect: %s", err)
//   }
//   defer conn.Close()

//   c := NewMainServiceClient(conn)

//   response, err := c.SayHello(context.Background(), &Message{Body: "Hello Client GRPC..."})
//   if err != nil {
//     log.Fatalf("Error when calling SayHello: %s", err)
//   }
//   log.Printf("Response from server: %s", response.Body)

//   response, err = c.BroadcastMessage(context.Background(), &Message{Body: "Message to Broadcast grpc tests!"})
//   if err != nil {
//     log.Fatalf("Error when calling Broadcast Message: %s", err)
//   }
//   log.Printf("Response from server: %s", response.Body)

// }
