package redispresence

import (
	"E-01/pkg/richerror"
	"context"
	"time"
)

func (d DB) Upsert(ctx context.Context, key string, timestamp int64, expTime time.Duration) error {
	const op = "redispresence.Upsert"
	
	_, err := d.adapter.Client().Set(ctx ,key,timestamp, expTime).Result()
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}