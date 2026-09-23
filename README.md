# ApiRed

## Problema a tratar
Hasta hace unos años, los apicultores podían anticipar en qué mes florecía cada planta siguiendo el calendario normal. Sin embargo, el cambio climático está provocando olas de calor inusuales en invierno o sequías en primavera. Actualmente, la única forma de saber si el campo tiene floración disponible es desplazarse físicamente hasta el colmenar. Estos viajes a ciegas suponen un gasto innecesario de tiempo y combustible, y si se llega tarde, las colmenas pueden quedarse sin alimento.

Una colmena está situada en un campo, por lo que está asociada a la flora que rodea esa ubicación. El objetivo del apicultor es que la floración produzca néctar para sus abejas, por lo cual la floración está vinculada a la producción del néctar. Cuando se produce la floración, las plantas segregan néctar que es recogido por las propias abejas. Este néctar les sirve de alimento en ese mismo momento y, tras procesarlo en la colmena durante varios días, se convierte en la miel.

Cuando hay un periodo largo sin floración útil alrededor, el único remedio para el apicultor es trasladar físicamente las colmenas a otra zona (trashumancia) o utilizar alimentación artificial.

## Conocimiento personal
Al vivir en Chillón, un pueblo con un entorno rural, conozco este problema de primera mano porque mi padre es apicultor. Veo a menudo cómo se enfrenta a este problema debido a estos cambios térmicos, por lo que se ve obligado a asumir gastos innecesarios para ir a comprobar los colmenares físicamente. Además, puedo comprobar cómo este mismo problema climático está afectando a otros sectores de la agricultura y la ganadería en mi zona.

## ¿Cómo se obtienen los datos? 
Para evaluar el estado del terreno vamos a utilizar las siguientes fuentes de datos:

1. Lo que aporta el usuario: El apicultor registra su colmenar poniendo las coordenadas y eligiendo la flor principal (por ejemplo, romero o jara). Esta ubicación coincidirá con la flora, ya que las colmenas se encuentran en espacios rodeados de campo.

2. El clima: La aplicación contará con archivos (CSV/JSON) con los registros meteorológicos recientes e históricos de temperatura y lluvia de esas zonas. Estos datos tendrán una precisión espacial a nivel municipal y una precisión temporal diaria.

Los datos históricos nos servirán para ver cómo afectaron las condiciones climáticas a floraciones pasadas para compararlas con los datos actuales y, de esa forma, establecer la predicción actual.

3. Heurística: Las reglas de cuánta lluvia y qué límites de temperatura necesita cada planta para que su flor genere néctar las he extraído hablando directamente con un apicultor profesional (mi padre). 

Además, para asegurar los números exactos de cada flor, se contrastarán sus consejos buscando en manuales y guías de agricultura. Estos umbrales estarán integrados manualmente dentro de la base de datos de la aplicación, por lo que no supondrán un subproblema.

## ¿Por qué requiere una lógica de negocio y no solo almacenamiento?
El sistema no se limita a guardar y mostrar el tiempo, sino que aplica una lógica de procesamiento sobre los datos obtenidos para:
- Calcular variables como el calor acumulado y la lluvia caída, sacando los datos de nuestros archivos para esa ubicación concreta.
- Analizar el calor acumulado y la falta de humedad cruzando esa información con el tipo de flor elegida y las reglas del apicultor.
- Generar una estimación que nos diga si las condiciones del terreno son buenas para que la flor dé néctar.
- Validar y alertar sobre el riesgo de escasez de néctar en las colmenas, para que el apicultor sepa si de verdad merece la pena hacer el viaje.

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