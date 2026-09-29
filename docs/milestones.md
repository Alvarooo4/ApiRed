# Milestones
Definen qué quiero entregar, no cómo entregarlo.

## Milestone 0: Modelo inicial de dominio

### Objetivo
Construir la base estructural del problema apícola apoyándose en la HU001.

### Qué se entrega
- Los elementos esenciales del entorno (parcelas, clima, plantas) plasmados en código limpio.
- Únicamente estructuras de datos, dejando fuera por ahora cualquier cálculo o regla de negocio.
- Un archivo en la carpeta docs/ que explique el diseño (DDD), aclarando cuáles son las entidades, los objetos valor y cómo se relacionan.
- Una pequeña justificación sobre las decisiones tomadas durante este diseño.

## Milestone 1: Lógica de negocio y cálculo de viabilidad

### Objetivo
Programar el motor de la aplicación que evaluará el Índice de Viabilidad Floral.

### Qué se entrega
- Funciones que crucen la meteorología de la zona con los requisitos de las flores.
- Un conjunto de tests automáticos integrados en el código.
- El hito se dará por superado si las pruebas confirman que el sistema evalúa la viabilidad correctamente.