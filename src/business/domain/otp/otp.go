package otp

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
	GetList(ctx context.Context, param entity.OTP, opts ...func(prefix, suffix *string) error) ([]entity.OTP, error)
	GetDetail(ctx context.Context, param entity.OTP, opts ...func(prefix, suffix *string) error) (entity.OTP, error)
	Create(ctx context.Context, param entity.OTP) (entity.OTP, error)
	Update(ctx context.Context, param entity.OTP) (entity.OTP, error)
	Delete(ctx context.Context, param entity.OTP) (entity.OTP, error)
}

type otp struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &otp{
		log: log,
		db:  db,
	}
}

func (o *otp) GetList(ctx context.Context, param entity.OTP, opts ...func(prefix, suffix *string) error) ([]entity.OTP, error) {
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

	rows, err := o.db.Follower().Query(ctx, "getListOTP", readOTP+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.OTP{}
	for rows.Next() {
		result := entity.OTP{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (o *otp) GetDetail(ctx context.Context, param entity.OTP, opts ...func(prefix, suffix *string) error) (entity.OTP, error) {
	qb, err := query.NewSQLQueryBuilder(o.db, "param", "db")
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.OTP{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := o.db.Follower().QueryRow(ctx, "getDetailOTP", readOTP+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.OTP{}
	if err := row.StructScan(&result); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (o *otp) Create(ctx context.Context, param entity.OTP) (entity.OTP, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txCreateOTP", sql.TxOptions{})
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createOTP", createOTP, param)
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no otp created")
	}

	if err := tx.Commit(); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (o *otp) Update(ctx context.Context, param entity.OTP) (entity.OTP, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txUpdateOTP", sql.TxOptions{})
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateOTP", updateOTP, param)
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no otp updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (o *otp) Delete(ctx context.Context, param entity.OTP) (entity.OTP, error) {
	tx, err := o.db.Leader().BeginTx(ctx, "txDeleteOTP", sql.TxOptions{})
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteOTP", deleteOTP, param)
	if err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no otp deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.OTP{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
