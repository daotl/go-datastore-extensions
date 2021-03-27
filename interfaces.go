package dsextensions

import (
	"github.com/daotl/go-datastore"
	"github.com/daotl/go-datastore/key"
	"github.com/daotl/go-datastore/query"
)

type QueryExt struct {
	query.Query
	SeekPrefix key.Key
}

type TxnExt interface {
	datastore.Txn
	QueryExtensions
}

type DatastoreExtensions interface {
	NewTransactionExtended(readOnly bool) (TxnExt, error)
	QueryExtensions
}

type QueryExtensions interface {
	QueryExtended(q QueryExt) (query.Results, error)
}
