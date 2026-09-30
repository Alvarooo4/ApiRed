# Historias de Usuario
Van con el tag 'user-stories' y cada una se irá resolviendo mediante issues en distintos milestones.

## [HU001] No sé si un campo tiene las condiciones climáticas adecuadas antes de trasladar mis colmenas
Como apicultor profesional, no tengo forma de comprobar si la flora de una ubicación específica ha sufrido por la sequía o las olas de calor sin desplazarme físicamente hasta allí, por lo que me arriesgo a hacer viajes a ciegas, perdiendo tiempo y dinero en combustible, o llevando a mis abejas a un lugar sin alimento.

**Jornada relacionada:** Jornada 1 (Manuel).

**Datos implicados y sus fuentes:**
*   **Clima local:** Archivos de datos estáticos (JSON o CSV) almacenados directamente en el repositorio del proyecto.
*   **Ubicaciones geográficas:** Zonas donde se encuentran los colmenares que la aplicación cruza con los archivos climáticos.
*   **Heurística de floración:** Reglas lógicas (necesidades de calor y agua de cada planta) configuradas y almacenadas en los propios ficheros (JSON/CSV) del sistema.

## [HU002] No sé cómo está afectando el cambio climático a los días útiles de floración en la comarca
Como técnica de cooperativa apícola, no tengo forma de comparar rápidamente los datos meteorológicos actuales con los históricos de distintos municipios, por lo que no puedo evaluar si los días de floración útil se están reduciendo ni asesorar a los socios sobre qué zonas deben evitar.

**Jornada relacionada:** Jornada 2 (Elena).

**Datos implicados y sus fuentes:**
*   **Clima histórico:** Datos históricos en formato JSON o CSV que tendremos en el repositorio de la aplicación.
*   **Días útiles de floración:** Dato que se calcula internamente por la aplicación, comparando los archivos climáticos históricos con la lógica configurada.