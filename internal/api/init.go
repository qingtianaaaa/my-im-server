package api

import (
	"context"
	"errors"
	"fmt"
	"my-im-server/pkg/common/config"
	kdis "my-im-server/pkg/common/discovery"
	"my-im-server/sdk/tools/errs"
	"my-im-server/sdk/tools/system/program"
	"my-im-server/sdk/tools/utils/network"
	"my-im-server/sdk/tools/utils/runtime"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	*config.AllConfig
	RuntimeEnv string
	ConfigPath string
}

func Start(ctx context.Context, config *Config) error {
	apiPort := config.API.Api.Port

	config.RuntimeEnv = runtime.PrintRuntimeEnv()

	client, err := kdis.NewDiscoveryRegister(&config.Discovery, config.RuntimeEnv)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}
	client.AddOptions(grpc.WithInsecure())

	var (
		netDone = make(chan struct{}, 1)
		netErr  error
	)
	// registerIp, err := network.NewRpcRegisterIp("")
	// if err != nil {
	// 	return err
	// }

	router, err := newGinRouter(ctx, client, config)
	if err != nil {
		return err
	}

	address := net.JoinHostPort(network.GetListenIp(config.API.Api.ListenIP), strconv.Itoa(apiPort))

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			netErr = errs.WrapErr(err, fmt.Sprintf("api start err: %s", server.Addr))
			netDone <- struct{}{}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	shutdown := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			return errs.WrapErr(err, "shutdown err")
		}
		return nil
	}

	select {
	case <-sigs:
		program.SIGTERMExit()
		if err := shutdown(); err != nil {
			return err
		}
	case <-netDone:
		close(netDone)
		return netErr
	}
	return nil
}
