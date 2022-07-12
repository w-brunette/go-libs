package moeda

import (
	"fmt"
	"math"
)

const Exp int = 2

type Moeda interface {
	Display() string
	Somar(m Moeda) Moeda
	SomarInt64(value int64) Moeda

	Dividir(m Moeda) Moeda
	DividirInt64(value int64) Moeda

	Subtrair(m Moeda) Moeda
	SubtrairInt64(value int64) Moeda

	Multiplicar(m Moeda) Moeda
	MultiplicarInt64(value int64) Moeda

	Float64() float64
	Int64() int64
	LessThan(m Moeda) bool
	LessThanInt64(value int64) bool
	GreaterThan(m Moeda) bool
	GreaterThanInt64(value int64) bool
}

type moeda struct {
	valor int64
}

func (m *moeda) Display() string {
	return fmt.Sprintf("R$ %.2f", float64(m.valor)/math.Pow10(2))
}

func FromInt64(valor int64) Moeda {
	return &moeda{
		valor: valor,
	}
}

func FromFloat64(valor float64) Moeda {
	r := valor*math.Pow10(Exp) + math.Copysign(0.5, valor)
	return &moeda{
		valor: int64(r),
	}
}

func (m *moeda) Float64() float64 {
	return float64(m.valor) / math.Pow10(Exp)
}

func (m *moeda) Int64() int64 {
	return m.valor
}

func (m *moeda) Somar(other Moeda) Moeda {
	return &moeda{
		valor: m.valor + other.Int64(),
	}
}
func (m *moeda) SomarInt64(value int64) Moeda {
	return &moeda{
		valor: m.valor + value,
	}
}
func (m *moeda) Subtrair(other Moeda) Moeda {
	return &moeda{
		valor: m.valor - other.Int64(),
	}
}
func (m *moeda) SubtrairInt64(value int64) Moeda {
	return &moeda{
		valor: m.valor - value,
	}
}

func (m *moeda) Multiplicar(other Moeda) Moeda {
	return &moeda{
		valor: m.valor * other.Int64(),
	}
}
func (m *moeda) MultiplicarInt64(value int64) Moeda {
	return &moeda{
		valor: m.valor * value,
	}
}

func (m *moeda) Dividir(other Moeda) Moeda {
	return &moeda{
		valor: m.valor / other.Int64(),
	}
}
func (m *moeda) DividirInt64(value int64) Moeda {
	return &moeda{
		valor: m.valor / value,
	}
}

func (m *moeda) LessThan(other Moeda) bool {
	return m.valor < other.Int64()
}
func (m *moeda) LessThanInt64(other int64) bool {
	return m.valor < other
}

func (m *moeda) GreaterThan(other Moeda) bool {
	return m.valor > other.Int64()
}

func (m *moeda) GreaterThanInt64(other int64) bool {
	return m.valor > other
}

func (m *moeda) Equals(other Moeda) bool {
	return m.valor == other.Int64()
}

func (m *moeda) IsZero() bool {
	return m.valor == 0.0
}

func Zero() Moeda {
	return &moeda{}
}
