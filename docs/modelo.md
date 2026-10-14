# Modelo de Problema

Modelado del problema expuesto en [HU001](historias-usuario.md#hu001-no-sé-si-un-campo-tiene-las-condiciones-climáticas-adecuadas-antes-de-trasladar-mis-colmenas) atendiendo a la metodología _DDD_ y usando el lenguaje ubicuo definido en [docs/dominio.md](dominio.md).

## Términos a modelar:
Ruta, Parcela, Coordenadas, Desplazamiento, Clima Local, Registro, Datos Metereológicos, Fenómenos metereológicos, Variables Meteorológicas -> solo se considera la temperatura -> Temperatura, Flora, Flor, Heurística de Floración.

## Quedan sin mencionar del lenguaje ubicuo:
Gastos - No relevante, el problema es saber si si el índice floral de una ubicación es óptimo, no se intenta reducir los gastos. (deducido del [journey](user-journeys.md#jornada-1-manuel-comprueba-la-viabilidad-del-terreno-antes-de-trasladar-el-colmenar), último párrafo, Manuel va probando ubicaciones hasta encontrar una con el índice floral óptimo, sin tener en cuenta el coste necesario para llegar allí).
Apicultor - No relevante, no diferenciamos entre un apicultor y otro, pues ambos obtienen el mismo resultado de introducir los mismos datos en la aplicación. 
Dinero - No relevante, por los mismos motivos que gastos.
Trashumancia - No relevante, se considera que todos los desplazamientos realizados por el apicultor son de trashumancia.
Tiempo de desplazamiento - Incluido en minutos en el propio Desplazamiento. 
Viabilidad Floral - Puntuación calculable dada una ubicación, metereología local y heurística de floración, es por tanto un resultado, no un objeto en sí. 
Colmenar - Está incluido como número de colmenas de una parcela.
Colmena - Cada parcela indica cuántas colmenas tiene.
Abejas - No relevante para la resolución del problema.
Alimento - No relevante para la resolución del problema.
