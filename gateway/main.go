package main
import(
  "net/http"
  "flag"
  "context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "github.com/golang/glog"
  gw "../proto"
  "google.golang.org/grpc"
)

var(
  endpoint = flag.String("endpoint", "localhost:8080", "input message")
)

func run(address string) (error) {
  ctx := context.Background()
  ctx,cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterPlayerServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)

  if err != nil {
    return err
  }
  return http.ListenAndServe(address, mux)
}
func main() {
  flag.Parse()
  defer glog.Flush()
  if err := run(":3000"); err != nil {
    glog.Fatal(err)
  }
}

