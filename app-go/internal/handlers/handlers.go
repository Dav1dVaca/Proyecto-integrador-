package handlers

import (
	"chat/internal/services"
	"fmt"
	"net"
)

type ManejadorRed struct {
	servicio *services.ServicioChat
}

func NuevoManejadorRed(servicio *services.ServicioChat) *ManejadorRed {
	return &ManejadorRed{
		servicio: servicio,
	}
}

func separarSolicitud(solicitud string) (string, string, string) {
	parte1 := ""
	parte2 := ""
	parte3 := ""
	etapa := 1
	for indice := 0; indice < len(solicitud); indice++ {
		if solicitud[indice] == '|' && etapa < 3 {
			etapa++
		} else {
			if etapa == 1 {
				parte1 = parte1 + string(solicitud[indice])
			} else if etapa == 2 {
				parte2 = parte2 + string(solicitud[indice])
			} else {
				parte3 = parte3 + string(solicitud[indice])
			}
		}
	}
	return parte1, parte2, parte3
}

func (manejador *ManejadorRed) AtenderCliente(conexion net.Conn) {
	defer conexion.Close()
	datosRecibidos := make([]byte, 1024)

	for {
		cantidadBytes, err := conexion.Read(datosRecibidos)
		if err != nil {
			return
		}

		solicitud := string(datosRecibidos[:cantidadBytes])
		comando, dato1, dato2 := separarSolicitud(solicitud)

		var respuesta string

		if comando == "LOGIN" {
			usuarioConectado, err := manejador.servicio.IniciarSesion(dato1)
			if err != nil {
				respuesta = "ERROR: " + err.Error()
			} else {
				respuesta = "OK: Bienvenido " + usuarioConectado.ObtenerNombre() + " (" + usuarioConectado.ObtenerSede() + ")"
			}
		} else if comando == "ENVIAR" {
			mensajeEnviado, err := manejador.servicio.EnviarMensaje(dato1, dato2)
			if err != nil {
				respuesta = "ERROR: " + err.Error()
			} else {
				respuesta = "OK: " + mensajeEnviado.Mostrar()
			}
		} else if comando == "HISTORIAL" {
			respuesta = manejador.servicio.ConsultarHistorial()
		} else if comando == "SALIR" {
			manejador.servicio.CerrarSesion(dato1)
			respuesta = "OK: Sesion cerrada"
		} else {
			respuesta = "ERROR: Opcion no valida"
		}

		fmt.Fprint(conexion, respuesta)
	}
}