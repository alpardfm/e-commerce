package payments

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
	GetList(ctx context.Context, param entity.Payments, opts ...func(prefix, suffix *string) error) ([]entity.Payments, error)
	GetDetail(ctx context.Context, param entity.Payments, opts ...func(prefix, suffix *string) error) (entity.Payments, error)
	Create(ctx context.Context, param entity.Payments) (entity.Payments, error)
	Update(ctx context.Context, param entity.Payments) (entity.Payments, error)
	Delete(ctx context.Context, param entity.Payments) (entity.Payments, error)
}

type payments struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &payments{
		log: log,
		db:  db,
	}
}

func (p *payments) GetList(ctx context.Context, param entity.Payments, opts ...func(prefix, suffix *string) error) ([]entity.Payments, error) {
	qb, err := query.NewSQLQueryBuilder(p.db, "param", "db")
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

	rows, err := p.db.Follower().Query(ctx, "getListPayments", readPayments+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Payments{}
	for rows.Next() {
		result := entity.Payments{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (p *payments) GetDetail(ctx context.Context, param entity.Payments, opts ...func(prefix, suffix *string) error) (entity.Payments, error) {
	qb, err := query.NewSQLQueryBuilder(p.db, "param", "db")
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Payments{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := p.db.Follower().QueryRow(ctx, "getDetailPayments", readPayments+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Payments{}
	if err := row.StructScan(&result); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (p *payments) Create(ctx context.Context, param entity.Payments) (entity.Payments, error) {
	tx, err := p.db.Leader().BeginTx(ctx, "txCreatePayments", sql.TxOptions{})
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createPayments", createPayments, param)
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no payments created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (p *payments) Update(ctx context.Context, param entity.Payments) (entity.Payments, error) {
	tx, err := p.db.Leader().BeginTx(ctx, "txUpdatePayments", sql.TxOptions{})
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updatePayments", updatePayments, param)
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no payments updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (p *payments) Delete(ctx context.Context, param entity.Payments) (entity.Payments, error) {
	tx, err := p.db.Leader().BeginTx(ctx, "txDeletePayments", sql.TxOptions{})
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deletePayments", deletePayments, param)
	if err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no payments deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Payments{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
