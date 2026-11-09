package services

import (
	"chat/internal/models"
	"chat/internal/storage"
	"fmt"
)

type ServicioChat struct {
	almacenamiento *storage.Almacenamiento
}

func NuevoServicioChat(almacenamiento *storage.Almacenamiento) *ServicioChat {
	return &ServicioChat{
		almacenamiento: almacenamiento,
	}
}

func (servicio *ServicioChat) IniciarSesion(cuenta string) (*models.Usuario, error) {
	if cuenta == "" {
		return nil, fmt.Errorf("el usuario no puede estar vacio")
	}
	usuarioEncontrado, err := servicio.almacenamiento.BuscarUsuario(cuenta)
	if err != nil {
		return nil, err
	}
	usuarioEncontrado.CambiarActivo(true)
	return usuarioEncontrado, nil
}

func (servicio *ServicioChat) CerrarSesion(cuenta string) error {
	usuarioEncontrado, err := servicio.almacenamiento.BuscarUsuario(cuenta)
	if err != nil {
		return err
	}
	usuarioEncontrado.CambiarActivo(false)
	return nil
}

func (servicio *ServicioChat) EnviarMensaje(cuenta string, texto string) (*models.Mensaje, error) {
	usuarioEncontrado, err := servicio.almacenamiento.BuscarUsuario(cuenta)
	if err != nil {
		return nil, err
	}
	if !usuarioEncontrado.ObtenerActivo() {
		return nil, fmt.Errorf("el usuario no tiene sesion activa")
	}
	if texto == "" {
		return nil, fmt.Errorf("el mensaje no puede estar vacio")
	}
	mensajeGuardado, err := servicio.almacenamiento.GuardarMensaje(usuarioEncontrado.ObtenerNombre(), usuarioEncontrado.ObtenerSede(), texto)
	if err != nil {
		return nil, err
	}
	return mensajeGuardado, nil
}

func (servicio *ServicioChat) ConsultarHistorial() string {
	listaMensajes := servicio.almacenamiento.ObtenerHistorial()
	if len(listaMensajes) == 0 {
		return "No hay mensajes en el historial."
	}
	textoHistorial := "HISTORIAL DE MENSAJES:\n"
	for _, mensajeActual := range listaMensajes {
		textoHistorial = textoHistorial + mensajeActual.Mostrar() + "\n"
	}
	return textoHistorial
}