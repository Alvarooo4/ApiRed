# Elección del gestor de tareas para ApiRed

## Contexto

El proyecto necesita automatizar operaciones repetitivas, como instalar las dependencias, comprobar la sintaxis del código o lanzar los tests, para no tener que escribir cada orden a mano. Go ya gestiona las dependencias mediante módulos (`go.mod`), así que lo que hay que elegir es la herramienta que lanza esas tareas.

## Alternativas consideradas

Se han comparado cuatro gestores de tareas que se usan habitualmente en proyectos Go:

- GNU Make
- Just
- Mage
- Task

## Criterios

La elección se basa en dos medidas objetivas, tomadas en el mismo equipo (Windows con Git Bash):

- **Velocidad de ejecución:** cuánto tiempo añade el gestor a la orden que ejecuta por dentro.
- **Velocidad de instalación:** cuánto tarda en instalarse la herramienta desde cero con su método habitual.

## Método de medición

En los cuatro gestores, la tarea `check` ejecuta la misma orden, que comprueba la sintaxis de los ficheros Go sin compilarlos:

`gofmt -e . > /dev/null`

**Ejecución.** Se ha usado `hyperfine` para comparar cada tarea con la orden que ejecuta por dentro:

`hyperfine -N --warmup 5 --min-runs 50 'gofmt -e .' '<gestor> check'`

- Se hacen 5 ejecuciones de calentamiento y al menos 50 ejecuciones medidas.
- La opción `-N` lanza las órdenes sin una shell intermedia, así que los tiempos no incluyen el arranque de la shell.
- La orden de referencia no lleva `> /dev/null` porque `hyperfine` ya descarta la salida por defecto.
- La diferencia entre las dos medias es el tiempo que añade el gestor.

**Instalación.** Se ha medido con `time`, partiendo de que la herramienta no estaba instalada. Solo se tiene en cuenta el tiempo `real`, porque `user` y `sys` no incluyen el tiempo de los procesos de Windows que lanzan las instalaciones.

## Resultados