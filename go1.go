package main

import (
	"fmt"
)

func main() {

	var nombre, nProducto string = "", "Pringles"
	var pPringles, cantidad, pAux, cedula int64 = 0, 0, 0, 0
	var iva, pInicial, pFinal float64 = 16.00, 0, 0

	fmt.Print("Introduce el precio del producto: " + "'" + nProducto + "': ")
	fmt.Scan(&pPringles)

	fmt.Print("\nIntroduce el nombre del cliente: ")
	fmt.Scan(&nombre)

	fmt.Print("\nIntroduce la cedula del cliente: ")
	fmt.Scan(&cedula)

	fmt.Println("\nBienvenido al sistema de compras")
	fmt.Print("\nIntroduce el producto que desea comprar el cliente: \n")
	fmt.Println("1. Pringles")
	fmt.Scan(&pAux)
	fmt.Print("\nIntroduce la cantidad que desea comprar el cliente: ")
	fmt.Scan(&cantidad)
	fmt.Println("Has seleccionado 'Pringles'. Tienen un costo de: ", pPringles, "Bs" )

	pInicial = float64(cantidad) * float64(pPringles)
	pFinal = float64(pInicial) + (float64(iva) / 100 * float64(pInicial)) 

	fmt.Println("----------Factura----------")
	fmt.Println("Nombre del cliente: " + nombre)
	fmt.Println("Cedula: ", cedula)
	fmt.Println("Producto seleccionado: " + nProducto)
	fmt.Println("Cantidad comprada: ", cantidad)
	fmt.Println("\nPrecio/Und:", pPringles, "Bs")
	fmt.Println("\nPrecio inicial:", pInicial, "Bs")
	fmt.Println("IVA:", iva , "%")
	fmt.Println("Precio final:", pFinal , "Bs")
	fmt.Println("----------Factura----------")
}