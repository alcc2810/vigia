# Vigía Platform

# Documentation Standard

| Campo | Valor |
|--------|--------|
| Documento | Documentation Standard |
| ID | VIG-DOC-00 |
| Versión | 1.0.0 |
| Estado | Approved |
| Clasificación | Interno |
| Proyecto | Vigía Platform |
| Autor | Adrián Collantes |
| Última actualización | 2026-07-23 |

---

# Tabla de Contenido

1. Objetivo
2. Alcance
3. Principios de Documentación
4. Estructura de la Documentación
5. Estructura de un Documento
6. Convenciones de Nomenclatura
7. Convenciones de Markdown
8. Versionado
9. Estados de los Documentos
10. Referencias Cruzadas
11. Diagramas e Imágenes
12. Plantillas
13. README
14. Architecture Decision Records (ADR)
15. Control de Versiones (Git)
16. Convenciones de Commits
17. Buenas Prácticas
18. Revisión y Aprobación
19. Jerarquía Documental
20. Documentación Oficial
21. Historial de Cambios

---

# 1. Objetivo

Establecer el estándar oficial para la creación, mantenimiento, organización y versionado de toda la documentación del proyecto **Vigía Platform**, garantizando consistencia, trazabilidad y facilidad de mantenimiento durante todo el ciclo de vida del software.

---

# 2. Alcance

Este estándar aplica a toda la documentación almacenada dentro del directorio `docs/`.

Incluye, entre otros:

* Documentación funcional
* Documentación técnica
* Arquitectura
* Modelo de datos
* Protocolos
* APIs
* Seguridad
* Roadmap
* Especificaciones del dispositivo
* Guías de desarrollo
* Convenciones
* Architecture Decision Records (ADR)

No aplica a documentación temporal, pruebas o investigaciones almacenadas fuera del directorio `docs/`.

---

# 3. Principios de Documentación

Toda la documentación del proyecto deberá cumplir los siguientes principios:

* Claridad
* Precisión
* Consistencia
* Trazabilidad
* Mantenibilidad
* Reutilización
* Versionado
* Una única fuente de verdad
* Independencia tecnológica cuando sea posible

---

# 4. Estructura de la Documentación

La documentación oficial del proyecto se organiza de la siguiente manera:

```text
docs/
│
├── README.md
├── templates/
├── adr/
├── images/
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
├── 13_CONVENTIONS.md
└── CHANGELOG.md
```

---

# 5. Estructura de un Documento

Todos los documentos oficiales deberán comenzar con el siguiente encabezado:

```md
# Vigía Platform

# Nombre del Documento

| Campo | Valor |
|--------|--------|
| Documento | <Document Name> |
| ID | <Document ID> |
| Versión | <Version> |
| Estado | <Status> |
| Clasificación | <Classification> |
| Proyecto | Vigía Platform |
| Autor | <Author> |
| Última actualización | <YYYY-MM-DD> |

---

# Tabla de Contenido

```

Asimismo, todos los documentos deberán finalizar con un historial de cambios.

---

# 6. Convenciones de Nomenclatura

Los documentos oficiales deberán cumplir las siguientes reglas:

* Numeración de dos dígitos
* Nombre del archivo en mayúsculas
* Palabras separadas mediante guion bajo (`_`)
* Extensión `.md`

Ejemplos:

```text
00_DOCUMENTATION_STANDARD.md
01_PRODUCT_VISION.md
02_ARCHITECTURE.md
```

Los títulos visibles dentro de los documentos utilizarán formato legible para facilitar su lectura.

Ejemplo:

```text
Archivo:
02_ARCHITECTURE.md

Título:
Architecture
```

---

# 7. Convenciones de Markdown

Toda la documentación utilizará Markdown como lenguaje de edición.

Se establecen las siguientes reglas:

* Utilizar `#` para títulos principales
* Numerar todas las secciones
* Utilizar `*` para listas no ordenadas
* Utilizar listas numeradas para procedimientos
* Utilizar tablas cuando faciliten la comprensión de la información
* Especificar siempre el lenguaje en los bloques de código
* Dejar una línea en blanco antes y después de listas, tablas y bloques de código

---

# 8. Versionado

Todos los documentos oficiales utilizan **Semantic Versioning (SemVer)**.

La estructura es la siguiente:

```text
MAJOR.MINOR.PATCH
```

Donde:

* **MAJOR:** Cambios importantes
* **MINOR:** Nuevas secciones o ampliaciones
* **PATCH:** Correcciones menores

Ejemplos:

```text
1.0.0
1.0.1
1.1.0
2.0.0
```

Cada documento mantiene su propio versionado de forma independiente.

---

# 9. Estados de los Documentos

Los documentos podrán encontrarse en uno de los siguientes estados:

| Estado | Descripción |
|---------|-------------|
| Draft | Documento en elaboración |
| Review | Documento en proceso de revisión |
| Approved | Documento aprobado para su utilización |
| Deprecated | Documento reemplazado o fuera de uso |

---

# 10. Referencias Cruzadas

Toda referencia entre documentos, diagramas, recursos o componentes deberá realizarse mediante identificadores únicos.

Ejemplos:

| Prefijo | Descripción |
|---------|-------------|
| VIG-DOC | Documentación oficial |
| VIG-ADR | Architecture Decision Records |
| VIG-ARC | Diagramas de arquitectura |
| VIG-DB | Diagramas de base de datos |
| VIG-UI | Interfaces gráficas |
| VIG-WF | Wireframes |
| VIG-BR | Reglas de negocio |
| VIG-API | Especificaciones de API |

---

# 11. Diagramas e Imágenes

Todos los recursos gráficos deberán almacenarse dentro del directorio `docs/images/`.

Las imágenes deberán organizarse por categoría.

```text
architecture/
database/
ui/
wireframes/
```

Cada recurso deberá utilizar un identificador único y un nombre descriptivo.

Ejemplo:

```text
VIG-ARC-01-System-Overview.drawio
VIG-DB-01-ERD.drawio
VIG-UI-01-Login.png
```

---

# 12. Plantillas

Las plantillas oficiales deberán almacenarse dentro del directorio `docs/templates/`.

Plantillas definidas:

* DOCUMENT_TEMPLATE.md
* ADR_TEMPLATE.md
* README_TEMPLATE.md
* DIAGRAM_TEMPLATE.md

---

# 13. README

Toda carpeta principal del proyecto deberá contener un archivo `README.md`.

El propósito del README es describir el contenido y la finalidad de la carpeta correspondiente.

Los README deberán mantenerse sincronizados con la estructura del proyecto.

---

# 14. Architecture Decision Records (ADR)

Las decisiones arquitectónicas deberán documentarse mediante Architecture Decision Records (ADR).

Los ADR deberán almacenarse en el directorio:

```text
docs/adr/
```

Convenciones:

* Cada ADR tendrá un identificador único.
* Los ADR aprobados constituyen un registro histórico.
* Un ADR aprobado no deberá modificarse; cualquier cambio se documentará mediante un nuevo ADR.

---

# 15. Control de Versiones (Git)

El proyecto utilizará la siguiente estrategia de ramas:

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

---

# 16. Convenciones de Commits

El proyecto utilizará el estándar **Conventional Commits** para mantener un historial de cambios consistente y facilitar el seguimiento de la evolución del software.

Prefijos permitidos:

* feat
* fix
* docs
* refactor
* test
* build
* style
* chore

Para la documentación oficial, los mensajes de commit deberán incluir el identificador (ID) y el nombre del documento cuando corresponda.

Ejemplos:

```text
docs: create VIG-DOC-00 Documentation Standard
docs: create VIG-DOC-01 Product Vision
docs: update VIG-DOC-02 Architecture
docs: update VIG-DOC-04 Protocol Specification

feat: implement device registration
fix: correct GPS parser
refactor: simplify rule engine
```

---

# 17. Buenas Prácticas

Toda la documentación deberá:

* Mantener una única fuente de verdad por tema.
* Evitar la duplicidad de información.
* Permanecer alineada con la evolución del software.
* Utilizar un lenguaje claro y objetivo.
* Actualizar su historial de cambios cuando corresponda.

---

# 18. Revisión y Aprobación

Todo documento oficial seguirá el siguiente ciclo de vida:

```text
Draft
   ↓
Review
   ↓
Approved
   ↓
Deprecated
```

Ningún documento deberá marcarse como **Approved** sin haber completado previamente su proceso de revisión.

---

# 19. Jerarquía Documental

En caso de conflicto entre documentos, prevalecerá el siguiente orden de prioridad:

1. VIG-DOC-00 — Documentation Standard.
2. Architecture Decision Records (ADR).
3. Documentación oficial.
4. README de directorios.
5. Comentarios en el código fuente.

---

# 20. Documentación Oficial

La documentación oficial del proyecto está compuesta por:

| ID | Documento |
|----|-----------|
| VIG-DOC-00 | Documentation Standard |
| VIG-DOC-01 | Product Vision |
| VIG-DOC-02 | Architecture |
| VIG-DOC-03 | Data Model |
| VIG-DOC-04 | Protocol Specification |
| VIG-DOC-05 | Rule Engine |
| VIG-DOC-06 | API Specification |
| VIG-DOC-07 | Security |
| VIG-DOC-08 | Deployment |
| VIG-DOC-09 | Roadmap |
| VIG-DOC-10 | Glossary |
| VIG-DOC-11 | Device Specification |
| VIG-DOC-12 | Development Guide |
| VIG-DOC-13 | Conventions |

---

# 21. Historial de Cambios

| Versión | Fecha | Autor | Descripción |
|----------|------------|-------------------|-----------------------------------------|
| 1.0.0 | 2026-07-23 | Adrián Collantes | Creación del estándar oficial de documentación de Vigía Platform. |