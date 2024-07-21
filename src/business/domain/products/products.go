package products

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
	GetList(ctx context.Context, param entity.Products, opts ...func(prefix, suffix *string) error) ([]entity.Products, error)
	GetDetail(ctx context.Context, param entity.Products, opts ...func(prefix, suffix *string) error) (entity.Products, error)
	Create(ctx context.Context, param entity.Products) (entity.Products, error)
	Update(ctx context.Context, param entity.Products) (entity.Products, error)
	Delete(ctx context.Context, param entity.Products) (entity.Products, error)
}

type products struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &products{
		log: log,
		db:  db,
	}
}

func (p *products) GetList(ctx context.Context, param entity.Products, opts ...func(prefix, suffix *string) error) ([]entity.Products, error) {
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

	rows, err := p.db.Follower().Query(ctx, "getListProducts", readProducts+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Products{}
	for rows.Next() {
		result := entity.Products{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (p *products) GetDetail(ctx context.Context, param entity.Products, opts ...func(prefix, suffix *string) error) (entity.Products, error) {
	qb, err := query.NewSQLQueryBuilder(p.db, "param", "db")
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Products{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := p.db.Follower().QueryRow(ctx, "getDetailProducts", readProducts+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Products{}
	if err := row.StructScan(&result); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (p *products) Create(ctx context.Context, param entity.Products) (entity.Products, error) {
	tx, err := p.db.Leader().BeginTx(ctx, "txCreateProducts", sql.TxOptions{})
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createProducts", createProducts, param)
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no products created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (p *products) Update(ctx context.Context, param entity.Products) (entity.Products, error) {
	tx, err := p.db.Leader().BeginTx(ctx, "txUpdateProducts", sql.TxOptions{})
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateProducts", updateProducts, param)
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no products updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (p *products) Delete(ctx context.Context, param entity.Products) (entity.Products, error) {
	tx, err := p.db.Leader().BeginTx(ctx, "txDeleteProducts", sql.TxOptions{})
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteProducts", deleteProducts, param)
	if err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no users deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Products{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
