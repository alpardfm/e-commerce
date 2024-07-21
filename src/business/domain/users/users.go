package users

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
	GetList(ctx context.Context, param entity.Users, opts ...func(prefix, suffix *string) error) ([]entity.Users, error)
	GetDetail(ctx context.Context, param entity.Users, opts ...func(prefix, suffix *string) error) (entity.Users, error)
	Create(ctx context.Context, param entity.Users) (entity.Users, error)
	Update(ctx context.Context, param entity.Users) (entity.Users, error)
	Delete(ctx context.Context, param entity.Users) (entity.Users, error)
}

type users struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &users{
		log: log,
		db:  db,
	}
}

func (u *users) GetList(ctx context.Context, param entity.Users, opts ...func(prefix, suffix *string) error) ([]entity.Users, error) {
	qb, err := query.NewSQLQueryBuilder(u.db, "param", "db")
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

	rows, err := u.db.Follower().Query(ctx, "getListUsers", readUsers+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Users{}
	for rows.Next() {
		result := entity.Users{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (u *users) GetDetail(ctx context.Context, param entity.Users, opts ...func(prefix, suffix *string) error) (entity.Users, error) {
	qb, err := query.NewSQLQueryBuilder(u.db, "param", "db")
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Users{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := u.db.Follower().QueryRow(ctx, "getDetailUsers", readUsers+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Users{}
	if err := row.StructScan(&result); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}
func (u *users) Create(ctx context.Context, param entity.Users) (entity.Users, error) {
	tx, err := u.db.Leader().BeginTx(ctx, "txCreateUsers", sql.TxOptions{})
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createUsers", createUsers, param)
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no users created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}
func (u *users) Update(ctx context.Context, param entity.Users) (entity.Users, error) {
	tx, err := u.db.Leader().BeginTx(ctx, "txUpdateUsers", sql.TxOptions{})
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateUsers", updateUsers, param)
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no users updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
func (u *users) Delete(ctx context.Context, param entity.Users) (entity.Users, error) {
	tx, err := u.db.Leader().BeginTx(ctx, "txDeleteUsers", sql.TxOptions{})
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteUsers", deleteUsers, param)
	if err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no users deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Users{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
