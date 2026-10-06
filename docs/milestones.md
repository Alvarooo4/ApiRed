## Milestone 0: Modelo inicial del dominio

### Objetivo
Definir en código los conceptos clave del entorno apícola que sacamos de la [HU001] y [HU002], creando las estructuras de datos iniciales.

### Qué se entrega
- Los conceptos representados directamente en los ficheros de código del proyecto.
- Las decisiones de diseño, que se irán debatiendo y registrando paso a paso en los issues del repositorio.
- Los commits con los avances de código, enlazados siempre a la issue en la que se está trabajando.

### Validación
- El hito se da por cerrado cuando se puede seguir el hilo completo del trabajo en GitHub. Para ello, hay que comprobar que el código subido va atado a un commit, que ese commit resuelve una issue concreta, y que esa issue viene directamente de los problemas de la [HU001] o [HU002].

## Milestone 1: Lógica de negocio y pruebas automáticas

### Objetivo
Desarrollar el motor principal del proyecto (la lógica que calcula el [Índice de Viabilidad Floral](docs/conceptos.md) que necesitan la [HU001] y [HU002]) usando el hito anterior, y montar los tests automáticos.

### Qué se entrega
- El código con la lógica central ya implementada y lista para funcionar.
- Una batería de tests automáticos integrada en el repositorio para asegurar que los cálculos hacen lo que deben.

### Validación
- El hito se da por cerrado con pruebas automáticas. Para ello, los tests deben ejecutarse solos, pasar en verde y estar cada uno asociado al issue que comprueban.