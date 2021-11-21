package gestion_cuentas

import (
	"fmt"

	"github.com/google/uuid"
)

type ListaCuentas struct {
	Cuentas []Cuenta `json:"cuentas"`
}

type Cuenta struct {
	Id              string  `json:"id"`
	Titular         string  `json:"titular" validate:"min=3,max=50"`
	Tipo            string  `json:"tipo" validate:"max=25"`
	FechaApertura   string  `json:"fechaApertura" validate:"required"`
	Numero          string  `json:"numero"`
	IBAN            string  `json:"IBAN"`
	CodigoBIC       string  `json:"codigoBIC"`
	SaldoDisponible float32 `json:"saldoDisponible"`
	SaldoContable   float32 `json:"saldoContable" validate:"required"`
}

//--DATOS SIMULACION ARRAY DE CUENTAS
var cuentas []Cuenta = []Cuenta{

	CreateCuenta("Peperrr Apellido1 Apellido2", "Cuenta NOMINA", "07-11-2021", "11110100889999999999", "ES99 1111 0100 88 1111111111", "YYYYYYCCCCC", 158000.00, 158000.00),
	CreateCuenta("Miriammm  Apellido1 Apellido2", "Cuenta AHORRO", "04-11-2021", "22220100889999999999", "ES99 2222 0100 88 2222222222", "YYYYYYCCCCC", 158000.00, 158000.00),
	CreateCuenta("Annnna Apellido1 Apellido2", "Cuenta NOMINA", "01-11-2021", "7777333889999999999", "ES99 3333 0100 88 3333333333", "YYYYYYCCCCC", 158000.00, 158000.00),
}

//-- FUNCIONES LOG.NEGOCIO
func CreateCuenta(titular string, tipo string, fechaApertura string, numero string, IBAN string, codigoBIC string, saldoDisponible float32, saldoContable float32) Cuenta {
	id := uuid.New().String()
	return Cuenta{id, titular, tipo, fechaApertura, numero, IBAN, codigoBIC, saldoDisponible, saldoContable}
}

func CreateNewCuenta(c *Cuenta) (*Cuenta, error) {

	if c.Titular == "" {
		return nil, fmt.Errorf("DEBE INDICAR NOMBRE Y APELLIDOS DEL TITULAR")
	}
	id := uuid.New().String()
	c.Id = id
	return c, nil
}

func AddCuenta(c Cuenta) {
	cuentas = append(cuentas, c)
	fmt.Println(cuentas)
}

func GetCuentas() []Cuenta {
	return cuentas
}

func GetListaCuentas() ListaCuentas {
	return ListaCuentas{GetCuentas()}
}

func GetCuentaByNumero(numerocuenta string) (*Cuenta, error) {
	for _, c := range cuentas {
		if c.Numero == numerocuenta {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("la cuenta con número %s no existe", numerocuenta)
}
