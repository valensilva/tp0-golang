package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()
	//log.Println("soy un log")

	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente
	log.Println(globals.ClientConfig.Mensaje)
	// loggeamos el valor de la config

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	//utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)
	// leer de la consola el mensaje

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto)
}
