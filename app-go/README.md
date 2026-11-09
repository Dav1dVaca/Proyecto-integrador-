# Prototipo Funcional - VARGA Chat (Go)

## 1. Descripción de la Aplicación
VARGA Chat es el prototipo funcional de un sistema de mensajería empresarial bajo arquitectura cliente-servidor desarrollado en lenguaje Go para la organización VARGA Solutions S.A. Su propósito es permitir la comunicación interna entre las sedes de Quito (Matriz) y Cuenca (Sucursal), aplicando los principios de Programación Orientada a Objetos (estructuras, constructores, encapsulamiento, métodos con receptores por puntero y composición), modularización y validación de datos.

## 2. Funcionalidades Implementadas
* Identificación de usuarios registrados de Quito y Cuenca (`josue`, `david`, `adrian`).
* Inicio y cierre de sesión de comunicación desde el cliente.
* Envío de mensajes de texto en tiempo real desde el cliente hacia el servidor.
* Consulta del historial de conversaciones almacenadas durante la sesión del servidor.
* Validación de entradas (usuarios inexistentes, datos vacíos, mensajes sin contenido y opciones de menú inválidas) y manejo de errores de red.

## 3. Requisitos de Software y Versión de Go
* **Sistema Operativo:** Windows, Linux, macOS o entorno GitHub Codespaces.
* **Versión de Go utilizada:** Go 1.22 (o superior).
* **Librerías utilizadas:** Librería estándar de Go (`fmt` y `net`). No requiere instalación de paquetes externos.

## 4. Instrucciones de Instalación
1. Clonar el repositorio en su equipo o abrirlo en GitHub Codespaces:
   ```bash
   git clone [https://github.com/Dav1dVaca/Proyecto-integrador-.git](https://github.com/Dav1dVaca/Proyecto-integrador-.git)