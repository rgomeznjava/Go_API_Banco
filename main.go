package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"

	//gestcuentas "github.com/raulgomezn/gestion_cuentas"

	"github.com/raulgomezn/config"
	handlers "github.com/raulgomezn/handlers"
)

var gorillaLambda *gorillamux.GorillaMuxAdapter

func init() {

	config.LeerConfiguraciones()

	r := mux.NewRouter()
	r.HandleFunc(config.Cfg.URL_REST_LISTA_CUENTAS, handlers.HandlerGetListaCuentas).Methods(http.MethodGet)
	r.HandleFunc(config.Cfg.URL_REST_CONSULTA_CUENTA, handlers.HandlerGetCuentaByNumero).Methods(http.MethodGet)
	r.HandleFunc(config.Cfg.URL_REST_ALTA_CUENTA, handlers.HandlerPostCuentas).Methods(http.MethodPost)
	r.HandleFunc(config.Cfg.URL_REST_LISTA_MOVIMIENTOS, handlers.HandlerGetMovimientosByNumCuenta).Methods(http.MethodGet)
	r.HandleFunc(config.Cfg.URL_REST_CONSULTA_MOVIMIENTO, handlers.HandlerGetMovimientoById).Methods(http.MethodGet)
	r.HandleFunc(config.Cfg.URL_REST_ALTA_MOVIMIENTO, handlers.HandlerPostMovimientos).Methods(http.MethodPost)
	r.HandleFunc(config.Cfg.URL_REST_BORRAR_MOVIMIENTO, handlers.HandlerDeleteMovimientos).Methods(http.MethodDelete)
	gorillaLambda = gorillamux.New(r)
	//gorillaLambda = handlers.Init()
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return gorillaLambda.ProxyWithContext(ctx, req)
}
