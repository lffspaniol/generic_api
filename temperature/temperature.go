package temperature

import (
	context "context"
	"log"

	"github.com/lffspaniol/generic_api/gen"
	"google.golang.org/grpc"
)

type Server struct{}

func Register(s *grpc.Server) {
	gen.RegisterTemperatureServiceServer(s, &Server{})
}

func (*Server) NewTemperature(ctx context.Context, req *gen.NewTemperatureRequest) (*gen.NewTemperatureResponse, error) {
	resutl := &gen.NewTemperatureResponse{Message: []string{"batata"}}
	log.Println("Request")
	return resutl, nil
}

func (*Server) GetLast5Temp(ctx context.Context, _ *gen.Nil) (*gen.GetLast5TempResponse, error) {
	result := &gen.GetLast5TempResponse{
		Temps: "batata",
	}
	return result, nil
}
