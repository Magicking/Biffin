package internal

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
)

type key int

var ctxValKey key = 0
var dbKey key = 1

type ctxValues struct {
	m map[key]interface{}
}

func (v *ctxValues) Set(k key, val interface{}) {
	v.m[k] = val
}

func (v *ctxValues) Get(k key) interface{} {
	val, ok := v.m[k]
	if !ok {
		log.Fatalf("Could not find key: %v", k)
	}
	return val
}

func NewCtxValues() *ctxValues {
	mm := make(map[key]interface{})
	cv := &ctxValues{
		m: mm,
	}
	return cv
}

func getContextValue(ctx context.Context, k key) interface{} {
	v, ok := ctx.Value(ctxValKey).(*ctxValues)
	if !ok {
		log.Fatalf("Could not obtain map context values")
	}
	return v.Get(k)
}

func setContextValue(ctx context.Context, k key, val interface{}) {
	v, ok := ctx.Value(ctxValKey).(*ctxValues)
	if !ok {
		log.Fatalf("Could not obtain map context values")
	}
	v.Set(k, val)
}

func InitContext(ctx context.Context) context.Context {
	values := NewCtxValues()
	return context.WithValue(ctx, ctxValKey, values)
}

func NewDBToContext(ctx context.Context, dbDsn string) {
	db, err := InitDatabase(dbDsn)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	setContextValue(ctx, dbKey, db)
}

func DBFromContext(ctx context.Context) *gorm.DB {
	key := dbKey
	ret, ok := getContextValue(ctx, key).(*gorm.DB)
	if !ok {
		log.Fatalf("Could not cast context with key %d", key)
	}
	return ret
}
