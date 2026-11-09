package main

import (
	"chat/internal/handlers"
	"chat/internal/services"
	"chat/internal/storage"
	"fmt"
	"net"
)

func main() {
	var ip string
	var puerto string

	fmt.Println("SERVIDOR DE CHAT")
	fmt.Print("Ingrese IP (Enter para 127.0.0.1): ")
	fmt.Scanln(&ip)
	if ip == "" {
		ip = "127.0.0.1"
	}

	fmt.Print("Ingrese Puerto (Enter para 9000): ")
	fmt.Scanln(&puerto)
	if puerto == "" {
		puerto = "9000"
	}

	direccion := ip + ":" + puerto
	escucha, err := net.Listen("tcp", direccion)
	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
		return
	}
	defer escucha.Close()

	fmt.Println("Servidor activo en", direccion)

	almacenamiento := storage.NuevoAlmacenamiento()
	servicio := services.NuevoServicioChat(almacenamiento)
	manejador := handlers.NuevoManejadorRed(servicio)

	for {
		conexion, err := escucha.Accept()
		if err != nil {
			fmt.Println("Error de conexion:", err)
			continue
		}
		fmt.Println("Cliente conectado")
		go manejador.AtenderCliente(conexion)
	}
}