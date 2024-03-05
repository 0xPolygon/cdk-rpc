package rpc

import (
	"context"
	"database/sql/driver"

	"github.com/jmoiron/sqlx"
)

// DBTxManager allows to do scopped DB txs
type DBTxManager struct{}

// Tx is the interface that defines functions a db tx has to implement
type dbTx interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
	driver.Tx
}

// DB defines functions that a DB instance should implement
type DB interface {
	BeginStateTransaction(ctx context.Context) (dbTx, error)
}

// DBTxScopedFn function to do scopped DB txs
type DBTxScopedFn func(ctx context.Context, dbTx dbTx) (interface{}, Error)

// NewDbTxScope function to initiate DB scopped txs
func (f *DBTxManager) NewDbTxScope(db DB, scopedFn DBTxScopedFn) (interface{}, Error) {
	ctx := context.Background()

	dbTx, err := db.BeginStateTransaction(ctx)
	if err != nil {
		return RPCErrorResponse(DefaultErrorCode, "failed to connect to the state", err)
	}

	v, rpcErr := scopedFn(ctx, dbTx)
	if rpcErr != nil {
		if txErr := dbTx.Rollback(); txErr != nil {
			return RPCErrorResponse(DefaultErrorCode, "failed to rollback db transaction", txErr)
		}
		return v, rpcErr
	}

	if txErr := dbTx.Commit(); txErr != nil {
		return RPCErrorResponse(DefaultErrorCode, "failed to commit db transaction", txErr)
	}
	return v, rpcErr
}
