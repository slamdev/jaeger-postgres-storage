package pkg

import (
	"github.com/hashicorp/go-hclog"
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/dependencystore"
	"time"
)

type depreader struct {
	repo Repository
	log  hclog.Logger
}

func NewDepReader(repo Repository, log hclog.Logger) dependencystore.Reader {
	return &depreader{
		repo: repo,
		log:  log.Named("dep-reader"),
	}
}

func (d *depreader) GetDependencies(endTs time.Time, lookback time.Duration) ([]model.DependencyLink, error) {
	log := d.log.Named("GetDependencies")
	log.Debug("inputs", "endTs", endTs, "lookback", lookback)
	return []model.DependencyLink{}, nil
}
