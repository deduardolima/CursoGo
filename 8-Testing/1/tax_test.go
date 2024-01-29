package tax

import "testing"

func TestCalculateTax( t * testing.T){
	amout := 500.0
	expected := 5.0

	result := CalculateTax(amout);

		if result != expected{
				t.Errorf("Expected %f but got %f", expected, result)
		}

}

func TestCalculateTaxBatch( t * testing.T){
	type calcTax struct {
		amout, expected float64
	}
	table := []calcTax {
		{500.0, 5.0},
		{1000.0, 10.0},
		{0.0, 0.0},
		{1300.0, 10.0},
	}

	for _, item := range table{
		result := CalculateTax(item.amout)
		if result != item.expected{
				t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}

}
func BenchmarkCalculateTax(b *testing.B) {
    for i := 0; i < b.N; i++ {
        CalculateTax(500.0)
    }
}

func FuzzCalculateTax(f *testing.F){
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0 }
	for _, amout := range seed {
		f.Add(amout)
	}
	f.Fuzz(func(t * testing.T, amout float64){
		result := CalculateTax(amout)
		if amout <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amout > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})

}


//go tool cover -html=coverage.out
//go test -coverprofile=coverage.out