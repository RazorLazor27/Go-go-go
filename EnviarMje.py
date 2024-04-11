# Importo la libreria socket para manejar conexiones
import socket as skt

# Se declara dirección del servidor y puerto
serverAddr = 'localhost'
serverPort = 63420

# Se crea un socket para hacer el manejo de la conexión
clientSocket = skt.socket(skt.AF_INET, skt.SOCK_DGRAM)

# Solicito mensajes y luego ejecuto las funciones para conectar al servidor y realizar el intercambio de mensajes.



toSend = input("Qué deseas hacer? \nMarca 1 si quieres agregar tu información\nMarca 2 en caso de querer la dirección IP de un nombre de dominio, \nEn caso de que quieras salir escribe STOP: \n")
flag = True
clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))

while flag:
	

	

	if toSend == "STOP":
		clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))
		print("Programa Finalizado")
		flag = True
	if toSend == "1":



		print()

		nombre_dominio = input("Ingrese el Nombre de dominio :\n>")
		Ip_direction = input("Ingrese Dirección IP :\n>")
		ttl = input("Ingrese TTL :\n>")
		tipo = input("Ingrese el Tipo :\n>")

		# Concatenar las cadenas con punto y coma
		enviar = nombre_dominio + ";" + Ip_direction + ";" + ttl + ";" + tipo

		# Enviar la cadena concatenada al servidor
		clientSocket.sendto(enviar.encode(), (serverAddr,serverPort))
	if toSend == "2":
		print()

		nombre_dominio = input("Ingrese el Nombre de dominio :\n>")
		clientSocket.sendto(nombre_dominio.encode(), (serverAddr,serverPort))

	#Falta poner el sistema de imprimir en pantalla la respuesta del servidor dado el nombre de dominio (debe devolver la direccion ip)


		msg, addr = clientSocket.recvfrom(1024)
		print("aqui respuesta")
		print(msg.decode()) 

	toSend = input("\nQué deseas hacer? \nMarca 1 si quieres agregar tu información\nMarca 2 en caso de querer la dirección IP de un nombre de dominio, \nEn caso de que quieras salir escribe STOP: \n")
	clientSocket.sendto(toSend.encode(), (serverAddr,serverPort))
