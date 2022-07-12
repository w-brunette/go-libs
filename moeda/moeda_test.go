package moeda

import (
	"testing"
)

func TestMoedaFromFloat64(t *testing.T) {
	moeda := FromFloat64(1000.30)
	if moeda.Int64() != 100030 {
		t.Fail()
	}
}
func TestMoedaToFloat64(t *testing.T) {
	moeda := FromInt64(100030)
	if moeda.Float64() != 1000.30 {
		t.Fail()
	}
}
