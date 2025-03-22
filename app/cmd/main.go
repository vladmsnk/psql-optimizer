package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"psql-optimizer/internal/app/optimizer"
	"psql-optimizer/internal/usecases/indexes"
	"syscall"

	"psql-optimizer/adapter/postgres"
	grpclib "psql-optimizer/utils/grpc_server"
	psqllib "psql-optimizer/utils/postgres"
)

var (
	defaultDSN        = "user=user password=password host=localhost port=5432 dbname=postgres sslmode=disable"
	defaultServerHost = ""
	defaultServerPort = 8080
)

func main() {
	ctx := context.Background()

	postgresConfig := &psqllib.Config{
		DSN: defaultDSN,
	}

	grpcConfig := &grpclib.Config{
		Host: defaultServerHost,
		Port: defaultServerPort,
	}

	psqlConn, err := psqllib.NewWithContext(ctx, postgresConfig)
	if err != nil {
		log.Fatalf("postgres.NewWithContext: %s", err.Error())
	}

	grpcServer, err := grpclib.New(grpcConfig)
	if err != nil {
		log.Fatalf("grpclib.New: %s", err.Error())
	}

	postgresInfoProvider := postgres.New(psqlConn.Pool)

	indexInfoGetter := indexes.New(postgresInfoProvider)

	transport := optimizer.New(indexInfoGetter)

	serv, err := runGRPCServer(transport, grpcServer)
	if err != nil {
		log.Fatalf("runGRPCServer: %s", err.Error())
	}
	defer serv.GracefulStop()

	Lock(make(chan os.Signal))
}

func Lock(ch chan os.Signal) {
	defer func() {
		ch <- os.Interrupt
	}()
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-ch
}
