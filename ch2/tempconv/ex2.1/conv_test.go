package tempconv

import (
	"math"
	"testing"
)

func TestCToF(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Fahrenheit
	}{
		// TODO: Add test cases.
		{"Freezing point of water", args{0}, 32},
		{"Water boils", args{100}, 212},
		{"Hot bath", args{40}, 104},
		{"Very cold day", args{-17.78}, 0},
		{"Extremely cold day", args{-40}, -40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			esp := 0.01 // acceptable floating point error
			if got := CToF(tt.args.c); math.Abs(float64(got-tt.want)) > esp {
				t.Errorf("CToF(%v) = %v, want %v", tt.args.c, got, tt.want)
			}
		})
	}
}

func TestFToC(t *testing.T) {
	type args struct {
		f Fahrenheit
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		// TODO: Add test cases.
		{"Freezing point of water", args{32}, 0},
		{"Water boils", args{212}, 100},
		{"Hot bath", args{104}, 40},
		{"Very cold day", args{0}, -17.78},
		{"Extremely cold day", args{-40}, -40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			esp := 0.01 // acceptable floating point error
			if got := FToC(tt.args.f); math.Abs(float64(got-tt.want)) > esp {
				t.Errorf("FToC(%v) = %v, want %v", tt.args.f, got, tt.want)
			}
		})
	}
}