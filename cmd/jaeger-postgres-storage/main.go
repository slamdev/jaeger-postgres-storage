package main

import (
	"flag"
	"github.com/hashicorp/go-hclog"
	"github.com/jaegertracing/jaeger/plugin/storage/grpc"
	"github.com/jaegertracing/jaeger/storage/dependencystore"
	"github.com/jaegertracing/jaeger/storage/spanstore"
	"github.com/slamdev/jaeger-postgres-storage/pkg"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

var configPath string

func main() {
	log := hclog.New(&hclog.LoggerOptions{
		Name:  "jaeger-postgres-storage",
		Level: hclog.Debug,
	})

	flag.StringVar(&configPath, "config", "", "A path to the plugin's configuration file")
	flag.Parse()

	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// TODO remove
	//configPath = "hack/config.yaml"
	//os.Setenv("POSTGRES_HOST", "localhost")

	if configPath != "" {
		v.SetConfigFile(configPath)
		err := v.ReadInConfig()
		if err != nil {
			log.Error("failed to parse configuration file", "err", err)
			os.Exit(1)
		}
	}

	c := pkg.Config{}
	if err := c.InitFromViper(v); err != nil {
		log.Error("failed to init config", "err", err)
		os.Exit(1)
	}

	// TODO waits until PG container starts. remove it
	time.Sleep(3 * time.Second)

	repo, err := pkg.NewRepository(c)
	if err != nil {
		log.Error("failed to create repository", "err", err)
		os.Exit(1)
	}

	log.Info("postgres connection established")

	if err := repo.Migrate(); err != nil {
		log.Error("failed to run database migrations", "err", err)
		os.Exit(1)
	}

	log.Info("migrations are applied")

	reader := pkg.NewReader(repo, log)
	writer := pkg.NewWriter(repo, log)
	depReader := pkg.NewDepReader(repo, log)

	p := plugin{
		reader:    reader,
		writer:    writer,
		depReader: depReader,
	}

	grpc.Serve(&p)

	if err := repo.Close(); err != nil {
		log.Error("failed to close repository", "err", err)
		os.Exit(1)
	}
}

type plugin struct {
	reader    spanstore.Reader
	writer    spanstore.Writer
	depReader dependencystore.Reader
}

func (p *plugin) SpanReader() spanstore.Reader {
	return p.reader
}

func (p *plugin) SpanWriter() spanstore.Writer {
	return p.writer
}

func (p *plugin) DependencyReader() dependencystore.Reader {
	return p.depReader
}
