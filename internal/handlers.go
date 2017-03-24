package internal

import (
	middleware "github.com/go-openapi/runtime/middleware"

	op "github.com/Magicking/Biffin/restapi/operations"
	"log"
)

func GetMap(ctx *Context, params op.GetMapParams) middleware.Responder {
	map_files, err := GetAllMapFile(ctx)
	if err != nil {
		log.Println(err)
		return op.NewGetMapDefault(500)
	}
	ret := op.NewGetMapOK()

	return ret.WithPayload(map_files)
}

func PostMap(ctx *Context, params op.PostMapParams) middleware.Responder {
	err := InsertMapFile(ctx, params.MapFile)
	if err != nil {
		log.Println(err)
		return op.NewPostMapDefault(500)
	}
	return op.NewPostMapOK()
}
