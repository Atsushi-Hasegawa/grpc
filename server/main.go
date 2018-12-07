package main
import(
  "context"
  "net"
  "log"
  pb "../proto"
  "google.golang.org/grpc"
)

type playerService struct {}
const(
  port = ":8080"
)
func (p *playerService) GetPlayer(ctx context.Context, req *pb.GetPlayerRequest) (*pb.PlayerResponse, error) {
  return &pb.PlayerResponse{
    Id: req.Id,
    Name: req.Name,
    Age: req.Age,
  }, nil
}

func (p *playerService) GetPlayerList(ctx context.Context, req *pb.PlayerCountryRequest) (*pb.PlayersResponse,error) {
  return &pb.PlayersResponse {
    Lang: req.Lang,
    Users: []*pb.PlayerResponse {
      {Id: 10, Name: "Tom", Age: 10},
      {Id: 20, Name: "Bob", Age: 20},
    },
  }, nil
}

func main() {
  listenPort, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatal(err)
  }
  s := grpc.NewServer()
  pb.RegisterPlayerServiceServer(s, new(playerService))
  if err = s.Serve(listenPort); err != nil {
    log.Fatal(err)
  }
}
