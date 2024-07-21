package categories

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
	GetList(ctx context.Context, param entity.Categories, opts ...func(prefix, suffix *string) error) ([]entity.Categories, error)
	GetDetail(ctx context.Context, param entity.Categories, opts ...func(prefix, suffix *string) error) (entity.Categories, error)
	Create(ctx context.Context, param entity.Categories) (entity.Categories, error)
	Update(ctx context.Context, param entity.Categories) (entity.Categories, error)
	Delete(ctx context.Context, param entity.Categories) (entity.Categories, error)
}

type categories struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &categories{
		log: log,
		db:  db,
	}
}

func (c *categories) GetList(ctx context.Context, param entity.Categories, opts ...func(prefix, suffix *string) error) ([]entity.Categories, error) {
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

	rows, err := c.db.Follower().Query(ctx, "getListCategories", readCategories+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Categories{}
	for rows.Next() {
		result := entity.Categories{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (c *categories) GetDetail(ctx context.Context, param entity.Categories, opts ...func(prefix, suffix *string) error) (entity.Categories, error) {
	qb, err := query.NewSQLQueryBuilder(c.db, "param", "db")
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Categories{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := c.db.Follower().QueryRow(ctx, "getDetailCategories", readCategories+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Categories{}
	if err := row.StructScan(&result); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (c *categories) Create(ctx context.Context, param entity.Categories) (entity.Categories, error) {
	tx, err := c.db.Leader().BeginTx(ctx, "txCreateCategories", sql.TxOptions{})
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createCategories", createCategories, param)
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no categories created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (c *categories) Update(ctx context.Context, param entity.Categories) (entity.Categories, error) {
	tx, err := c.db.Leader().BeginTx(ctx, "txUpdateCategories", sql.TxOptions{})
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateCategories", updateCategories, param)
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no categories updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (c *categories) Delete(ctx context.Context, param entity.Categories) (entity.Categories, error) {
	tx, err := c.db.Leader().BeginTx(ctx, "txDeleteCategories", sql.TxOptions{})
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteCategories", deleteCategories, param)
	if err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no categories deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Categories{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
