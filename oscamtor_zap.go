// Declarar un paquete principal (un paquete es una forma de agrupar funciones y está compuesto por todos los archivos en el mismo directorio).
package main

// UNO
// **************
// Importar el popular paquete fmt, que contiene funciones para formatear texto, incluida la impresión en la consola. Este paquete es uno de los paquetes de biblioteca estándar que obtuvo cuando instaló Go.
import "fmt"

func Hello() string {
	return "Hello, world"
}

// Implementar una función main para imprimir un mensaje en la consola. Una función man se ejecuta de forma predeterminada cuando ejecuta el paquete principal
func main() {
	fmt.Println(Hello())
}

// DOS (not tested)
// *****
// import (
// "fmt"
// "rsc.io/quote"
// )

// func main() {
    // Example_presets()
// }



// TRES OK (ok tested)
//*************

// import (
// "fmt"
// "rsc.io/quote"
// )

// func main() {
    // fmt.Println(quote.Go())
// }

