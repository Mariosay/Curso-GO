#Curso GO
https://go.dev/dl/
##Son ejemplos muy basicospero necesarios para entender este curso
##Es necesario correr estos comandos desde la consola


| Código  | Descripcion |
| ------------- | ------------- |
|1 ```go version ``` | muestra la version de GO instalada |
|2 ```go build nombrArchivo.go``` | genera el ejecutable de la app |
|3```go run nombrArchivo.go ```| corre el codigo sin generar el ejecutable|
|4```go mod init carpetadelproyecto  ```| crea un archi o go.mod el cual nos sirve para guardar drivers en este caso para hacer la conexion a mysql|
|5```go get -u github.com/go-sql-driver/mysql  ```| descarga el driver genera un archivo go.sum|
|6```go env ```| muestra las variables de entorno|
|7```go get -u github.com/gorilla/mux ``` |esto se tiene que hacer cuando se pasa de linux a windows junto con el paso 4, primero 4 y despues este paso, segun sea la libreria que se necesita|
###
#nota 
En GO solo hay el ciclo for


