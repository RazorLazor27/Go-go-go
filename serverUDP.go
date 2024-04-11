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

	if strings.Contains(msg, "1") {
		// guardar informacion
		fmt.Println("El mensaje recibido fue: ", msg)
		buffer := make([]byte, 1024)
		n, addr, err := conexion.ReadFromUDP(buffer)
		fmt.Print(" --- ", string(buffer[0:n]))
		fmt.Println("El mensaje recibido de: ", addr.String(), ", fue:", string(buffer[:n]))

		aux := string(buffer[:n])

		if err != nil {
			fmt.Println(err)
		}

		subcadena := strings.Split(aux, ";")

		array[contador][0] = subcadena[0]
		array[contador][1] = subcadena[1]
		array[contador][2] = subcadena[2]
		array[contador][3] = subcadena[3]
		contador++

	}
	if strings.Contains(msg, "2") {
		fmt.Println("El mensaje recibido fue: ", msg)
		fmt.Println("El mensaje recibido fue: ", msg)
		buffer := make([]byte, 1024)
		n, addr, err := conexion.ReadFromUDP(buffer)
		fmt.Print(" --- ", string(buffer[0:n]))
		fmt.Println("El mensaje recibido de: ", addr.String(), ", fue:", string(buffer[:n]))

		aux := string(buffer[:n])

		if err != nil {
			fmt.Println(err)
		}

		for i := range array {
			if aux == array[i][0] {
				fmt.Println(" La direccion ip es: ", addr.String())

				auxbuf := []byte(array[i][1])
				_, err := conexion.WriteToUDP(auxbuf, addr)

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("El mensaje fue enviado: ")

			}
		}

	}
	if strings.Contains(msg, "STOP") {
		fmt.Println("La conexion UDP fue cerrada")
		// Si UDPServerStatus es falso, se cierra el servidor
		UDPServerStatus = false
	}
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
