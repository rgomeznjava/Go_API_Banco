package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"

	//--"github.com/raulgomezn/gestion_cuentas" //sin alias

	config "github.com/raulgomezn/config"
	gestcuentas "github.com/raulgomezn/gestion_cuentas"
	gestmov "github.com/raulgomezn/gestion_movimientos"
	validaciones "github.com/raulgomezn/validaciones"
)

var cfg = &config.Cfg

var gorillaLambda *gorillamux.GorillaMuxAdapter

func HandlerPostCuentas(w http.ResponseWriter, r *http.Request) {

	var datos *gestcuentas.Cuenta
	json.NewDecoder(r.Body).Decode(&datos)

	err := validaciones.ValidateStruct(*datos)

	if err != nil {
		w.WriteHeader(http.StatusConflict)

		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	cuentaOut, err := gestcuentas.CreateNewCuenta(datos)

	gestcuentas.AddCuenta(*cuentaOut)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusConflict)

		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cuentaOut)
}

func HandlerGetListaCuentas(w http.ResponseWriter, r *http.Request) {

	listaCuentas := gestcuentas.GetListaCuentas()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listaCuentas)
}

func HandlerGetCuentaByNumero(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	numCuenta := vars["numCuenta"]

	cuenta, err := gestcuentas.GetCuentaByNumero(numCuenta)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cuenta)
}

//-------------------------------
//--- HANDLERS MOVIMIENTOS ------

func HandlerPostMovimientos(w http.ResponseWriter, r *http.Request) {

	var datos *gestmov.Movimiento
	json.NewDecoder(r.Body).Decode(&datos)

	err := validaciones.ValidateStruct(*datos)

	if err != nil {
		w.WriteHeader(http.StatusConflict)

		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	movimientoOut, err := gestmov.CreateNewMovimiento(datos)

	gestmov.AddMovimiento(*movimientoOut)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusConflict)

		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movimientoOut)
}

func HandlerDeleteMovimientos(w http.ResponseWriter, r *http.Request) {

	var datos *gestmov.Movimiento
	json.NewDecoder(r.Body).Decode(&datos)

	err := validaciones.ValidateStruct(*datos)

	if err != nil {
		w.WriteHeader(http.StatusConflict)

		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	//--no--movimientoOut, err := gestmov.CreateNewMovimiento(datos)

	gestmov.DeleteMovimiento(datos)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusConflict)

		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	//--json.NewEncoder(w).Encode(movimientoOut)
}

func HandlerGetMovimientosByNumCuenta(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	numCuenta := vars["numCuenta"]

	//movimientos de la cuenta
	movimientos := gestmov.GetMovimientosByNumCuenta(numCuenta)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movimientos)
}

func HandlerGetMovimientoById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	//movimiento
	movimiento, err := gestmov.GetMovimientoById(id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		msg := map[string]string{
			"msg": err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movimiento)
}

func Init() *gorillamux.GorillaMuxAdapter {

	fmt.Println("--> Arrancando Servidor...")

	r := mux.NewRouter()
	r.HandleFunc(cfg.URL_REST_LISTA_CUENTAS, HandlerGetListaCuentas).Methods(http.MethodGet)
	r.HandleFunc(cfg.URL_REST_CONSULTA_CUENTA, HandlerGetCuentaByNumero).Methods(http.MethodGet)
	r.HandleFunc(cfg.URL_REST_ALTA_CUENTA, HandlerPostCuentas).Methods(http.MethodPost)
	r.HandleFunc(cfg.URL_REST_LISTA_MOVIMIENTOS, HandlerGetMovimientosByNumCuenta).Methods(http.MethodGet)
	r.HandleFunc(cfg.URL_REST_CONSULTA_MOVIMIENTO, HandlerGetMovimientoById).Methods(http.MethodGet)
	r.HandleFunc(cfg.URL_REST_ALTA_MOVIMIENTO, HandlerPostMovimientos).Methods(http.MethodPost)
	r.HandleFunc(cfg.URL_REST_BORRAR_MOVIMIENTO, HandlerDeleteMovimientos).Methods(http.MethodDelete)
	gorillaLambda = gorillamux.New(r)
	return gorillaLambda
	//http.ListenAndServe(cfg.SERVIDOR_1, r)

}
