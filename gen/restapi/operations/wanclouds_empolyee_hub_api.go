// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewWancloudsEmpolyeeHubAPI creates a new WancloudsEmpolyeeHub instance
func NewWancloudsEmpolyeeHubAPI(spec *loads.Document) *WancloudsEmpolyeeHubAPI {
	return &WancloudsEmpolyeeHubAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		AddEmployeeHandler: AddEmployeeHandlerFunc(func(params AddEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation AddEmployee has not yet been implemented")
		}),
		AddHomeHandler: AddHomeHandlerFunc(func(params AddHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation AddHome has not yet been implemented")
		}),
		AddOfficeHandler: AddOfficeHandlerFunc(func(params AddOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation AddOffice has not yet been implemented")
		}),
		DeleteEmployeeHandler: DeleteEmployeeHandlerFunc(func(params DeleteEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteEmployee has not yet been implemented")
		}),
		DeleteHomeHandler: DeleteHomeHandlerFunc(func(params DeleteHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteHome has not yet been implemented")
		}),
		DeleteOfficeHandler: DeleteOfficeHandlerFunc(func(params DeleteOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteOffice has not yet been implemented")
		}),
		GetEmployeeHandler: GetEmployeeHandlerFunc(func(params GetEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation GetEmployee has not yet been implemented")
		}),
		GetHomeHandler: GetHomeHandlerFunc(func(params GetHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHome has not yet been implemented")
		}),
		GetOfficeHandler: GetOfficeHandlerFunc(func(params GetOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation GetOffice has not yet been implemented")
		}),
		ListEmployeesHandler: ListEmployeesHandlerFunc(func(params ListEmployeesParams) middleware.Responder {
			return middleware.NotImplemented("operation ListEmployees has not yet been implemented")
		}),
		ListHomesHandler: ListHomesHandlerFunc(func(params ListHomesParams) middleware.Responder {
			return middleware.NotImplemented("operation ListHomes has not yet been implemented")
		}),
		ListOfficesHandler: ListOfficesHandlerFunc(func(params ListOfficesParams) middleware.Responder {
			return middleware.NotImplemented("operation ListOffices has not yet been implemented")
		}),
		UpdateEmployeeHandler: UpdateEmployeeHandlerFunc(func(params UpdateEmployeeParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateEmployee has not yet been implemented")
		}),
		UpdateHomeHandler: UpdateHomeHandlerFunc(func(params UpdateHomeParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateHome has not yet been implemented")
		}),
		UpdateOfficeHandler: UpdateOfficeHandlerFunc(func(params UpdateOfficeParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateOffice has not yet been implemented")
		}),
	}
}

/*WancloudsEmpolyeeHubAPI Employee Hub API */
type WancloudsEmpolyeeHubAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AddEmployeeHandler sets the operation handler for the add employee operation
	AddEmployeeHandler AddEmployeeHandler
	// AddHomeHandler sets the operation handler for the add home operation
	AddHomeHandler AddHomeHandler
	// AddOfficeHandler sets the operation handler for the add office operation
	AddOfficeHandler AddOfficeHandler
	// DeleteEmployeeHandler sets the operation handler for the delete employee operation
	DeleteEmployeeHandler DeleteEmployeeHandler
	// DeleteHomeHandler sets the operation handler for the delete home operation
	DeleteHomeHandler DeleteHomeHandler
	// DeleteOfficeHandler sets the operation handler for the delete office operation
	DeleteOfficeHandler DeleteOfficeHandler
	// GetEmployeeHandler sets the operation handler for the get employee operation
	GetEmployeeHandler GetEmployeeHandler
	// GetHomeHandler sets the operation handler for the get home operation
	GetHomeHandler GetHomeHandler
	// GetOfficeHandler sets the operation handler for the get office operation
	GetOfficeHandler GetOfficeHandler
	// ListEmployeesHandler sets the operation handler for the list employees operation
	ListEmployeesHandler ListEmployeesHandler
	// ListHomesHandler sets the operation handler for the list homes operation
	ListHomesHandler ListHomesHandler
	// ListOfficesHandler sets the operation handler for the list offices operation
	ListOfficesHandler ListOfficesHandler
	// UpdateEmployeeHandler sets the operation handler for the update employee operation
	UpdateEmployeeHandler UpdateEmployeeHandler
	// UpdateHomeHandler sets the operation handler for the update home operation
	UpdateHomeHandler UpdateHomeHandler
	// UpdateOfficeHandler sets the operation handler for the update office operation
	UpdateOfficeHandler UpdateOfficeHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *WancloudsEmpolyeeHubAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *WancloudsEmpolyeeHubAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *WancloudsEmpolyeeHubAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WancloudsEmpolyeeHubAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WancloudsEmpolyeeHubAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WancloudsEmpolyeeHubAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WancloudsEmpolyeeHubAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WancloudsEmpolyeeHubAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WancloudsEmpolyeeHubAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WancloudsEmpolyeeHubAPI
func (o *WancloudsEmpolyeeHubAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AddEmployeeHandler == nil {
		unregistered = append(unregistered, "AddEmployeeHandler")
	}
	if o.AddHomeHandler == nil {
		unregistered = append(unregistered, "AddHomeHandler")
	}
	if o.AddOfficeHandler == nil {
		unregistered = append(unregistered, "AddOfficeHandler")
	}
	if o.DeleteEmployeeHandler == nil {
		unregistered = append(unregistered, "DeleteEmployeeHandler")
	}
	if o.DeleteHomeHandler == nil {
		unregistered = append(unregistered, "DeleteHomeHandler")
	}
	if o.DeleteOfficeHandler == nil {
		unregistered = append(unregistered, "DeleteOfficeHandler")
	}
	if o.GetEmployeeHandler == nil {
		unregistered = append(unregistered, "GetEmployeeHandler")
	}
	if o.GetHomeHandler == nil {
		unregistered = append(unregistered, "GetHomeHandler")
	}
	if o.GetOfficeHandler == nil {
		unregistered = append(unregistered, "GetOfficeHandler")
	}
	if o.ListEmployeesHandler == nil {
		unregistered = append(unregistered, "ListEmployeesHandler")
	}
	if o.ListHomesHandler == nil {
		unregistered = append(unregistered, "ListHomesHandler")
	}
	if o.ListOfficesHandler == nil {
		unregistered = append(unregistered, "ListOfficesHandler")
	}
	if o.UpdateEmployeeHandler == nil {
		unregistered = append(unregistered, "UpdateEmployeeHandler")
	}
	if o.UpdateHomeHandler == nil {
		unregistered = append(unregistered, "UpdateHomeHandler")
	}
	if o.UpdateOfficeHandler == nil {
		unregistered = append(unregistered, "UpdateOfficeHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WancloudsEmpolyeeHubAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WancloudsEmpolyeeHubAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *WancloudsEmpolyeeHubAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WancloudsEmpolyeeHubAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *WancloudsEmpolyeeHubAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WancloudsEmpolyeeHubAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the wanclouds empolyee hub API
func (o *WancloudsEmpolyeeHubAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WancloudsEmpolyeeHubAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/employee"] = NewAddEmployee(o.context, o.AddEmployeeHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/home"] = NewAddHome(o.context, o.AddHomeHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/office"] = NewAddOffice(o.context, o.AddOfficeHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/employee/{ID}"] = NewDeleteEmployee(o.context, o.DeleteEmployeeHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/home/{ID}"] = NewDeleteHome(o.context, o.DeleteHomeHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/office/{ID}"] = NewDeleteOffice(o.context, o.DeleteOfficeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/employee/{ID}"] = NewGetEmployee(o.context, o.GetEmployeeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/home/{ID}"] = NewGetHome(o.context, o.GetHomeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/office/{ID}"] = NewGetOffice(o.context, o.GetOfficeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/employee"] = NewListEmployees(o.context, o.ListEmployeesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/home"] = NewListHomes(o.context, o.ListHomesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/office"] = NewListOffices(o.context, o.ListOfficesHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/employee/{ID}"] = NewUpdateEmployee(o.context, o.UpdateEmployeeHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/home/{ID}"] = NewUpdateHome(o.context, o.UpdateHomeHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/office/{ID}"] = NewUpdateOffice(o.context, o.UpdateOfficeHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WancloudsEmpolyeeHubAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WancloudsEmpolyeeHubAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WancloudsEmpolyeeHubAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WancloudsEmpolyeeHubAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WancloudsEmpolyeeHubAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
