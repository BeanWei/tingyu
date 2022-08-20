package driver

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/BeanWei/tingyu/g"
)

var _ dialect.Driver = (*multiDriver)(nil)

type multiDriver struct {
	r, w dialect.Driver
}

func NewDriver() *multiDriver {
	return &multiDriver{
		r: entsql.OpenDB(dialect.Postgres, g.RDB()),
		w: entsql.OpenDB(dialect.Postgres, g.WDB()),
	}
}

func (d *multiDriver) Dialect() string {
	return dialect.Postgres
}

func (d *multiDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return d.r.Query(ctx, query, args, v)
}

func (d *multiDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return d.w.Exec(ctx, query, args, v)
}

func (d *multiDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.w.Tx(ctx)
}

func (d *multiDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return d.w.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

func (d *multiDriver) Close() error {
	rerr := d.r.Close()
	werr := d.w.Close()
	if rerr != nil {
		return rerr
	}
	if werr != nil {
		return werr
	}
	return nil
}
