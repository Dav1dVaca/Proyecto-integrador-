# VARGA Chat – Plataforma Empresarial Distribuida

Proyecto Integrador de Tercer Semestre de Ingeniería en Sistemas de Información (UIDE).  
Sistema interno de mensajería para empresas, control de acceso y red entre dos sedes para la organización VARGA Solutions.

---

## 1. Descripción del Sistema

VARGA Chat es una aplicación cliente-servidor desarrollada en Go para resolver los problemas de comunicación, fuga de datos importantes y pérdida de historial entre la matriz y la sucursal de la empresa VARGA Solutions.

El sistema permite que los empleados chateen en tiempo real de forma segura, utiliza un servicio adicional en FastAPI para verificar usuarios y contraseñas mediante pases digitales de acceso, y una base de datos local para guardar los mensajes y consultarlos cuando sea necesario, además, funciona sobre una red privada empresarial conectada entre ambas ciudades.

---

## 2. Integrantes del Equipo

| Integrante | Responsabilidades Generales |
| :--- | :--- |
| **Josue Ramírez** | Encargado de la seguridad, el servicio de inicio de sesión con FastAPI y la protección de contraseñas de los usuarios. |
| **David Vaca** | Encargado de programar el cliente y el servidor de chat en Go, y de guardar los mensajes en la base de datos. |
| **Adrian Garcia** | Encargado de configurar la red que une Quito y Cuenca, y de tomar las medidas numéricas para el análisis de estadística. |

---

---

## 3. Alcance Preliminar 
### Lo que hemos completado en esta etapa
- [x] Descripción del problema de la empresa entre Quito y Cuenca.
- [x] Justificación de por qué elegimos el chat empresarial y el lenguaje Go
- [x] Lista de requisitos organizados con códigos para poder probarlos después
- [x] Diagrama general y descripción paso a paso de los 3 casos de uso principales.
- [x] Matriz de trazabilidad que conecta necesidades, objetivos y requisitos
- [x] Estructura inicial del repositorio en Git y reparto de trabajo
### Lo que desarrollaremos en las siguientes fases
- [ ] Definir el formato exacto en el que viajarán los textos de los mensajes.
- [ ] Asignar las direcciones de red para las computadoras de Quito y Cuenca
- [ ] Programar el sistema de inicio de sesión en FastAPI
- [ ] Programar el servidor y el programa de chat en Go
- [ ] Probar la velocidad de envío y retraso de los mensajes para los cálculos de estadística

---
## 4. Estructura Inicial del Repositorio

El proyecto se organiza en las siguientes carpetas[cite: 2]:

```text
proyecto-integrador/
├── app-go/             # Programa del chat en Go (servidor y cliente)
│   ├── cmd/            # Archivos para ejecutar el servidor y el cliente
│   ├── internal/       # Lógica del chat y guardado de mensajes en SQLite
│   └── README.md       # Guía de uso del programa en Go
├── auth-service/       # Servicio de inicio de sesión en FastAPI
│   ├── app/            # Código para comprobar usuarios y contraseñas
│   ├── tests/          # Pruebas de acceso
│   └── requirements.txt# Librerías de Python necesarias
├── network/            # Archivos y respaldos de la red
│   ├── configs/        # Respaldos de configuración de routers y switches
│   ├── addressing/     # Lista de direcciones de red para Quito y Cuenca
│   └── diagrams/       # Dibujos y diagramas de la red
├── research/           # Carpeta para la materia de Estadística
│   ├── instruments/    # Guía de cómo se toman las mediciones
│   ├── data/           # Archivos con los datos de retraso y velocidad
│   └── analysis/       # Gráficos y cálculos estadísticos
├── docs/               # Documentos e informes de la materia
├── LICENSE             # Licencia libre del proyecto
└── README.md           # Explicación general del repositorio
