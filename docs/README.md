# VARGA Chat – Plataforma Empresarial Distribuida

Proyecto Integrador de Tercer Semestre de Ingeniería en Sistemas de Información (UIDE).  
Sistema interno de mensajería para empresas, control de acceso y red entre dos sedes para la organización **VARGA Solutions**

## 1. Descripción del Sistema

VARGA Chat es una aplicación cliente-servidor desarrollada en lenguaje Go para resolver los problemas de comunicación, fuga de datos importantes y pérdida de historial técnico entre la sede matriz (Quito) y la sucursal (Cuenca) de la empresa VARGA Solutions S.A.

El sistema permite que los empleados se comuniquen en tiempo real mediante sockets TCP aplicando los pilares de Programación Orientada a Objetos (estructuras, constructores, encapsulamiento, métodos con receptores por puntero y composición). Además, cuenta con el diseño de seguridad para autenticación mediante FastAPI, almacenamiento de historial y segmentación de red entre ambas ciudades.

## 2. Integrantes del Equipo y Contribuciones

| Integrante | Responsabilidades y Contribuciones |
| :--- | :--- |
| **Josue Ramírez** | Diseño de estructuras POO (`internal/models`), validaciones de datos, diseño de arquitectura de seguridad (Fase 1 y Fase 2) y plan de pruebas de seguridad. |
| **David Vaca** | Desarrollo del almacenamiento en memoria (`internal/storage`), lógica de negocio del servicio de chat (`internal/services`) y programación del cliente en Go (`cmd/client`). |
| **Adrian Garcia** | Implementación del manejador de red TCP (`internal/handlers`), programación del servidor en Go (`cmd/server`), diseño de red inter-sedes y documentación técnica. |

## 3. Estado y Avance del Proyecto

### Lo que hemos completado hasta la Semana 10
- [x] Descripción del problema de la empresa entre Quito y Cuenca y justificación tecnológica.
- [x] Definición de requisitos funcionales, requisitos de seguridad, casos de uso y matriz de trazabilidad.
- [x] Diseño de seguridad (Fase 2): flujo de autenticación, matriz de autorización por roles, protección de credenciales (almacenamiento vs. transmisión), controles de seguridad y plan de pruebas.
- [x] Definición del protocolo de aplicación por comandos de texto (`LOGIN`, `ENVIAR`, `HISTORIAL`, `SALIR`) sobre sockets TCP.
- [x] **Prototipo funcional cliente-servidor en Go (`app-go/`):**
  - Procesos independientes para servidor y cliente con IP y puerto configurables.
  - Estructuras encapsuladas (`Persona`, `Usuario`, `Mensaje`), constructores, métodos y composición.
  - Gestión de usuarios y mensajes en memoria (`map` y `slice`).
  - Validación de entradas vacías, usuarios inexistentes y manejo explícito de errores.

### Lo que desarrollaremos en las siguientes fases (Semanas 11 y 12)
- [ ] Integración del cliente-servidor en Go con el servicio de autenticación en FastAPI.
- [ ] Conexión con base de datos SQLite para persistencia permanente y hash de contraseñas.
- [ ] Pruebas de comunicación sobre la topología de red entre Quito y Cuenca y toma de mediciones estadísticas.

## 4. Estructura Actual del Repositorio

El repositorio se organiza de forma modular de la siguiente manera:

```text
Proyecto-integrador-/
├── app-go/                     # Prototipo funcional del chat en Go
│   ├── cmd/
│   │   ├── client/
│   │   │   └── main.go         # Programa ejecutable del cliente
│   │   └── server/
│   │       └── main.go         # Programa ejecutable del servidor
│   ├── internal/
│   │   ├── models/
│   │   │   └── models.go       # Estructuras Persona, Usuario y Mensaje (POO)
│   │   ├── services/
│   │   │   └── services.go     # Lógica de negocio y validaciones del chat
│   │   ├── handlers/
│   │   │   └── handlers.go     # Manejador de conexiones y comandos TCP
│   │   └── storage/
│   │       └── storage.go      # Almacenamiento en memoria (map y slice)
│   ├── data/
│   │   └── usuarios.txt        # Usuarios precargados para pruebas
│   ├── go.mod                  # Definición del módulo en Go (Go 1.22)
│   └── README.md               # Manual detallado del prototipo en Go
├── docs/                       # Documentación e informes de avance
├── LICENSE                     # Licencia MIT del proyecto
└── README.md                   # Descripción general del proyecto
```

## 5. Ejecución Rápida del Prototipo en Go

Para probar el prototipo funcional desarrollado en esta entrega (ver detalles completos en [`app-go/README.md`](./app-go/README.md)):

1. **Iniciar el servidor** (Terminal 1):
   ```bash
   cd app-go
   go run cmd/server/main.go
   ```
2. **Iniciar el cliente** (Terminal 2):
   ```bash
   cd app-go
   go run cmd/client/main.go
   ```
   *Parámetros por defecto:* IP `127.0.0.1`, Puerto `9000`.  
   *Usuarios de prueba:* `josue`, `david`, `adrian`.

## 6. Licencia Seleccionada

Este proyecto se distribuye bajo la **Licencia MIT**. Consulta el archivo [`LICENSE`](./LICENSE) en la raíz del repositorio para conocer los términos completos de uso y distribución.

## 7. Créditos y Declaración del Uso de Herramientas de IA

* **Créditos Académicos:** Proyecto desarrollado para las asignaturas de *Programación Orientada a Objetos*, *Fundamentos de la Seguridad de los Sistemas* y *Tecnologías de Comunicación de Información en Red* del Tercer Semestre de Ingeniería en Sistemas de Información en la Universidad Internacional del Ecuador (UIDE).
* **Declaración del uso de herramientas de Inteligencia Artificial:** Durante el desarrollo del proyecto se empleó herramientas de inteligencia artificial como apoyo de consulta para estructurar la organización modular del repositorio, revisar buenas prácticas de encapsulamiento en Go y dar formato a las plantillas de documentación técnica. Todo el código fuente, validaciones y pruebas de ejecución fueron implementados, verificados y comprendidos íntegramente por los miembros del equipo.