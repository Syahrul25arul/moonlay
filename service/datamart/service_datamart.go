package service_datamart

import (
	"context"
	"moonlay/model/domain"
)

type ServiceDatamart interface {
	GetDatamar1(ctx context.Context, interval int) []domain.Datamart1
	GetDatamar2(ctx context.Context, interval int) []domain.Datamart2
	GetDatamar3(ctx context.Context, interval int) []domain.Datamart3
	GetTotalData(ctx context.Context) int
}
