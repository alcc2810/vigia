# Vigía Platform

# 00_DOCUMENTATION_STANDARD

**Versión:** 1.0.0

**Estado:** Approved

**Autor:** Adrián Collantes

**Última actualización:** 2026-07-23

---

# Tabla de Contenido

1. Objetivo
2. Alcance
3. Estructura de la documentación
4. Convenciones de nombres
5. Estructura de los documentos
6. Versionado
7. Estados de los documentos
8. Historial de cambios
9. Convenciones de Markdown
10. Convenciones para diagramas e imágenes
11. Convenciones para ADR
12. Convenciones para README
13. Convenciones para Git
14. Convenciones para commits
15. Buenas prácticas
16. Documentos oficiales

---

# 1. Objetivo

Establecer el estándar oficial para la creación, mantenimiento y versionado de toda la documentación técnica y funcional del proyecto Vigía Platform.

Este documento define las reglas que deberán seguir todos los documentos del repositorio para garantizar consistencia, trazabilidad y facilidad de mantenimiento.

---

# 2. Alcance

Este estándar aplica a toda la documentación almacenada dentro del directorio `/docs`.

Incluye:

* Documentación funcional.
* Documentación técnica.
* Arquitectura.
* APIs.
* Protocolos.
* Seguridad.
* Roadmaps.
* Architecture Decision Records (ADR).

No aplica a documentación temporal almacenada dentro de `/research`.

---

# 3. Estructura de la documentación

```text
docs/
│
├── templates/
│   ├── DOCUMENT_TEMPLATE.md
│   └── ADR_TEMPLATE.md
│
├── adr/
│
├── images/
│   ├── architecture/
│   ├── database/
│   ├── ui/
│   └── wireframes/
│
├── 00_DOCUMENTATION_STANDARD.md
├── 01_PRODUCT_VISION.md
├── 02_ARCHITECTURE.md
├── 03_DATA_MODEL.md
├── 04_PROTOCOL_SPEC.md
├── 05_RULE_ENGINE.md
├── 06_API_SPEC.md
├── 07_SECURITY.md
├── 08_DEPLOYMENT.md
├── 09_ROADMAP.md
├── 10_GLOSSARY.md
├── 11_DEVICE_SPEC.md
├── 12_DEVELOPMENT_GUIDE.md
└── CHANGELOG.md
```

---

# 4. Convenciones de nombres

Todos los documentos oficiales deberán cumplir las siguientes reglas:

* Numeración de dos dígitos.
* Nombre en MAYÚSCULAS.
* Palabras separadas mediante "_".
* Extensión `.md`.

Ejemplos:

```text
01_PRODUCT_VISION.md
02_ARCHITECTURE.md
03_DATA_MODEL.md
```

---

# 5. Estructura de los documentos

Todo documento oficial deberá comenzar con el siguiente encabezado:

```md
# Vigía Platform

# Nombre del Documento

**Versión:** x.y.z

**Estado:** Draft | Review | Approved | Deprecated

**Autor:** Nombre del autor

**Última actualización:** AAAA-MM-DD
```

Posteriormente deberá incluir:

* Tabla de contenido.
* Desarrollo del documento.
* Historial de cambios.

---

# 6. Versionado

La documentación utilizará Semantic Versioning.

```text
MAJOR.MINOR.PATCH
```

* **MAJOR:** cambios importantes.
* **MINOR:** nuevas secciones o ampliaciones.
* **PATCH:** correcciones menores.

Ejemplo:

```text
1.0.0
1.0.1
1.1.0
2.0.0
```

---

# 7. Estados de los documentos

| Estado | Descripción |
|---------|-------------|
| Draft | Documento en elaboración |
| Review | Documento en revisión |
| Approved | Documento aprobado |
| Deprecated | Documento obsoleto o reemplazado |

---

# 8. Historial de cambios

Todos los documentos finalizarán con:

| Versión | Fecha | Autor | Descripción |
|----------|-------|--------|-------------|

---

# 9. Convenciones de Markdown

* Utilizar `#` para el título principal.
* Numerar todas las secciones.
* Utilizar `*` para listas no ordenadas.
* Utilizar listas numeradas para procedimientos.
* Utilizar tablas para información comparativa.
* Especificar siempre el lenguaje en bloques de código.
* Dejar una línea en blanco antes y después de listas, tablas y bloques de código.

---

# 10. Convenciones para diagramas e imágenes

Todas las imágenes deberán almacenarse en:

```text
docs/images/
```

Organizadas en:

* architecture/
* database/
* ui/
* wireframes/

Nombres descriptivos:

```text
system-overview.drawio
deployment-diagram.png
database-er.png
```

---

# 11. Convenciones para ADR

Los Architecture Decision Records deberán almacenarse en:

```text
docs/adr/
```

Formato:

```text
ADR-0001-Database-Selection.md
ADR-0002-Microservices.md
ADR-0003-Protocol.md
```

Reglas:

* Nunca modificar un ADR aprobado.
* Crear un nuevo ADR cuando una decisión cambie.
* Todos los ADR deberán utilizar la plantilla oficial.

---

# 12. Convenciones para README

Cada carpeta principal del proyecto deberá contener un archivo `README.md`.

Contenido mínimo:

* Objetivo.
* Descripción.
* Estructura interna.
* Tecnologías utilizadas (si aplica).
* Responsable (opcional).

Ejemplo:

```text
backend/
frontend/
database/
hardware/
firmware/
research/
```

---

# 13. Convenciones para Git

Modelo de ramas oficial:

```text
main
│
└── develop
     ├── feature/*
     ├── bugfix/*
     ├── hotfix/*
     ├── release/*
     ├── docs/*
     └── research/*
```

Descripción:

* **main:** versión estable.
* **develop:** integración del desarrollo.
* **feature/***: nuevas funcionalidades.
* **bugfix/***: corrección de errores.
* **hotfix/***: correcciones urgentes.
* **release/***: preparación de nuevas versiones.
* **docs/***: cambios en documentación.
* **research/***: investigación y experimentación.

---

# 14. Convenciones para commits

Se utilizará Conventional Commits.

Ejemplos:

```text
feat: implement device registration

fix: correct geofence calculation

docs: update architecture

refactor: simplify notification service

test: add unit tests

build: update docker image

chore: update dependencies
```

Prefijos permitidos:

* feat
* fix
* docs
* refactor
* test
* build
* style
* chore

---

# 15. Buenas prácticas

Toda la documentación deberá cumplir los siguientes principios:

* Claridad.
* Simplicidad.
* Consistencia.
* Trazabilidad.
* Reutilización.
* Independencia tecnológica cuando sea posible.
* Actualización continua.
* Una única fuente de verdad por tema.
* Evitar duplicar información entre documentos.

---

# 16. Documentos oficiales

| Nº | Documento |
|----|-----------|
| 00 | DOCUMENTATION_STANDARD |
| 01 | PRODUCT_VISION |
| 02 | ARCHITECTURE |
| 03 | DATA_MODEL |
| 04 | PROTOCOL_SPEC |
| 05 | RULE_ENGINE |
| 06 | API_SPEC |
| 07 | SECURITY |
| 08 | DEPLOYMENT |
| 09 | ROADMAP |
| 10 | GLOSSARY |
| 11 | DEVICE_SPEC |
| 12 | DEVELOPMENT_GUIDE |
| -- | CHANGELOG |

---

# Historial de Cambios

| Versión | Fecha | Autor | Descripción |
|----------|------------|-------------------|-----------------------------------------|
| 1.0.0 | 2026-07-23 | Adrián Collantes | Creación del estándar documental del proyecto. |