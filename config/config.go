package config

import (
	"fmt"
	"log"

	"github.com/magiconair/properties"
)

var Cfg Config

type Config struct {
	SERVIDOR_1                   string `properties:"SERVIDOR_1"`
	SERVIDOR_2                   string `properties:"SERVIDOR_2"`
	URL_REST_ALTA_CUENTA         string `properties:"URL_REST_ALTA_CUENTA"`
	URL_REST_LISTA_CUENTAS       string `properties:"URL_REST_LISTA_CUENTAS"`
	URL_REST_CONSULTA_CUENTA     string `properties:"URL_REST_CONSULTA_CUENTA"`
	URL_REST_LISTA_MOVIMIENTOS   string `properties:"URL_REST_LISTA_MOVIMIENTOS"`
	URL_REST_ALTA_MOVIMIENTO     string `properties:"URL_REST_ALTA_MOVIMIENTO"`
	URL_REST_CONSULTA_MOVIMIENTO string `properties:"URL_REST_CONSULTA_MOVIMIENTO"`
	URL_REST_BORRAR_MOVIMIENTO   string `properties:"URL_REST_BORRAR_MOVIMIENTO"`

	//Port                 int    `properties:"port,default=9000"`
	//Accept  []string      `properties:"accept,default=image/png;image;gif"`
	//Timeout time.Duration `properties:"timeout,default=5s"`
}

func LeerConfiguraciones() {

	fmt.Println("--> Leyendo configuraciones......")

	// init from a file
	//p := properties.MustLoadFile("${HOME}/Documents/_CURSOS_2020/curso-golang-at-master/api-gestion-banco/app.properties", properties.UTF8)
	p := properties.MustLoadFile("config/app.properties", properties.UTF8)

	// or through Decode∫
	if err := p.Decode(&Cfg); err != nil {
		log.Fatal("HA FALLADO AL DECODIFICAR PROPERTIES.....")
		log.Fatal(err)
	}

}

func GetConfiguracion() Config {
	return Cfg
}
