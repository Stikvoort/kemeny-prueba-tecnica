# Arquitectura y Decisiones

## 1. Code Review — Resumen Ejecutivo

¿Cuáles son los 2-3 problemas más graves y por qué los priorizaste?

Los 3 problemas criticas para mi eran los que afectan seguridad, especialmente para un MVP temas de seguridad necesitan su propio atencion. Sin embargo es muy facil de enfocarse en los funciones del application para tener algo que sirve lo mas rapido posible. Problemas con seguridad son un peligro real y existencial para un producto web.

## 2. Integración de IA

- ¿Qué modelo/provider elegiste y por qué?
- ¿Cómo diseñaste el prompt? ¿Qué iteraciones hiciste?
- ¿Cómo manejas los casos donde el LLM falla, tarda, o retorna datos inválidos?
- ¿Cómo manejarías el costo en producción? (caching, rate limiting, batch processing)
- Si tuvieras que clasificar 10,000 tareas existentes, ¿cómo lo harías?

- Para la integracion de IA, elige OpenAI porque tengo mas conociemento con ellos. 
- Por falta de tiempo no ha diseñado mas que un prompt, que probablemente puede ser mucho mejor que es ahora. 
- En todo en el implementacion hay un gran dependencia en el client de OpenAI para resolver la mayoria de problemas mentionadas.
- Manejar el costo depende tambien en el uso, puede ser suficiente para tener las limites en la cuenta de openAI. Monitorear para adaptar los necesidades es recomendable en caso de un MVP.
- Por grandes cantidades de tareas que necesitan ser clasificadas un job que comenza a en la noche, que puede procesar todo en forma batch probablemente es lo mejor. Pero mucho depende de necesidad del cliente y sus ideas por el producto.


## 3. Docker y Orquestación

¿Qué decisiones tomaste? (multi-stage builds, networking, volumes, health checks, etc.)

De verdad, no tome mucho deciciones por falta de tiempo, y no tener un setup de docker activo en my computador. IA era una guia bueno para tener algo en este caso.
El setup de ahora a menos falta un network dedicado al application, para aislar los servicos del ambiente. Un otro aumentacion es el inclusion de un health check, pero para hacer esto en mejor forma, tambien unos cambios son necesarios en el codigo del backend y frontend, por ejemplo falta un check por la conexion con el base de datos.

## 4. Arquitectura para Producción

Si este proyecto fuera a producción con 10,000 usuarios concurrentes:

- ¿Qué cambiarías en la arquitectura?
- ¿Qué servicios agregarías? (cache, queue, CDN, etc.)
- ¿Cómo manejarías el deployment?
- Incluye un diagrama si te parece útil (ASCII, Mermaid, o imagen)

Con 10.000 usuarios, la cantidad de tasks va a aumentar significante. Para manejar los repuestos mas comon, la introducion de Redis para caching puede ser necesario. Para el frontend, cloudflare es un buen opcion para acelar y manejar el delivery de sitio. Con rate limiting en los respuestos al LLM y al backend puede tener un flujo stabil por los usuarios independiente del numero de usarios. Probablemente, multiples instancias del los servicios son necesarios para servir. Usar kubernetes para el deployment de los servicios genera mas control sobre los instancias, y recuperacion si una falla.


## 5. Trade-offs

¿Qué decisiones tomaste donde había más de una opción válida?

Mis decisiones fueron influido mas por el poco tiempo que tuve, que por un analysis fundamental de los opciones. 

## 6. Qué Mejorarías con Más Tiempo

Sé específico y prioriza.

Eso depende mucho de los stakeholders, sus planes son lo mas importante para mejorar el app. Unas temas importantes de pensar son: los roles de usuarios, la implementacion de LLM para clasificar, seperacion de estylo y el react code, tailwind puede ser un buen opcion.
Priorizacion depende entonces mucho en los funciones requridos por las stakeholders. Para un MVP, este 'prototypo' necesita convertirse en un applicacion apta par producion. Si este es importante, lo primero es segurar un sierto nivel de seguridad, y funcionalidad que este alienada con las esperanzas del usarios.

## 7. Uso de IA (como herramienta de desarrollo)

¿Usaste IA para desarrollar? ¿Para qué? ¿Modificaste lo que te sugirió?

Por el poco tiempo que tuve, y mi desconociamento del idioma go, IA fue un gran ayuda durante este processo. Para hacer review, para desarrollar la integracion con IA y para hacer el containerization del app.
