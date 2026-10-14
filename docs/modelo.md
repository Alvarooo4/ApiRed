# Modelo de Problema

Modelado del problema expuesto en [HU001](historias-usuario.md#hu001-no-sé-si-un-campo-tiene-las-condiciones-climáticas-adecuadas-antes-de-trasladar-mis-colmenas) atendiendo a la metodología _DDD_ y usando el lenguaje ubicuo definido en [docs/dominio.md](dominio.md).

## Términos a modelar:
Ruta, Ubicación, Coordenadas, Desplazamiento, Clima Local, Registro, Datos Metereológicos, Fenómenos metereológicos, Variables Meteorológicas -> solo se considera la temperatura -> Temperatura, Flora, Flor, Heurística de Floración.

## Quedan sin mencionar del lenguaje ubicuo:
Gastos - No relevante, el problema es saber si si el índice floral de una ubicación es óptimo, no se intenta reducir los gastos. (deducido del [journey](user-journeys.md#jornada-1-manuel-comprueba-la-viabilidad-del-terreno-antes-de-trasladar-el-colmenar), último párrafo, Manuel va probando ubicaciones hasta encontrar una con el índice floral óptimo, sin tener en cuenta el coste necesario para llegar allí).
Apicultor, 
Dinero, 
Trashumancia - No relevante, se considera que todos los desplazamientos realizados por el apicultor son de trashumancia.
Tiempo de desplazamiento - Incluido en minutos en el propio Desplazamiento. 
Viabilidad Floral - Puntuación calculable dada una ubicación, metereología local y heurística de floración, es por tanto un resultado, no un objeto en sí. 
Colmenar - Solo debería incluirse si influye en el proceso de Cargar, refs #12
Colmena - No relevante, el problema es saber si una ubicación posee una viabilidad floral adecuada.
, Abejas, Alimento.
