# Documentación de Vigía Platform

Bienvenido a la documentación oficial de **Vigía Platform**.

Este directorio contiene toda la documentación funcional, técnica y de arquitectura del proyecto. Su objetivo es servir como fuente única de información para el diseño, desarrollo, despliegue y mantenimiento de la plataforma.

# Tabla de Contenido

1. Objetivos
2. Estructura
3. Orden recomendado de lectura
4. Documentación Oficial
5. Directorios
6. Versionado
7. Referencias
8. Documento de Referencia
9. Licencia

---

# 1. Objetivos

* Centralizar la documentación oficial del proyecto.
* Mantener la trazabilidad de las decisiones técnicas.
* Facilitar la incorporación de nuevos colaboradores.
* Servir como referencia durante todo el ciclo de vida del producto.
* Garantizar la consistencia mediante estándares documentales.

---

# 2. Estructura

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

# 3. Orden recomendado de lectura

Para comprender la plataforma de forma progresiva, se recomienda consultar los documentos en el siguiente orden:

1. VIG-DOC-000 — DOCUMENTATION_STANDARD
2. VIG-DOC-001 — PRODUCT_VISION
3. VIG-DOC-002 — ARCHITECTURE
4. VIG-DOC-003 — DATA_MODEL
5. VIG-DOC-004 — PROTOCOL_SPEC
6. VIG-DOC-005 — RULE_ENGINE
7. VIG-DOC-006 — API_SPEC
8. VIG-DOC-007 — SECURITY
9. VIG-DOC-008 — DEPLOYMENT
10. VIG-DOC-009 — ROADMAP
11. VIG-DOC-010 — GLOSSARY
12. VIG-DOC-011 — DEVICE_SPEC
13. VIG-DOC-012 — DEVELOPMENT_GUIDE
14. VIG-DOC-013 — CONVENTIONS

---

# 4. Documentación Oficial

| ID | Documento | Descripción |
|----|-----------|-------------|
| VIG-DOC-000 | DOCUMENTATION_STANDARD | Estándar oficial para la documentación del proyecto. |
| VIG-DOC-001 | PRODUCT_VISION | Visión del producto, objetivos, alcance y propuesta de valor. |
| VIG-DOC-002 | ARCHITECTURE | Arquitectura general de la plataforma. |
| VIG-DOC-003 | DATA_MODEL | Modelo de datos y entidades del sistema. |
| VIG-DOC-004 | PROTOCOL_SPEC | Especificación del protocolo de comunicación con los dispositivos. |
| VIG-DOC-005 | RULE_ENGINE | Motor de reglas de negocio y generación de eventos. |
| VIG-DOC-006 | API_SPEC | Especificación de las APIs del sistema. |
| VIG-DOC-007 | SECURITY | Políticas y arquitectura de seguridad. |
| VIG-DOC-008 | DEPLOYMENT | Estrategia de despliegue e infraestructura. |
| VIG-DOC-009 | ROADMAP | Evolución planificada del producto. |
| VIG-DOC-010 | GLOSSARY | Glosario de términos técnicos y de negocio. |
| VIG-DOC-011 | DEVICE_SPEC | Especificación funcional del dispositivo IoT. |
| VIG-DOC-012 | DEVELOPMENT_GUIDE | Guía para el desarrollo del proyecto. |
| VIG-DOC-013 | CONVENTIONS | Convenciones generales del proyecto. |

---

# 5. Directorios

## templates/

Contiene las plantillas oficiales utilizadas para crear nuevos documentos.

## adr/

Contiene los **Architecture Decision Records (ADR)** del proyecto.

Cada ADR documenta una decisión arquitectónica importante y constituye un registro histórico.

## images/

Almacena diagramas, wireframes, mockups e imágenes utilizadas por la documentación, es decir, almacena todos los recursos gráficos utilizados por la documentación.

Estructura:

* architecture/
* database/
* ui/
* wireframes/

---

# 6. Versionado

Todos los documentos siguen **Semantic Versioning**.

Cada documento mantiene su propio versionado de manera independiente.

```text
MAJOR.MINOR.PATCH
```

Ejemplo:

| Prefijo | Descripción |
|---------|-------------|
| VIG-DOC | Documentación oficial |
| VIG-ADR | Architecture Decision Records |
| VIG-ARC | Diagramas de arquitectura |
| VIG-DB | Diagramas de base de datos |
| VIG-UI | Interfaces gráficas |
| VIG-WF | Wireframes |
| VIG-BR | Reglas de negocio |
| VIG-API | APIs |

---

# 7. Referencias

La documentación utiliza identificadores únicos para facilitar las referencias cruzadas.

Ejemplos:

```text
VIG-DOC-002
VIG-ADR-0001
VIG-ARC-001
VIG-DB-001
VIG-UI-001
VIG-WF-001
VIG-BR-001
VIG-API-001
```

---

# 8. Documento de referencia

Toda la documentación del proyecto deberá cumplir las reglas establecidas en:

**VIG-DOC-000 — DOCUMENTATION_STANDARD**

Este documento constituye el estándar oficial de documentación para Vigía Platform.

---

# 9. Licencia

Toda la documentación forma parte del proyecto Vigía Platform y deberá mantenerse alineada con la evolución funcional y técnica del software.

Cualquier modificación deberá seguir las normas definidas en VIG-DOC-000.