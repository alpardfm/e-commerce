package orders

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
	GetList(ctx context.Context, param entity.Orders, opts ...func(prefix, suffix *string) error) ([]entity.Orders, error)
	GetDetail(ctx context.Context, param entity.Orders, opts ...func(prefix, suffix *string) error) (entity.Orders, error)
	Create(ctx context.Context, param entity.Orders) (entity.Orders, error)
	Update(ctx context.Context, param entity.Orders) (entity.Orders, error)
	Delete(ctx context.Context, param entity.Orders) (entity.Orders, error)
}

type orders struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &orders{
		log: log,
		db:  db,
	}
}

func (o *orders) GetList(ctx context.Context, param entity.Orders, opts ...func(prefix, suffix *string) error) ([]entity.Orders, error) {
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

	rows, err := o.db.Follower().Query(ctx, "getListOrders", readOrders+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Orders{}
	for rows.Next() {
		result := entity.Orders{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (o *orders) GetDetail(ctx context.Context, param entity.Orders, opts ...func(prefix, suffix *string) error) (entity.Orders, error) {
	qb, err := query.NewSQLQueryBuilder(o.db, "param", "db")
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Orders{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := o.db.Follower().QueryRow(ctx, "getDetailOrders", readOrders+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Orders{}
	if err := row.StructScan(&result); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (o *orders) Create(ctx context.Context, param entity.Orders) (entity.Orders, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txCreateOrders", sql.TxOptions{})
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createOrders", createOrders, param)
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no orders created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (o *orders) Update(ctx context.Context, param entity.Orders) (entity.Orders, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txUpdateOrders", sql.TxOptions{})
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateOrders", updateOrders, param)
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no orders updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (o *orders) Delete(ctx context.Context, param entity.Orders) (entity.Orders, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txDeleteOrders", sql.TxOptions{})
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteOrders", deleteOrders, param)
	if err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no orders deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Orders{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
