package table

import (
	mtype "github.com/slamdev/jaeger-postgres-storage/pkg/dbmodel/type"
	"time"
)

type DependenciesV2 struct {
	TsBucket     time.Time          `db:"ts_bucket"`
	Ts           time.Time          `db:"ts"`
	Dependencies []mtype.Dependency `db:"dependencies"`
}

func (d *DependenciesV2) TableName() string {
	return "dependencies_v2"
}
