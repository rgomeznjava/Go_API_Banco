package gestion_movimientos

import (
	"fmt"

	"github.com/google/uuid"
)

type Movimiento struct {
	Id          string  `json:"id"`
	NumCuenta   string  `json:"numCuenta"`
	Fecha       string  `json:"fecha"`
	Categoria   string  `json:"categoria"`
	Descripcion string  `json:"descripcion"`
	Importe     float32 `json:"importe"`
	Saldo       float32 `json:"saldo"`
}

var movimientos []Movimiento = []Movimiento{
	NewMovimiento("11110100889999999991", "07-11-2021", "Deportes", "Pago Gimnasio Judo Ippon", 45.00, 1900.44),
	NewMovimiento("11110100889999999991", "04-11-2021", "Telefono", "Recibo Yoigor", 22.00, 1945.44),
	NewMovimiento("11110100889999999991", "01-11-2021", "Subcripciones", "Colección Star Wars ", 32.56, 1967.44),
	NewMovimiento("11110100889999999998", "07-11-2021", "Deportes", "Pago Gimnasio Judo Ippon", 45.00, 1900.44),
	NewMovimiento("11110100889999999998", "04-11-2021", "Telefono", "Recibo Yoigor", 22.00, 1945.44),
	NewMovimiento("11110100889999999998", "01-11-2021", "Subcripciones", "Colección Star Wars ", 32.56, 1967.44),
	NewMovimiento("11110100889999999993", "07-11-2021", "Deportes", "Pago Gimnasio Judo Ippon", 45.00, 1900.44),
	NewMovimiento("11110100889999999993", "04-11-2021", "Telefono", "Recibo Yoigor", 22.00, 1945.44),
	NewMovimiento("11110100889999999993", "01-11-2021", "Subcripciones", "Colección Star Wars ", 32.56, 1967.44),
	NewMovimiento("11110100889999999995", "07-11-2021", "Deportes", "Pago Gimnasio Judo Ippon", 45.00, 1900.44),
	NewMovimiento("11110100889999999995", "04-11-2021", "Telefono", "Recibo Yoigor", 22.00, 1945.44),
	NewMovimiento("11110100889999999995", "01-11-2021", "Subcripciones", "Colección Star Wars ", 32.56, 1967.44),
}

//NEW MOVIMIENTO (SIMULACION)
func NewMovimiento(numCuenta string, fecha string, categoria string, descripcion string, importe float32, saldo float32) Movimiento {
	id := uuid.New().String()
	return Movimiento{id, numCuenta, fecha, categoria, descripcion, importe, saldo}
}

//CREATE NEW MOVIMIENTO
func CreateNewMovimiento(mov *Movimiento) (*Movimiento, error) {

	if mov.NumCuenta == "" {
		return nil, fmt.Errorf("DEBE INDICAR NUMERO DE CUENTA DEL MOVIMIENTO")
	}
	id := uuid.New().String()
	mov.Id = id
	return mov, nil
}

// ADD MOVIMIENTO
func AddMovimiento(mov Movimiento) {
	movimientos = append(movimientos, mov)
	fmt.Println(movimientos)
}

// DELETE MOVIMIENTO
func DeleteMovimiento(mov *Movimiento) {

	var movimientosAux []Movimiento

	fmt.Printf("DeleteMovientos-movimientos.len: %d", len(movimientos))
	fmt.Println()

	for _, m := range movimientos {

		if m.Id == mov.Id {
			fmt.Println("--> ENCONTRADO MOVIENTO DELETE: " + m.Id)
		} else {

			movimientosAux = append(movimientosAux, m)
		}

	}
	movimientos = movimientosAux

	fmt.Printf("DeleteMovientos-movimientos.len: %d", len(movimientos))
}

// GET MOVIMIENTOS BY NUM.CUENTA
func GetMovimientosByNumCuenta(numCuenta string) []Movimiento {

	var movimientosCuenta []Movimiento = []Movimiento{}

	for _, mov := range movimientos {

		if numCuenta == mov.NumCuenta {
			movimientosCuenta = append(movimientosCuenta, mov)
		}
	}

	fmt.Println("movimientosCuenta:")
	for _, m := range movimientosCuenta {

		fmt.Println(m)
	}

	return movimientosCuenta
}

// GET MOVIMIENTO BY ID
func GetMovimientoById(id string) (*Movimiento, error) {
	for _, m := range movimientos {
		if m.Id == id {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("el movimiento con id %s no existe", id)
}

