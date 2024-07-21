package cart

import (
	"context"

	"github.com/alpardfm/e-commerce/src/entity"
	"github.com/alpardfm/go-toolkit/codes"
	"github.com/alpardfm/go-toolkit/errors"
	"github.com/alpardfm/go-toolkit/log"
	"github.com/alpardfm/go-toolkit/query"
	"github.com/alpardfm/go-toolkit/sql"
)

type Interface interface {
	GetList(ctx context.Context, param entity.Cart, opts ...func(prefix, suffix *string) error) ([]entity.Cart, error)
	GetDetail(ctx context.Context, param entity.Cart, opts ...func(prefix, suffix *string) error) (entity.Cart, error)
	Create(ctx context.Context, param entity.Cart) (entity.Cart, error)
	Update(ctx context.Context, param entity.Cart) (entity.Cart, error)
	Delete(ctx context.Context, param entity.Cart) (entity.Cart, error)
}

type cart struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &cart{
		log: log,
		db:  db,
	}
}

func (c *cart) GetList(ctx context.Context, param entity.Cart, opts ...func(prefix, suffix *string) error) ([]entity.Cart, error) {
	qb, err := query.NewSQLQueryBuilder(c.db, "param", "db")
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return nil, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return nil, err
	}

	rows, err := c.db.Follower().Query(ctx, "getListCart", readCart+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Cart{}
	for rows.Next() {
		result := entity.Cart{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (c *cart) GetDetail(ctx context.Context, param entity.Cart, opts ...func(prefix, suffix *string) error) (entity.Cart, error) {
	qb, err := query.NewSQLQueryBuilder(c.db, "param", "db")
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Cart{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := c.db.Follower().QueryRow(ctx, "getDetailCart", readCart+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Cart{}
	if err := row.StructScan(&result); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (c *cart) Create(ctx context.Context, param entity.Cart) (entity.Cart, error) {
	tx, err := c.db.Leader().BeginTx(ctx, "txCreateCart", sql.TxOptions{})
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createCart", createCart, param)
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no cart created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (c *cart) Update(ctx context.Context, param entity.Cart) (entity.Cart, error) {
	tx, err := c.db.Leader().BeginTx(ctx, "txUpdateCart", sql.TxOptions{})
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateCart", updateCart, param)
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no cart updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (c *cart) Delete(ctx context.Context, param entity.Cart) (entity.Cart, error) {
	tx, err := c.db.Leader().BeginTx(ctx, "txDeleteCart", sql.TxOptions{})
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteCart", deleteCart, param)
	if err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no cart deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Cart{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
