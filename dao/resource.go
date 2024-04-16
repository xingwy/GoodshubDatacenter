package dao

import (
	"context"
)

type ProjectResourceDB interface {
}

type ProjectResourceDBInstance struct {
	D *Dao
}

func (i *ProjectResourceDBInstance) CreateProjectResource(ctx context.Context, resource any) error {
	return nil
}
