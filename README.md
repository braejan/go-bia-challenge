# Bia Energy Golang Tech Challenge
Una implementación de un reto técnico usando Golang y microservicios.

## Objetivo

Se requiere crear un microservicio que esté conectado a una base de datos SQL que tiene una tabla donde se almacenan los consumos de energía de los medidores de nuestros clientes. Este microservicio está integrado por API al microservicio de direcciones el cual almacena las direcciones donde están instalados los medidores.

### Consideraciones y restricciones:
- Los identificadores de los medidores se deben enviar cómo query string parameter separados por coma (en caso de enviar más de uno) y sin espacios entre sí. Ejemplo: `1,2,3`
- El único formato válido para los campos de tipo fecha es `YYYY-MM-DD`
- Los únicos valores permitidos para el tipo de consulta son `daily`, `weekly` y `monthly`


## Requerimientos

- Golang 1.19 o superior (https://golang.org/dl/)
- Docker (https://docs.docker.com/get-docker/)
- Docker Compose (https://docs.docker.com/compose/install/)

## Instalación

1. Clona este repositorio en tu máquina local.
2. Ubica tu terminal en el directorio del repositorio `go-bia-challenge`.
### Construir el proyecto localmente usando docker compose:

```console
docker-compose up
```
#### Puede tomar un tiempo mientras descarga las dependencias y construye el proyecto.
En este paso, se está realizando tambien el cargue del archivo [up.sql](internal/infraestructure/database/postgres/up.sql) que realiza el cargue de los datos de manera local.
#### No te gustan las cosas por defecto?
Puedes entrar a este link oficial de [postgres en docker](https://hub.docker.com/_/postgres) para construir tu base de datos personalizada.

## Uso

Despues de haber ejecutado el proyecto localmente con `docker-compose` puedes validar por consola que estén los dos servicios arriba con el comando:
```bash
docker-compose ps
```
Si los servicios se encuentran listos para recibir peticiones, deberá arrojar el resultado de los dos:
```bash
            Name                          Command              State                 Ports              
--------------------------------------------------------------------------------------------------------
go-bia-challenge_database_1    docker-entrypoint.sh postgres   Up      0.0.0.0:5432->5432/tcp,:::5432-  
                                                                       >5432/tcp                        
go-bia-challenge_file-rest_1   ./consumption-rest              Up      0.0.0.0:8080->8080/tcp,:::8080-  
                                                                       >8080/tcp                        
```

Para ejecutar el servicio local, se puede hacer uso del software [curl](https://curl.se/).

## Url
    http://localhost:8080/consumption\?meter_ids\={{ids}}\&kind_period\={{kindPeriod}}\&start_date\={{startDate}}\&end_date\={{endDate}}
### Consideraciones
- Todas las variables son requeridas en la URL.
- `{{ids}}`: Los ID a consultar separados por coma y sin espacios. Ej: `1,2,3`
- `{{kindPeriod}}`: Tipo de consulta a realizar `daily`, `weekly` o `monthly`
- `{{startDate}}`: Rango de fecha inicial de la consulta
- `{{endDate}}`: Rango de fecha final de la consulta

# Ejemplo
```bash
curl http://localhost:8080/consumption\?meter_ids\=1,2\&kind_period\=monthly\&start_date\=2023-06-01\&end_date\=2023-07-31
```
Con los datos locales arrojará el siguiente resultado:

```json
{
  "period": [
    "JUN 2023",
    "JUL 2023"
  ],
  "data_graph": [
    {
      "meter_id": 1,
      "address": "Dirección Mock",
      "active_energy": [
        28853895.42664002,
        22129138.077179987
      ],
      "reactive_energy": [
        4486645.46164,
        1643312.1089800023
      ],
      "capacitive_reactive": [
        7372.619,
        2341.153
      ],
      "solar": [
        5943.037877272408,
        3114.0005055555116
      ]
    },
    {
      "meter_id": 2,
      "address": "Dirección Mock",
      "active_energy": [
        39174684.245770015,
        22931124.777489986
      ],
      "reactive_energy": [
        43211872.66273,
        3588724.2794
      ],
      "capacitive_reactive": [
        0,
        0
      ],
      "solar": [
        468.73814626165364,
        223.37283915773872
      ]
    }
  ]
}
```


## Pruebas

Para ejecutar las pruebas unitarias, debes ejecutar los siguientes comando:
```bash
go test -v ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```
Esta ejecución generará en consola el resultado de la ejecución del set de pruebas de todos los archivos *_test.go. Además, generará un archivo coverage.html que puedes abrir en tu navegador para ver el porcentaje de cobertura de las pruebas.
## Deuda técnica.
El proyecto debe implementar un set de pruebas de integración.

## Licencia

Este proyecto está bajo la licencia [MIT License](LICENSE).