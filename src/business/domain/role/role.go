package role

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
	GetList(ctx context.Context, param entity.Role, opts ...func(prefix, suffix *string) error) ([]entity.Role, error)
	GetDetail(ctx context.Context, param entity.Role, opts ...func(prefix, suffix *string) error) (entity.Role, error)
	Create(ctx context.Context, param entity.Role) (entity.Role, error)
	Update(ctx context.Context, param entity.Role) (entity.Role, error)
	Delete(ctx context.Context, param entity.Role) (entity.Role, error)
}

type role struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &role{
		log: log,
		db:  db,
	}
}

func (r *role) GetList(ctx context.Context, param entity.Role, opts ...func(prefix, suffix *string) error) ([]entity.Role, error) {
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

	rows, err := r.db.Follower().Query(ctx, "getListRole", readRole+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Role{}
	for rows.Next() {
		result := entity.Role{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *role) GetDetail(ctx context.Context, param entity.Role, opts ...func(prefix, suffix *string) error) (entity.Role, error) {
	qb, err := query.NewSQLQueryBuilder(r.db, "param", "db")
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Role{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := r.db.Follower().QueryRow(ctx, "getDetailRole", readRole+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Role{}
	if err := row.StructScan(&result); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (r *role) Create(ctx context.Context, param entity.Role) (entity.Role, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txCreateRole", sql.TxOptions{})
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createRole", createRole, param)
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no role created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (r *role) Update(ctx context.Context, param entity.Role) (entity.Role, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txUpdateRole", sql.TxOptions{})
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateRole", updateRole, param)
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no role updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (r *role) Delete(ctx context.Context, param entity.Role) (entity.Role, error) {
	tx, err := r.db.Leader().BeginTx(ctx, "txDeleteRole", sql.TxOptions{})
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteRole", deleteRole, param)
	if err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no role deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Role{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
