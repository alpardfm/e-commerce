package order_items

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
	GetList(ctx context.Context, param entity.OrderItems, opts ...func(prefix, suffix *string) error) ([]entity.OrderItems, error)
	GetDetail(ctx context.Context, param entity.OrderItems, opts ...func(prefix, suffix *string) error) (entity.OrderItems, error)
	Create(ctx context.Context, param entity.OrderItems) (entity.OrderItems, error)
	Update(ctx context.Context, param entity.OrderItems) (entity.OrderItems, error)
	Delete(ctx context.Context, param entity.OrderItems) (entity.OrderItems, error)
}

type orderItems struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &orderItems{
		log: log,
		db:  db,
	}
}

func (o *orderItems) GetList(ctx context.Context, param entity.OrderItems, opts ...func(prefix, suffix *string) error) ([]entity.OrderItems, error) {
	qb, err := query.NewSQLQueryBuilder(o.db, "param", "db")
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

	rows, err := o.db.Follower().Query(ctx, "getListOrderItems", readOrderItems+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.OrderItems{}
	for rows.Next() {
		result := entity.OrderItems{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (o *orderItems) GetDetail(ctx context.Context, param entity.OrderItems, opts ...func(prefix, suffix *string) error) (entity.OrderItems, error) {
	qb, err := query.NewSQLQueryBuilder(o.db, "param", "db")
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.OrderItems{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := o.db.Follower().QueryRow(ctx, "getDetailOrderItems", readOrderItems+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.OrderItems{}
	if err := row.StructScan(&result); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (o *orderItems) Create(ctx context.Context, param entity.OrderItems) (entity.OrderItems, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txCreateOrderItems", sql.TxOptions{})
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createOrderItems", createOrderItems, param)
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no order items created")
	}

	if err := tx.Commit(); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (o *orderItems) Update(ctx context.Context, param entity.OrderItems) (entity.OrderItems, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txUpdateOrderItems", sql.TxOptions{})
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateOrderItems", updateOrderItems, param)
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no order items updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (o *orderItems) Delete(ctx context.Context, param entity.OrderItems) (entity.OrderItems, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txDeleteOrderItems", sql.TxOptions{})
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteOrderItems", deleteOrderItems, param)
	if err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no order items deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.OrderItems{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
