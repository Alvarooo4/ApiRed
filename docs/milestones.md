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
**Definición del Índice de Viabilidad Floral (IVF):**
Es el cálculo que determina si una zona concreta va a tener alimento suficiente para las abejas. El sistema observa los litros de lluvia o la temperatura de los archivos JSON/CSV y los cruza con las reglas lógicas que hemos desarrollado con la ayuda de un profesiional de la apicultura. Por tanto, evalúa si el agua y el calor se han dado en el momento y cantidad exacta que necesita cada tipo de flor para producir néctar.

### Qué se entrega
- El código empaquetado como un módulo, para que cualquier persona pueda descargarlo, instalarlo y usarlo fácilmente.
- Una serie de tests automáticos integrados en el código para validar que todos los cálculos y procesos funcionan de la manera adecuada.
- El hito estará superado cuando los tests se ejecuten automáticamente al subir el código a GitHub y nos confirmen que todo está correctamente.