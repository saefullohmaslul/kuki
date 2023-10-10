package todos

import "github.com/saefullohmaslul/kuki/internal/grpc"

type GrpcHandler interface {
	grpc.TodosHandlerServer
}
