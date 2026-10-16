# Diagrama del modelo

Diagrama de las estructuras de datos implementadas en `internal/modelo`, generado a partir del código. La clasificación en entidad, objeto valor y registro es la definida en [modelo.md](modelo.md).

```mermaid
classDiagram
    class Clima {
        <<Entidad>>
        -coordenada Coordenada
        -registroMeteorologico Map~time.Time, Meteorologia~
        +NuevoClima(coord Coordenada) Clima
        +AgregarMeteorologia(meteorologia Meteorologia) error
        +ModificarMeteorologia(fecha, agua, temperaturaMin, temperaturaMax) error
    }

    class Meteorologia {
        <<Objeto valor>>
        -fecha time.Time
        -aguaLlovida float64
        -temperaturaMin float64
        -temperaturaMax float64
        +NuevaMeteorologia(fecha, agua, tempMin, tempMax) Meteorologia, error
        +NormalizarFecha(fecha time.Time) time.Time
    }

    class Coordenada {
        <<Objeto valor>>
        -latitudY float64
        -longitudX float64
        +NuevaCoordenada(lY, lX float64) Coordenada, error
    }

    class Flora {
        <<Registro>>
        -flores Map~string, Flor~
        +NuevaFlora() Flora
        +AgregarFlor(flor Flor) error
        +ModificarFlor(nombre, agua, temperatura) error
    }

    class Flor {
        <<Objeto valor>>
        -nombre string
        -aguaIdeal float64
        -temperaturaIdeal float64
        +NuevaFlor(nombre string, agua, temperatura float64) Flor, error
    }

    Clima *-- Coordenada : identifica
    Clima *-- "0..*" Meteorologia : indexada por fecha
    Flora *-- "0..*" Flor : indexada por nombre
```

## Invariantes comprobadas en los constructores

| Tipo | Comprobación |
|---|---|
| `Coordenada` | `latitudY` en `[-90, 90]`; `longitudX` en `[-180, 180]` |
| `Flor` | `nombre` no vacío; `aguaIdeal > 0`; `temperaturaIdeal` en `[-50, 60]` |
| `Meteorologia` | `aguaLlovida >= 0`; `temperaturaMin <= temperaturaMax`; `temperaturaMin >= -273.15` |
| `Clima` | no se admiten dos meteorologías con la misma fecha |
| `Flora` | no se admiten dos flores con el mismo nombre |
