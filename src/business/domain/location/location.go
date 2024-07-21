package location

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
	GetList(ctx context.Context, param entity.Location, opts ...func(prefix, suffix *string) error) ([]entity.Location, error)
	GetDetail(ctx context.Context, param entity.Location, opts ...func(prefix, suffix *string) error) (entity.Location, error)
	Create(ctx context.Context, param entity.Location) (entity.Location, error)
	Update(ctx context.Context, param entity.Location) (entity.Location, error)
	Delete(ctx context.Context, param entity.Location) (entity.Location, error)
}

type location struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &location{
		log: log,
		db:  db,
	}
}

func (l *location) GetList(ctx context.Context, param entity.Location, opts ...func(prefix, suffix *string) error) ([]entity.Location, error) {
	qb, err := query.NewSQLQueryBuilder(l.db, "param", "db")
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

	rows, err := l.db.Follower().Query(ctx, "getListLocation", readLocation+additionalQuery, additionalArgs...)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	results := []entity.Location{}
	for rows.Next() {
		result := entity.Location{}
		if err := rows.StructScan(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (l *location) GetDetail(ctx context.Context, param entity.Location, opts ...func(prefix, suffix *string) error) (entity.Location, error) {
	qb, err := query.NewSQLQueryBuilder(l.db, "param", "db")
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	prefix, suffix := "", ""
	for _, opt := range opts {
		if err := opt(&prefix, &suffix); err != nil {
			return entity.Location{}, err
		}
	}

	qb.AddPrefixQuery(prefix)
	qb.AddSuffixQuery(suffix)

	additionalQuery, additionalArgs, _, _, err := qb.Build(&param)
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLBuilder, err.Error())
	}

	row, err := l.db.Follower().QueryRow(ctx, "getDetailLocation", readLocation+additionalQuery, additionalArgs...)
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLRead, err.Error())
	}

	result := entity.Location{}
	if err := row.StructScan(&result); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLRowScan, err.Error())
	}

	return result, nil
}

func (l *location) Create(ctx context.Context, param entity.Location) (entity.Location, error) {
	tx, err := l.db.Leader().BeginTx(ctx, "txCreateLocation", sql.TxOptions{})
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("createLocation", createLocation, param)
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no location created")
	}

	if err := tx.Commit(); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	if param.ID, err = res.LastInsertId(); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	}

	return param, nil
}

func (l *location) Update(ctx context.Context, param entity.Location) (entity.Location, error) {
	tx, err := l.db.Leader().BeginTx(ctx, "txUpdateLocation", sql.TxOptions{})
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("updateCategories", updateLocation, param)
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no location updated")
	}

	if err := tx.Commit(); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}

func (l *location) Delete(ctx context.Context, param entity.Location) (entity.Location, error) {
	tx, err := l.db.Leader().BeginTx(ctx, "txDeleteLocation", sql.TxOptions{})
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxBegin, err.Error())
	}
	defer tx.Rollback()

	res, err := tx.NamedExec("deleteLocation", deleteLocation, param)
	if err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxExec, err.Error())
	}

	if num, err := res.RowsAffected(); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, err.Error())
	} else if num < 1 {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLNoRowsAffected, "no location deleted")
	}

	if err := tx.Commit(); err != nil {
		return entity.Location{}, errors.NewWithCode(codes.CodeSQLTxCommit, err.Error())
	}

	return param, nil
}
