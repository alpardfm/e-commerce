package refund

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
	GetList(ctx context.Context, param entity.Refund, opts ...func(prefix, suffix *string) error) ([]entity.Refund, error)
	GetDetail(ctx context.Context, param entity.Refund, opts ...func(prefix, suffix *string) error) (entity.Refund, error)
	Create(ctx context.Context, param entity.Refund) (entity.Refund, error)
	Update(ctx context.Context, param entity.Refund) (entity.Refund, error)
	Delete(ctx context.Context, param entity.Refund) (entity.Refund, error)
}

type refund struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &refund{
		log: log,
		db:  db,
	}
}

func (r *refund) GetList(ctx context.Context, param entity.Refund, opts ...func(prefix, suffix *string) error) ([]entity.Refund, error) {
	qb, err := query.NewSQLQueryBuilder(r.db, "param", "db")
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

	rows, err := r.db.Follower().Query(ctx, "getListRefund", readRefund+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Refund{}
	for rows.Next() {
		result := entity.Refund{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *refund) GetDetail(ctx context.Context, param entity.Refund, opts ...func(prefix, suffix *string) error) (entity.Refund, error) {
	qb, err := query.NewSQLQueryBuilder(r.db, "param", "db")
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Refund{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := r.db.Follower().QueryRow(ctx, "getDetailRefund", readRefund+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Refund{}
	if err := row.StructScan(&result); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (r *refund) Create(ctx context.Context, param entity.Refund) (entity.Refund, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txCreateRefund", sql.TxOptions{})
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createRefund", createRefund, param)
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no refund created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (r *refund) Update(ctx context.Context, param entity.Refund) (entity.Refund, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txUpdateRefund", sql.TxOptions{})
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateRefund", updateRefund, param)
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no refund updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (r *refund) Delete(ctx context.Context, param entity.Refund) (entity.Refund, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txDeleteRefund", sql.TxOptions{})
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteRefund", deleteRefund, param)
	if err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no refund deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Refund{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
