# Importo la libreria socket para manejar conexiones
import socket as skt

# Se declara dirección del servidor y puerto
serverAddr = 'localhost'
serverPort = 63420

# Se crea un socket para hacer el manejo de la conexión
clientSocket = skt.socket(skt.AF_INET, skt.SOCK_DGRAM)

# Solicito mensajes y luego ejecuto las funciones para conectar al servidor y realizar el intercambio de mensajes.

toSend = input("Qué deseas hacer? \nMarca 1 si quieres agregar tu información, Marca 2 en caso de querer la dirección IP de un nombre de dominio, \nEn caso de que quieras salir escribe STOP: \n")



if toSend == "STOP":
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))
	print("Programa Finalizado")
if toSend == "1":

	nombre_dominio = input("Ingrese el Nombre de dominio :\n>")
	Ip_direction = input("Ingrese Dirección IP :\n>")
	ttl = input("Ingrese TTL :\n>")
	type = input("Ingrese el Tipo :\n>")

	print()

	clientSocket.sendto(nombre_dominio.encode(), (serverAddr,serverPort))
	clientSocket.sendto(Ip_direction.encode(), (serverAddr,serverPort))
	clientSocket.sendto(ttl.encode(), (serverAddr,serverPort))
	clientSocket.sendto(type.encode(), (serverAddr,serverPort))

	msg, addr = clientSocket.recvfrom(1024)
	print(msg.decode()) 

if toSend == "2":
	print()

	nombre_dominio = input("Ingrese el Nombre de dominio :\n>")

  #Falta poner el sistema de imprimir en pantalla la respuesta del servidor dado el nombre de dominio (debe devolver la direccion ip)

	msg, addr = clientSocket.recvfrom(1024)
	print(msg.decode()) 
