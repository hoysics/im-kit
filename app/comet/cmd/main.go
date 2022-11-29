package main

import (
	"github.com/hoysics/im-kit/app/comet/internal/data"
	"github.com/hoysics/im-kit/app/comet/internal/server"
	"github.com/hoysics/im-kit/app/comet/internal/service"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen(`tcp`, `0.0.0.0:7000`)
	if err != nil {
		log.Fatalf(`net listen err: %v`, err)
	}
	lc, lcClose, err := data.NewLogicClient()
	if err != nil {
		log.Fatalf(`connect logic err: %v`, err)
	}
	defer lcClose()
	cr := data.NewConnRepo()
	lr := data.NewLogicRepo(lc)
	srv := server.NewGRPCServer(service.NewBrokerService(cr), service.NewCometService(lr, cr))
	if err := srv.Serve(lis); err != nil {
		log.Fatalf(`grpc serve err: %v`, err)
	}
}
