package apiserver

import (
	"log"

	"github.com/daz2yy/go-base/internal/apiserver/config"
	"github.com/daz2yy/go-base/internal/apiserver/store/mysql"
	"github.com/daz2yy/go-base/pkg/shutdown"
	"github.com/daz2yy/go-base/pkg/shutdown/shutdownmanagers/posixsignal"

	genericoptions "github.com/daz2yy/go-base/internal/pkg/options"
	genericapiserver "github.com/daz2yy/go-base/internal/pkg/server"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	genericAPIServer *genericapiserver.GenericAPIServer
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	// TODO: GRPC Server
	// extraConfig, err := buildExtraConfig(cfg)
	// if err != nil {
	// 	return nil, err
	// }

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	// extraServer, err := extraConfig.complete().New()

	server := &apiServer{
		gs:               gs,
		genericAPIServer: genericServer,
	}

	return server, nil
}

type preparedAPIServer struct {
	*apiServer
}

func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	// s.initRedisStore()

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		if mysqlStore != nil {
			return mysqlStore.Close()
		}

		// s.gRPCAPIServer.Close()
		s.genericAPIServer.Close()

		return nil
	}))

	return preparedAPIServer{s}
}

func (s *apiServer) Run() error {
	// go s.gRPCAPIServer.Run()

	// start shutdown managers.
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericAPIServer.Run()
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

// ExtraConfig defines extra configuration for the iam-apiserver.
type ExtraConfig struct {
	Addr       string
	MaxMsgSize int
	ServerCert genericoptions.GeneratableKeyCert
	// mysqlOptions *genericoptions.MySQLOptions
}

type completedExtraConfig struct {
	*ExtraConfig
}

func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}

//nolint: unparam
func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	// return &ExtraConfig{
	// 	Addr:         fmt.Sprintf("%s:%d", cfg.GRPCOptions.BindAddress, cfg.GRPCOptions.BindPort),
	// 	MaxMsgSize:   cfg.GRPCOptions.MaxMsgSize,
	// 	ServerCert:   cfg.SecureServing.ServerCert,
	// 	mysqlOptions: cfg.MySQLOptions,
	// 	// etcdOptions:      cfg.EtcdOptions,
	// }, nil
	return &ExtraConfig{}, nil
}
