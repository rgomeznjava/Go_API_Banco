# Go_API_Banco Lambda 

Aplicación Ejemplo GOLANG API REST  Banco Lambda

Proposito: Aprendizaje desarrollo en GO, AWS Lambdas..etc

-  Despliegue en AWS - Lambdas
-  Router Mux + AWS
-  Lectura Config app.properties
-  Validaciones de Structs JSON
-  API REST /Cuentas GET/POST
-  API REST /Movimientos GET/POST
-  Capa modelo Fake (in memory)


URLS PRUEBAS REST API EN AWS:

- CONSULTA LISTA DE CUENTAS

https://d3g1k8jx21.execute-api.us-east-2.amazonaws.com/v1/cuentas

- CONSULTA MOVIMIENTOS DE UNA CUENTA por numCuenta: 

{numCuenta}: 11110100889999999991,..2,..4

https://d3g1k8jx21.execute-api.us-east-2.amazonaws.com/v1/movimientos/11110100889999999991


- CONSULTA DE 1 MOVIMIENTO POR idMovimiento: 

{idMovimiento}: 73bafcfb-a16d-4693-8579-0380174001b0

https://d3g1k8jx21.execute-api.us-east-2.amazonaws.com/v1/movimiento/73bafcfb-a16d-4693-8579-0380174001b0


- AÑADIR MOVIMIENTOS A UNA CUENTA - POST

https://d3g1k8jx21.execute-api.us-east-2.amazonaws.com/v1/movimientos

{
        "numCuenta": "11110100889999999992",
        "fecha": "11-11-2021",
        "categoria": "Compras",
        "descripcion": "Frigorifico Bosh",
        "importe": 600,
        "saldo": 1300.44
    }

