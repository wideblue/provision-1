package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchTemplateHandlerFunc turns a function with the right signature into a patch template handler
type PatchTemplateHandlerFunc func(PatchTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchTemplateHandlerFunc) Handle(params PatchTemplateParams) middleware.Responder {
	return fn(params)
}

// PatchTemplateHandler interface for that can handle valid patch template params
type PatchTemplateHandler interface {
	Handle(PatchTemplateParams) middleware.Responder
}

// NewPatchTemplate creates a new http.Handler for the patch template operation
func NewPatchTemplate(ctx *middleware.Context, handler PatchTemplateHandler) *PatchTemplate {
	return &PatchTemplate{Context: ctx, Handler: handler}
}

/*PatchTemplate swagger:route PATCH /template/{uuid} Templates patchTemplate

Patch Template

*/
type PatchTemplate struct {
	Context *middleware.Context
	Handler PatchTemplateHandler
}

func (o *PatchTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPatchTemplateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}