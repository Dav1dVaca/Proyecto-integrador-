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

