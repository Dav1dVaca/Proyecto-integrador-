package storage

import (
	"chat/internal/models"
	"fmt"
)

type Almacenamiento struct {
	usuarios map[string]*models.Usuario
	mensajes []*models.Mensaje
}

func NuevoAlmacenamiento() *Almacenamiento {
	almacenamiento := &Almacenamiento{
		usuarios: make(map[string]*models.Usuario),
		mensajes: make([]*models.Mensaje, 0),
	}
	almacenamiento.cargarUsuarios()
	return almacenamiento
}

func (almacenamiento *Almacenamiento) cargarUsuarios() {
	usuario1, _ := models.NuevoUsuario("josue", "Josue_Ramirez", "Quito", "Administrador")
	usuario2, _ := models.NuevoUsuario("david", "David_Vaca", "Quito", "Desarrollo")
	usuario3, _ := models.NuevoUsuario("adrian", "Adrian_Garcia", "Cuenca", "Soporte")

	almacenamiento.usuarios[usuario1.ObtenerCuenta()] = usuario1
	almacenamiento.usuarios[usuario2.ObtenerCuenta()] = usuario2
	almacenamiento.usuarios[usuario3.ObtenerCuenta()] = usuario3
}

func (almacenamiento *Almacenamiento) BuscarUsuario(cuenta string) (*models.Usuario, error) {
	usuarioEncontrado, existe := almacenamiento.usuarios[cuenta]
	if !existe {
		return nil, fmt.Errorf("el usuario no existe")
	}
	return usuarioEncontrado, nil
}

func (almacenamiento *Almacenamiento) GuardarMensaje(autor string, sede string, texto string) (*models.Mensaje, error) {
	nuevoMensaje, err := models.NuevoMensaje(autor, sede, texto)
	if err != nil {
		return nil, err
	}
	almacenamiento.mensajes = append(almacenamiento.mensajes, nuevoMensaje)
	return nuevoMensaje, nil
}

func (almacenamiento *Almacenamiento) ObtenerHistorial() []*models.Mensaje {
	return almacenamiento.mensajes
}