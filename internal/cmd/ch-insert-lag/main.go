package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/go-faster/ch"
	"github.com/go-faster/ch/internal/cmd/app"
	"github.com/go-faster/ch/proto"
)

const ddl = `CREATE TABLE IF NOT EXISTS ch_insert_lag  (
    ts DateTime64(9)
) ENGINE MergeTree() ORDER BY (ts)`

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger) error {
		const precision = proto.PrecisionNano

		g, ctx := errgroup.WithContext(ctx)
		done := make(chan struct{})
		ready := make(chan struct{})
		g.Go(func() error {
			conn, err := ch.Dial(ctx, ch.Options{})
			if err != nil {
				return errors.Wrap(err, "dial")
			}
			if err := conn.Do(ctx, ch.Query{Body: `DROP TABLE IF EXISTS ch_insert_lag`}); err != nil {
				return errors.Wrap(err, "drop table")
			}
			if err := conn.Do(ctx, ch.Query{Body: ddl}); err != nil {
				return errors.Wrap(err, "create")
			}
			close(ready)
			data := make(proto.ColDateTime64, 50_000)
			fill := func() {
				now := proto.ToDateTime64(time.Now(), precision)
				for i := range data {
					data[i] = now
				}
			}
			fill()
			return conn.Do(ctx, ch.Query{
				Body: `INSERT INTO ch_insert_lag VALUES`,
				OnInput: func(ctx context.Context) error {
					time.Sleep(time.Millisecond * 20)

					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-done:
						return io.EOF
					default:
						fill()
						return nil
					}
				},
				Input: proto.Input{
					{Name: "ts", Data: data.Wrap(precision)},
				},
			})
		})
		g.Go(func() error {
			defer close(done)
			conn, err := ch.Dial(ctx, ch.Options{})
			if err != nil {
				return errors.Wrap(err, "dial")
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ready:
			}
			ticker := time.NewTicker(time.Millisecond * 300)

			var data proto.ColDateTime64
			for range ticker.C {
				if err := conn.Do(ctx, ch.Query{
					Body: `SELECT max(ts) as latest FROM ch_insert_lag`,
					Result: proto.Results{
						{Name: "latest", Data: &data},
					},
				}); err != nil {
					return errors.Wrap(err, "select")
				}
				if len(data) == 0 {
					continue
				}
				v := data[0]
				if v == 0 {
					continue
				}
				latest := v.Time(precision)
				lag := time.Since(latest)
				fmt.Println(lag.Round(time.Millisecond))
			}
			return nil
		})
		return g.Wait()
	})
}
