package main

import (
	"fmt"
	"net"
)

func enviarDatos(conexion net.Conn, mensaje string) (string, error) {
	_, err := fmt.Fprint(conexion, mensaje)
	if err != nil {
		return "", err
	}
	datosRespuesta := make([]byte, 2048)
	cantidadBytes, err := conexion.Read(datosRespuesta)
	if err != nil {
		return "", err
	}
	return string(datosRespuesta[:cantidadBytes]), nil
}

func leerMensaje() string {
	var palabra1, palabra2, palabra3, palabra4, palabra5, palabra6 string
	cantidad, _ := fmt.Scanln(&palabra1, &palabra2, &palabra3, &palabra4, &palabra5, &palabra6)
	listaPalabras := []string{palabra1, palabra2, palabra3, palabra4, palabra5, palabra6}
	textoFinal := ""
	for indice := 0; indice < cantidad; indice++ {
		if indice > 0 {
			textoFinal = textoFinal + " "
		}
		textoFinal = textoFinal + listaPalabras[indice]
	}
	return textoFinal
}

func main() {
	var ip string
	var puerto string

	fmt.Println("CLIENTE DE CHAT")
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
	conexion, err := net.Dial("tcp", direccion)
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer conexion.Close()

	var cuentaActual string

	for {
		if cuentaActual == "" {
			var cuentaIngresada string
			fmt.Print("\nUsuario (josue, david, adrian) o salir: ")
			fmt.Scanln(&cuentaIngresada)

			if cuentaIngresada == "salir" {
				return
			}

			respuesta, err := enviarDatos(conexion, "LOGIN|"+cuentaIngresada)
			if err != nil {
				fmt.Println("Error de red:", err)
				return
			}

			fmt.Println(respuesta)
			if len(respuesta) >= 3 && respuesta[:3] == "OK:" {
				cuentaActual = cuentaIngresada
			}
		} else {
			var opcion string
			fmt.Println("\nMENU (" + cuentaActual + ")")
			fmt.Println("1. Enviar mensaje")
			fmt.Println("2. Ver historial")
			fmt.Println("3. Cerrar sesion")
			fmt.Print("Opcion: ")
			fmt.Scanln(&opcion)

			if opcion == "1" {
				fmt.Print("Mensaje: ")
				texto := leerMensaje()

				respuesta, err := enviarDatos(conexion, "ENVIAR|"+cuentaActual+"|"+texto)
				if err != nil {
					fmt.Println("Error de red:", err)
					return
				}
				fmt.Println(respuesta)
			} else if opcion == "2" {
				respuesta, err := enviarDatos(conexion, "HISTORIAL|")
				if err != nil {
					fmt.Println("Error de red:", err)
					return
				}
				fmt.Println(respuesta)
			} else if opcion == "3" {
				respuesta, err := enviarDatos(conexion, "SALIR|"+cuentaActual)
				if err != nil {
					fmt.Println("Error de red:", err)
					return
				}
				fmt.Println(respuesta)
				cuentaActual = ""
			} else {
				fmt.Println("ERROR: Opcion no valida")
			}
		}
	}
}