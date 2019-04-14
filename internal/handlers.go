package internal

import (
	"context"

	middleware "github.com/go-openapi/runtime/middleware"

	"log"

	op "github.com/Magicking/Biffin/restapi/operations"
)

func GetMap(ctx context.Context, params op.GetMapParams) middleware.Responder {
	map_files, err := GetAllMapFile(ctx)
	if err != nil {
		log.Println(err)
		return op.NewGetMapDefault(500)
	}
	ret := op.NewGetMapOK()

	return ret.WithPayload(map_files)
}

func PostMap(ctx context.Context, params op.PostMapParams) middleware.Responder {
	err := InsertMapFile(ctx, params.MapFile)
	if err != nil {
		log.Println(err)
		return op.NewPostMapDefault(500)
	}
	return op.NewPostMapOK()
}
