package reviews

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
	GetList(ctx context.Context, param entity.Reviews, opts ...func(prefix, suffix *string) error) ([]entity.Reviews, error)
	GetDetail(ctx context.Context, param entity.Reviews, opts ...func(prefix, suffix *string) error) (entity.Reviews, error)
	Create(ctx context.Context, param entity.Reviews) (entity.Reviews, error)
	Update(ctx context.Context, param entity.Reviews) (entity.Reviews, error)
	Delete(ctx context.Context, param entity.Reviews) (entity.Reviews, error)
}

type reviews struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &reviews{
		log: log,
		db:  db,
	}
}

func (r *reviews) GetList(ctx context.Context, param entity.Reviews, opts ...func(prefix, suffix *string) error) ([]entity.Reviews, error) {
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

	rows, err := r.db.Follower().Query(ctx, "getListReviews", readReviews+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Reviews{}
	for rows.Next() {
		result := entity.Reviews{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *reviews) GetDetail(ctx context.Context, param entity.Reviews, opts ...func(prefix, suffix *string) error) (entity.Reviews, error) {
	qb, err := query.NewSQLQueryBuilder(r.db, "param", "db")
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Reviews{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := r.db.Follower().QueryRow(ctx, "getDetailReviews", readReviews+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Reviews{}
	if err := row.StructScan(&result); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (r *reviews) Create(ctx context.Context, param entity.Reviews) (entity.Reviews, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txCreateReviews", sql.TxOptions{})
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createReviews", createReviews, param)
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no reviews created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (r *reviews) Update(ctx context.Context, param entity.Reviews) (entity.Reviews, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txUpdateReviews", sql.TxOptions{})
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateReviews", updateReviews, param)
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no reviews updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (r *reviews) Delete(ctx context.Context, param entity.Reviews) (entity.Reviews, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txDeleteReviews", sql.TxOptions{})
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteReviews", deleteReviews, param)
	if err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no reviews deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Reviews{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
