package main

import (
	"fmt"
	"net"
	"strings"
)

const Puerto = ":63420"

// Mostrar si el servido ha validado la conexion UDP
var UDPServerStatus bool = true
var NextMsg = false

var Ip string

var codeExit string = ""

var contador = 0
var posicion = 0

var array [][]string

// Esta funcion mantiene la conexion UDP abierta mientras sea necesario, ayuda a tener un codigo mas conciso en la funcion main
func cicloCliente(conexion *net.UDPConn) {

	buffer := make([]byte, 1024)
	n, addr, err := conexion.ReadFromUDP(buffer)
	fmt.Print(" --- ", string(buffer[0:n]))

	if err != nil {
		fmt.Println("Error en la conexion, cerrando el servidor", err)
		return
	}

	fmt.Println("El mensaje recibido de: ", addr.String(), ", fue:", string(buffer[:n]))
	Ip = addr.String()
	_, err = conexion.WriteToUDP([]byte("Soy la respuesta del servidor"), addr)
	if err != nil {
		fmt.Println("Error al enviar la respuesta al cliente", err)
		return
	}

	// msg := strings.ToUpper(string(buffer[:n]))
	msg := string(buffer[:n])
	// fmt.Println("El mensaje recibido fue: ", msg)

	if strings.Contains(msg, "STOP") {
		fmt.Println("La conexion UDP fue cerrada")
		// Si UDPServerStatus es falso, se cierra el servidor
		UDPServerStatus = false

	} else {
		array[posicion][contador] = msg
		contador++

		if contador == 4 {
			posicion++
		}
		fmt.Println("El mensaje recibido fue: ", msg)

	}
	for _, fila := range array {
		for _, elemento := range fila {
			fmt.Printf("%s ", elemento)
		}
	}
	// else {
	// fmt.Println("Conexion no aceptada, procediendo a cerrar la conexion UDP")
	// UDPServerStatus = false
	//}
}

func main() {

	array = make([][]string, 4)
	for i := range array {
		array[i] = make([]string, 100)
	}

	s, err := net.ResolveUDPAddr("udp4", Puerto)

	if err != nil {
		fmt.Println(err)
		return
	}

	conexion, err := net.ListenUDP("udp4", s)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conexion.Close()

	fmt.Println("El servidor UDP esta en funcionamiento desde ahora en el puerto:", Puerto)

	for UDPServerStatus {
		cicloCliente(conexion)
	}

}
