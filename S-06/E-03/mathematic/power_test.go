package mathematic

import "testing"

func TestPowerH(t *testing.T) {
	type test struct {
		Base           float64
		Power          float64
		ExpectedResult float64
	}
	testCases := []test{
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 2, Power: 5, ExpectedResult: 32},
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 0, ExpectedResult: 1},
	}

	for _, c := range testCases {
		res := PowerH(c.Base, c.Power)
		if res != c.ExpectedResult {
			t.Errorf("result : %f ||| Expected Result : %f", res, c.ExpectedResult)
		}
	}
}

func TestPowerMath(t *testing.T) {
	type test struct {
		Base           float64
		Power          float64
		ExpectedResult float64
	}
	testCases := []test{
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 2, Power: 5, ExpectedResult: 32},
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 0, ExpectedResult: 1},
	}

	for _, c := range testCases {
		res := PowerMath(c.Base, c.Power)
		if res != c.ExpectedResult {
			t.Errorf("result : %f ||| Expected Result : %f", res, c.ExpectedResult)
		}
	}
}

var result float64

func BenchmarkPowerH(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res := PowerH(2, 4)
		result = res
	}
}

func BenchmarkPowerMath(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res := PowerMath(2, 4)
		result = res
	}
}

func BenchmarkPowerSlc (b *testing.B) {
	b.ReportAllocs()
	for i := 0 ; i < b.N ; i++ {
		PowerSlc(2 , 4)
	}
}
