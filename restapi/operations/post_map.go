// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostMapHandlerFunc turns a function with the right signature into a post map handler
type PostMapHandlerFunc func(PostMapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostMapHandlerFunc) Handle(params PostMapParams) middleware.Responder {
	return fn(params)
}

// PostMapHandler interface for that can handle valid post map params
type PostMapHandler interface {
	Handle(PostMapParams) middleware.Responder
}

// NewPostMap creates a new http.Handler for the post map operation
func NewPostMap(ctx *middleware.Context, handler PostMapHandler) *PostMap {
	return &PostMap{Context: ctx, Handler: handler}
}

/*PostMap swagger:route POST /map postMap

Create a new map

Create the new map, return the associated id, an error othewise


*/
type PostMap struct {
	Context *middleware.Context
	Handler PostMapHandler
}

func (o *PostMap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostMapParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
