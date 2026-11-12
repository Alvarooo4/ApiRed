# Elección del gestor de tareas para ApiRed

## Contexto

El desarrollo del proyecto ApiRed requiere la automatización de operaciones repetitivas fundamentales, tales como la instalación de dependencias, la validación de la sintaxis del código y la ejecución de la batería de pruebas. 

Dado que el ecosistema de Go ya proporciona una gestión de dependencias nativa y robusta a través de los módulos (`go.mod`), este análisis se centra exclusivamente en seleccionar la herramienta óptima para orquestar y automatizar las tareas del repositorio.

## Alternativas consideradas

Tras investigar las opciones más estandarizadas y recomendadas en foros de desarrollo y comunidades de Go (como Reddit y StackOverflow), se han evaluado las siguientes cuatro herramientas:

- GNU Make
- Just
- Mage
- Task

## Criterios de evaluación

Para cumplir con la exigencia de fundamentar las decisiones arquitectónicas en datos objetivos y no en popularidad, estética o familiaridad previa, se han establecido los siguientes criterios empíricos enfocados en minimizar la deuda técnica de ApiRed:

- **Impacto en el rendimiento:** Tiempo real que añade la herramienta a la ejecución del comando subyacente.
- **Deuda técnica (Huella en disco):** Peso en el sistema del ejecutable requerido.
- **Complejidad de adopción:** Necesidad de instalar binarios o entornos adicionales pesados.
- **Alineación con el proyecto:** Suficiencia para gestionar los hitos actuales de la aplicación.

## Mediciones

Para obtener datos empíricos, se ha instrumentado la orden básica de validación de sintaxis en Go:

`go build ./...`

Las métricas se han extraído utilizando la herramienta de benchmarking `hyperfine`. Para garantizar la fiabilidad estadística y descartar fluctuaciones del sistema operativo, se han configurado 5 ciclos de calentamiento previo (`--warmup 5`) y un mínimo de 50 ejecuciones consecutivas (`--min-runs 50`) por cada gestor.

## Comparativa de resultados

Los resultados arrojados por las pruebas en el entorno local (Windows / Git Bash) se exponen en la siguiente tabla para facilitar la toma de decisiones:

| Herramienta | Tiempo de la tarea `check` | Sobrecarga temporal | Peso del ejecutable (Deuda técnica) |
| :--- | :--- | :--- | :--- |
| **Make** | **365.3 ms** | **1.4x más lento** | **172 KB** |
| **Just** | 398.2 ms | 1.6x más lento | 5.7 MB |
| **Mage** | 1107.6 ms | 4.2x más lento | 7.2 MB |
| **Task** | 2329.7 ms | 9.6x más lento | 72 MB |

### Evidencias de ejecución
*(Imágenes capturadas durante el benchmarking)*

![Medición de Make](media/Make.png)
![Medición de Just](media/Just.png)
![Medición de Mage](media/Mage.png)
![Medición de Task](media/Task.png)

## Elección final

A la vista de la evidencia empírica, **se establece GNU Make como la herramienta oficial para la automatización de tareas en ApiRed**.

La decisión se justifica principalmente en la drástica reducción de deuda técnica en comparación con las alternativas. Make resuelve las necesidades del proyecto añadiendo una penalización de rendimiento mínima (1.4x sobre el comando base) e introduciendo un peso casi nulo en el sistema (172 KB). 

Aunque Just presenta tiempos de ejecución muy competitivos (398.2 ms) y una sintaxis moderna, exige introducir un ejecutable de 5.7 MB. Por su parte, herramientas escritas en Go como Mage y Task quedan totalmente descartadas para este contexto, al multiplicar los tiempos de compilación por factores inasumibles (hasta 9.6x en el caso de Task) y arrastrar binarios excesivamente pesados (hasta 72 MB) para la simple orquestación de comandos.