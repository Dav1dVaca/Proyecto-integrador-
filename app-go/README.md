# Prototipo Funcional - VARGA Chat (Go)

## 1. Descripción de la Aplicación
VARGA Chat es el prototipo funcional de un sistema de mensajería empresarial bajo arquitectura cliente-servidor desarrollado en lenguaje Go para la organización VARGA Solutions. Su propósito es permitir la comunicación interna entre las sedes de Quito (Matriz) y Cuenca (Sucursal), aplicando los principios de Programación Orientada a Objetos (estructuras, constructores, encapsulamiento, métodos con receptores por puntero y composición), modularización y validación de datos.

## 2. Funcionalidades Implementadas
* Configuración de dirección IP y puerto tanto en el servidor como en el cliente al iniciar la ejecución.
* Identificación de usuarios registrados de Quito y Cuenca (`josue`, `david`, `adrian`).
* Inicio y finalización de sesión de comunicación desde el cliente.
* Envío de mensajes de texto en tiempo real desde el cliente hacia el servidor a través de sockets TCP.
* Consulta del historial básico de conversaciones almacenadas durante la sesión del servidor.
* Validación de datos de entrada (usuarios inexistentes, entradas vacías, mensajes sin texto y opciones de menú no válidas) y manejo explícito de errores sin uso de `panic`.

## 3. Requisitos de Software y Versión de Go
* **Sistema Operativo:** Windows, Linux, macOS o entorno GitHub Codespaces.
* **Versión de Go utilizada:** Go 1.22 (o superior).
* **Librerías utilizadas:** Únicamente librería estándar de Go (`fmt` y `net`). No requiere instalar dependencias externas.

## 4. Instrucciones de Instalación
1. Clonar el repositorio en su computadora o abrirlo directamente en GitHub Codespaces:
   ```bash
   git clone [https://github.com/Dav1dVaca/Proyecto-integrador-.git](https://github.com/Dav1dVaca/Proyecto-integrador-.git)
   ```
2. Ingresar a la carpeta de la aplicación en Go:
   ```bash
   cd Proyecto-integrador-/app-go
   ```

## 5. Instrucciones para Iniciar el Servidor
En una primera terminal ubicada dentro de la carpeta `app-go`, ejecutar:
```bash
go run cmd/server/main.go
```
* El programa pedirá ingresar la IP y el puerto. Presione `Enter` en ambos campos para utilizar los valores por defecto (`127.0.0.1` y `9000`).

## 6. Instrucciones para Ejecutar el Cliente
En una segunda terminal independiente (ubicada también dentro de la carpeta `app-go`), ejecutar:
```bash
go run cmd/client/main.go
```
* Ingrese la misma dirección IP y puerto donde se encuentra activo el servidor (o presione `Enter` para conectarse a `127.0.0.1:9000`).

## 7. Dirección IP, Puerto y Parámetros Configurables
* **Dirección IP por defecto:** `127.0.0.1` (configurable al inicio del programa para pruebas locales o de red).
* **Puerto por defecto:** `9000` (configurable al inicio del programa).
* **Protocolo de comunicación:** TCP (`net.Listen` en el servidor y `net.Dial` en el cliente).

## 8. Datos y Archivos de Prueba
Los usuarios precargados en memoria para comprobar el funcionamiento del prototipo están documentados en `data/usuarios.txt`:
* `josue` (Josue_Ramirez - Sede Quito - Rol Administrador)
* `david` (David_Vaca - Sede Quito - Rol Desarrollo)
* `adrian` (Adrian_Garcia - Sede Cuenca - Rol Soporte)

**Escenarios de prueba recomendados:**
1. **Validación de usuario inexistente:** Ingresar un nombre no registrado (ej. `carlos`) y verificar el mensaje `ERROR: el usuario no existe`.
2. **Inicio de sesión válido:** Ingresar `josue` o `adrian` y verificar el mensaje de bienvenida con su sede.
3. **Validación de mensaje vacío:** Elegir la opción `1` (Enviar mensaje) y presionar `Enter` sin escribir texto para comprobar el mensaje `ERROR: el mensaje no puede estar vacio`.
4. **Envío y consulta de historial:** Enviar un mensaje válido con la opción `1` y luego elegir la opción `2` (Ver historial) para verificar que quedó guardado con sede y autor.

## 9. Estructura de Carpetas
```text
app-go/
├── cmd/
│   ├── client/
│   │   └── main.go          # Proceso independiente del cliente con menú interactivo
│   └── server/
│       └── main.go          # Proceso independiente del servidor TCP concurrente
├── internal/
│   ├── models/
│   │   └── models.go        # Estructuras Persona, Usuario y Mensaje con constructores y métodos
│   ├── services/
│   │   └── services.go      # ServicioChat con reglas de negocio y validaciones
│   ├── handlers/
│   │   └── handlers.go      # ManejadorRed para procesar comandos TCP del cliente
│   └── storage/
│       └── storage.go       # Almacenamiento en memoria mediante map y slice
├── data/
│   └── usuarios.txt         # Detalle de usuarios de prueba precargados
├── go.mod                   # Archivo de módulo Go
└── README.md                # Manual técnico del prototipo
```

## 10. Limitaciones Conocidas
* **Almacenamiento temporal en memoria:** En esta fase (Semana 10), los usuarios y mensajes se almacenan en estructuras en memoria (`map` y `slice`). Si el servidor se apaga, el historial se reinicia.
* **Autenticación preparada para integración:** Siguiendo los lineamientos de la guía, no se duplicó un motor definitivo de contraseñas dentro de Go, dejando preparado el flujo de sesión para conectarse con el servicio FastAPI y la base de datos SQLite en las semanas 11 y 12.

## 11. Autores y Contribuciones
* **Josue Ramírez:** Implementación de modelos POO (`Persona`, `Usuario`, `Mensaje`), constructores, encapsulamiento y validaciones de entrada.
* **David Vaca:** Implementación del almacenamiento en memoria (`Almacenamiento`), lógica del servicio (`ServicioChat`) y cliente interactivo.
* **Adrian Garcia:** Implementación del manejador de red TCP (`ManejadorRed`), servidor concurrente y pruebas de comunicación.

## 12. Créditos, Licencia y Declaración de Uso de IA
* **Licencia:** Proyecto distribuido bajo la **Licencia MIT** (consultar el archivo `LICENSE` en la raíz del repositorio).
* **Créditos:** Universidad Internacional del Ecuador (UIDE) - Facultad de Ingenierías y Ciencias Aplicadas - Carrera de Ingeniería en Sistemas de Información.
* **Declaración del uso de herramientas de Inteligencia Artificial:** Se utilizó herramientas de inteligencia artificial como apoyo para la organización modular de paquetes, consulta de sintaxis básica en Go y formato de documentación. Todo el código fue revisado, probado y explicado por los integrantes del equipo.