# ApiRed

## Problema a tratar
Hasta hace unos años, los apicultores podían anticipar en qué mes florecía cada planta siguiendo el calendario normal. Sin embargo, el cambio climático está provocando picos de calor inusuales en invierno o sequías prolongadas en primavera. Actualmente, la única forma de saber si el campo tiene floración disponible es desplazarse físicamente hasta el colmenar. Estos viajes a ciegas suponen un gasto innecesario de tiempo y combustible, y si se llega tarde, las colmenas pueden de quedarse sin alimento.

## Conocimiento personal
Al vivir en Chillón, un pueblo con un entorno rural, conozco este problema de primera mano porque mi padre es apicultor. Veo a menudo cómo se enfrenta a este problema debido a estos cambios térmicos, por lo que se ve obligado a asumir gastos innecesarios para ir a comprobar los colmenares físicamente. Además, puedo comprobar cómo esta mismo problema climático está afectando a otros sectores de la agricultura y la ganadería en mi entorno.

## ¿Cómo se obtienen los datos? 
Toda la información (temperaturas, precipitaciones y previsiones de las distintas zonas) se extrae conectando el sistema a portales de datos públicos y APIs de centros meteorológicos, como por ejemplo AEMET u OpenWeather.

## ¿Por qué requiere una lógica de negocio y no solo almacenamiento?
El sistema no se limita a guardar y mostrar el tiempo, sino que aplica una lógica de procesamiento sobre los datos obtenidos para:
- Calcular variables como el calor acumulado y la humedad de esa ubicación concreta.
- Analizar el calor acumulado por las plantas y la falta de humedad del terreno.
- Generar una estimación que prediga si la floración se va a adelantar o retrasar respecto al ciclo habitual.
- Validar y alertar sobre el riesgo de escasez de néctar para las colmenas en las semanas siguientes.

## Por qué requiere un despliegue en la nube
Porque el sistema procesa continuamente el clima de distintas zonas geográficas, y un apicultor necesita conocer el estado de floración desde su casa antes de iniciar un viaje largo con sus colmenas. Sin un entorno en la nube que cruce constantemente esa información y genere la predicción, el apicultor no podrá saber si debe realizar ese desplazamiento.

## Tarjetas de rol
A continuación se muestran las imágenes de los roles asignados para el desarrollo del proyecto:

**Perspectiva del Cliente (El problema):**
![Tarjeta del cliente](media/Cliente.jpeg)

**Perspectiva del Desarrollador (La solución):**
![Tarjeta del desarrollador](media/Desarrollador.jpeg)

## Configuración
La configuración del entorno local se encuentra detallada paso a paso en el siguiente enlace:

[Ver detalles y capturas de la configuración](doc/configuracion.md)