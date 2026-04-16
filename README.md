# API-REST-GR2-TADEO
API REST de gestion de tareas con Go

# Objetivo General
Diseñar e implementar una API REST funcional y segura usando el lenguaje Go, aplicando principios de arquitectura de software limpia, persistencia de datos con SQLite y autenticación mediante tokens JWT, desarrollando competencias en el ciclo completo de desarrollo backend.

OE · 01
Fundamentos del Lenguaje
Comprender la sintaxis y tipos de datos primitivos de Go
Definir structs, interfaces y métodos para modelar el dominio
Aplicar el manejo idiomático de errores como valores de retorno
Organizar código en paquetes y módulos con go mod


OE · 02
Desarrollo de API REST
Construir un servidor HTTP con el paquete nativo net/http
Implementar los métodos CRUD completos: GET, POST, PUT, DELETE, PATCH
Serializar y deserializar datos con encoding/json
Diseñar rutas RESTful con manejo de parámetros de ruta


OE · 03
Persistencia de Datos
Conectar la aplicación a SQLite usando database/sql
Implementar el patrón Repository para separar capas de datos
Ejecutar migraciones para crear esquemas de tablas relacionales
Gestionar valores nulos con sql.NullTime y foreign keys

OE · 04
Seguridad y Calidad
Implementar autenticación con JWT y contraseñas seguras con bcrypt
Construir middlewares encadenados para logging y autorización
Aplicar inyección de dependencias mediante el contexto HTTP
Escribir pruebas unitarias con testing y httptest


# Plan de desarrollo · semana a semana

01
Fundamentos Go + Servidor HTTP
Sintaxis básica de Go
Structs e interfaces
Servidor net/http
Endpoint /health
3h · Semana 1

02
Modelos y CRUD en Memoria
Struct Task y User
Almacén en map
Handlers CRUD
Validaciones
3h · Semana 2

03
Persistencia con SQLite
database/sql
Migraciones
Patrón Repository
Foreign keys
3h · Semana 3

04
Autenticación JWT
Registro + bcrypt
Login + JWT
Middleware Auth
Context values
3h · Semana 4


05
Pruebas y Presentación
testing + httptest
Variables de entorno
Documentación API
Demo final
3h · Semana 5

# Stack tecnológico
⚙️
Go 1.21+
Lenguaje principal

🌐
net/http
Servidor nativo

🗄️
SQLite
Base de datos


🔐
JWT + bcrypt
Autenticación

🧪
testing
Pruebas unitarias
