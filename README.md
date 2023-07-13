# go-bia-challenge
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
    En este paso, se está realizando tambien el cargue del archivo [up.sql](internal/infraestructure/database/postgres/up.sql).
#### No te gustan las cosas por defecto?
Puedes entrar a este link oficial de [postgres en docker](https://hub.docker.com/_/postgres) para construir tu base de datos personalizada.

## Uso