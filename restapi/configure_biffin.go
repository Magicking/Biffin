package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"
	swag "github.com/go-openapi/swag"

	"github.com/Magicking/Biffin/restapi/operations"
	"github.com/Magicking/Biffin/internal"
	"github.com/rs/cors"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../biffin-api.yaml

func configureFlags(api *operations.BiffinAPI) {
	ethOpts := swag.CommandLineOptionsGroup{
		LongDescription:  "",
		ShortDescription: "Ethereum options",
		Options:          &ethopts,
	}
	serviceOpts := swag.CommandLineOptionsGroup{
		LongDescription:  "",
		ShortDescription: "Service options",
		Options:          &serviceopts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ethOpts, serviceOpts}
}

var ethopts struct {
	RawURL string `long:"ws-uri" env:"WS_URI" description:"Ethereum WS URI (e.g: ws://HOST:8546)"`
	PrivateKey string `long:"pkey" env:"PRIVATE_KEY" description:"hex encoded private key"`
}

var serviceopts struct {
	DbDSN string `long:"db-dsn" env:"DB_DSN" description:"Database DSN (e.g: /tmp/test.sqlite)"`
}

func configureAPI(api *operations.BiffinAPI) http.Handler {
	// configure the api here
	ctx, err := internal.InitContext(ethopts.RawURL, ethopts.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to initialize connection to ethereum node: %v", err)
	}
	ctx.DB, err = internal.InitDatabase(serviceopts.DbDSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetMapHandler = operations.GetMapHandlerFunc(func(params operations.GetMapParams) middleware.Responder {
		return internal.GetMap(ctx, params)
	})
	api.PostMapHandler = operations.PostMapHandlerFunc(func(params operations.PostMapParams) middleware.Responder {
		return internal.PostMap(ctx, params)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}
