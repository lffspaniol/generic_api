package temperature

import (
	context "context"

	"github.com/golang/protobuf/ptypes"
	"github.com/lffspaniol/generic_api/gen/pb-go/gen"
	"github.com/lffspaniol/generic_api/internal/storage"
	"google.golang.org/grpc"
)

type Server struct{}

func Register(s *grpc.Server) {
	gen.RegisterTemperatureServiceServer(s, &Server{})
}

func (*Server) NewTemperature(ctx context.Context, req *gen.NewTemperatureRequest) (*gen.NewTemperatureResponse, error) {
	storage.Add(ptypes.TimestampNow(), req.Temp)
	resutl := &gen.NewTemperatureResponse{Message: "sucesso"}
	return resutl, nil
}

func (*Server) GetLast5Temp(ctx context.Context, _ *gen.Nil) (*gen.GetLast5TempResponse, error) {
	temps := storage.Get()

	result := &gen.GetLast5TempResponse{
		Temps: temps,
	}
	return result, nil
}
