# Modelo de Problema

Modelado del problema expuesto en [HU001](historias-usuario.md#hu001-no-sé-si-un-campo-tiene-las-condiciones-climáticas-adecuadas-antes-de-trasladar-mis-colmenas) atendiendo a la metodología _DDD_ y usando el lenguaje ubicuo definido en [docs/dominio.md](dominio.md).

## Términos a modelar:
Clima Local, Registro

## Términos modelados:
### Coordenadas
Objeto valor representado por dos valores numéricos flotantes, uno para la latitud y otro para la longitud.

### Flor
Objeto valor representado:
    - Nombre
    - Cantidad de agua óptima: En milímetros (cada milímetro equivale a 1 l/m²) representada por un flotante.
    - Temperatura óptima: En grados centígrados representado con un flotante.

### Situación Meteoreológica
Objeto valor representado por:
    - Coordenadas: del lugar al que alude
    - Fecha: del día al que alude, representada por una variable de tipo _date_.  
    - Agua: Cantidad de agua llovida ese día en esas coordenadas en mm.
    - Temperatura: 

## Quedan sin mencionar del lenguaje ubicuo:
Gastos - No relevante, el problema es saber si si el índice floral de una ubicación es óptimo, no se intenta reducir los gastos. (deducido del [journey](user-journeys.md#jornada-1-manuel-comprueba-la-viabilidad-del-terreno-antes-de-trasladar-el-colmenar), último párrafo, Manuel va probando ubicaciones hasta encontrar una con el índice floral óptimo, sin tener en cuenta el coste necesario para llegar allí).
Ruta - No relevante, comprobamos el índice floral de una ubicación, no cómo llegar hasta ella.
Desplazamiento - No relevante por los mismos motivos que ruta.
Apicultor - No relevante, no diferenciamos entre un apicultor y otro, pues ambos obtienen el mismo resultado de introducir los mismos datos en la aplicación. 
Dinero - No relevante, por los mismos motivos que gastos.
Trashumancia - No relevante, se considera que todos los desplazamientos realizados por el apicultor son de trashumancia.
Tiempo de desplazamiento - Incluido en minutos en el propio Desplazamiento. 
Viabilidad Floral - Puntuación calculable dada una ubicación, metereología local y heurística de floración, es por tanto un resultado, no un objeto en sí. 
Parcela - Como solo es representada por una coordenada, es sinónimo de coordenada.
Colmenar - Está incluido como número de colmenas de una parcela.
Colmena - Cada parcela indica cuántas colmenas tiene.
Abejas - No relevante para la resolución del problema.
Alimento - No relevante para la resolución del problema.
Heurística de Floración - Queda reflejada en los valores de temperatura y agua requeridos de cada flor.
Flora - Se utiliza el término para referirse al conjunto de flores, pero no operamos con dicho conjunto en ningún momento, solo con flores particulares.
Temperatura - Queda incluida en Situación Meteoreológica.
Fenómenos metereológicos - Quedan representados en situación metereológica
