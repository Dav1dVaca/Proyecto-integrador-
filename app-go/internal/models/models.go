package models

import (
	"fmt"
)

type Persona struct {
	nombre string
	sede   string
}

func NuevaPersona(nombre string, sede string) (*Persona, error) {
	if nombre == "" {
		return nil, fmt.Errorf("el nombre no puede estar vacio")
	}
	if sede == "" {
		return nil, fmt.Errorf("la sede no puede estar vacia")
	}
	return &Persona{
		nombre: nombre,
		sede:   sede,
	}, nil
}

func (persona *Persona) ObtenerNombre() string {
	return persona.nombre
}

func (persona *Persona) CambiarNombre(nombre string) error {
	if nombre == "" {
		return fmt.Errorf("el nombre no puede estar vacio")
	}
	persona.nombre = nombre
	return nil
}

func (persona *Persona) ObtenerSede() string {
	return persona.sede
}

func (persona *Persona) CambiarSede(sede string) error {
	if sede == "" {
		return fmt.Errorf("la sede no puede estar vacia")
	}
	persona.sede = sede
	return nil
}

type Usuario struct {
	Persona
	cuenta string
	rol    string
	activo bool
}

func NuevoUsuario(cuenta string, nombre string, sede string, rol string) (*Usuario, error) {
	if cuenta == "" {
		return nil, fmt.Errorf("el usuario no puede estar vacio")
	}
	if rol == "" {
		return nil, fmt.Errorf("el rol no puede estar vacio")
	}
	nuevaPersona, err := NuevaPersona(nombre, sede)
	if err != nil {
		return nil, err
	}
	return &Usuario{
		Persona: *nuevaPersona,
		cuenta:  cuenta,
		rol:     rol,
		activo:  false,
	}, nil
}

func (usuario *Usuario) ObtenerCuenta() string {
	return usuario.cuenta
}

func (usuario *Usuario) CambiarCuenta(cuenta string) error {
	if cuenta == "" {
		return fmt.Errorf("el usuario no puede estar vacio")
	}
	usuario.cuenta = cuenta
	return nil
}

func (usuario *Usuario) ObtenerRol() string {
	return usuario.rol
}

func (usuario *Usuario) CambiarRol(rol string) error {
	if rol == "" {
		return fmt.Errorf("el rol no puede estar vacio")
	}
	usuario.rol = rol
	return nil
}

func (usuario *Usuario) ObtenerActivo() bool {
	return usuario.activo
}

func (usuario *Usuario) CambiarActivo(estado bool) {
	usuario.activo = estado
}

type Mensaje struct {
	autor string
	sede  string
	texto string
}

func NuevoMensaje(autor string, sede string, texto string) (*Mensaje, error) {
	if autor == "" {
		return nil, fmt.Errorf("el autor no puede estar vacio")
	}
	if texto == "" {
		return nil, fmt.Errorf("el mensaje no puede estar vacio")
	}
	return &Mensaje{
		autor: autor,
		sede:  sede,
		texto: texto,
	}, nil
}

func (mensaje *Mensaje) ObtenerAutor() string {
	return mensaje.autor
}

func (mensaje *Mensaje) ObtenerSede() string {
	return mensaje.sede
}

func (mensaje *Mensaje) ObtenerTexto() string {
	return mensaje.texto
}

func (mensaje *Mensaje) CambiarTexto(texto string) error {
	if texto == "" {
		return fmt.Errorf("el mensaje no puede estar vacio")
	}
	mensaje.texto = texto
	return nil
}

func (mensaje *Mensaje) Mostrar() string {
	return "[" + mensaje.sede + "] " + mensaje.autor + ": " + mensaje.texto
}