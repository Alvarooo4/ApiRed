## Milestone 0: Análisis inicial del problema

### Objetivo
Empezar a darle forma a la [HU001] usando la metodología DDD. En este paso nos centramos en llevar el problema al código, sin entrar en la lógica de negocio.

### Validación
Para dar este hito por superado, se va a comprobar que:
- Cada issue que abramos sirve para atacar un problema concreto derivado de la [HU001].
- Todos los problemas contenidos en la [HU001] tienen que estar representados con su propio issue.
- Cualquier trozo de código que subamos está enlazado directamente al issue que intenta resolver.
- Cada commit hace referencia o cierra un solo issue.

## Milestone 1: Lógica de negocio y pruebas automáticas

### Objetivo
Aprovechando la base que dejamos montada en el hito cero, ahora sí nos metemos de lleno a programar la lógica principal que soluciona la [HU001]. Lo haremos aplicando principios SOLID y creando las pruebas que verifican dicha lógica[cite: 8].

### Validación
Aquí la validación la dejamos puramente en manos de la máquina. El hito se considerará cerrado de forma objetiva cuando los tests que comprueban la [HU001] salten solos al subir los cambios, pasen todos en verde y nos confirmen que el código hace exactamente lo que le toca.