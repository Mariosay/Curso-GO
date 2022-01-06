#Curso GO
https://go.dev/dl/
##Son ejemplos muy basicospero necesarios para entender este curso
##Es necesario correr estos comandos desde la consola


| Código  | Descripcion |
| ------------- | ------------- |
| ```go version ``` | muestra la version de GO instalada |
| ```go build nombrArchivo.go``` | genera el ejecutable de la app |
|```go run nombrArchivo.go ```| corre el codigo sin generar el ejecutable|


#nota 
En GO solo hay el ciclo for

###
go mod init carpetadelproyecto -> crea un archivo .mod el cual nos sirve para guardar drivers en este caso para hacer la conexion a mysql
go get -u github.com/go-sql-driver/mysql
