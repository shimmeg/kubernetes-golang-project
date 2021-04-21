package validation

import (
	"reflect"
	"testing"
)

func TestValidateIsin(t *testing.T) {
	type args struct {
		isin string
	}
	tests := []struct {
		name string
		args args
		want Status
	}{
		{name: "", args: args{isin: "IE00B4X9L533"}, want: OK},
		{name: "", args: args{isin: "DE00B4X9L53⅐"}, want: NotValid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIsin(tt.args.isin); !(got.Status == tt.want) {
				t.Errorf("ValidateIsin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateNumberEquiualent(t *testing.T) {
	type args struct {
		symbol rune
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateNumberEquiualent(tt.args.symbol); got != tt.want {
				t.Errorf("calculateNumberEquiualent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{numbers: []int{3, 0, 2, 8, 0, 3, 7, 8, 3, 3, 1, 0, 0}}, want: 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSum(tt.args.numbers); got != tt.want {
				t.Errorf("calculateSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToNumbers(t *testing.T) {
	type args struct {
		isin string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{isin: "US037833100"}, want: []int{3, 0, 2, 8, 0, 3, 7, 8, 3, 3, 1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToNumbers(getRuneSlice(tt.args.isin)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickValidations(t *testing.T) {
	type args struct {
		isin string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "When provided valid ISIN shouldn't be errors", args: args{
			isin: "IE00B4X9L533",
		}, want: 0},
		{name: "When invalid ISIN - result should contain errors", args: args{
			isin: "I00B4X9L533",
		}, want: 2}, // length error + alphabetical at the beginning
		{name: "When invalid ISIN - result should contain errors", args: args{
			isin: "IE00B4X9L53A",
		}, want: 1}, // last symbol is not the number
		{name: "When invalid ISIN - result should contain errors", args: args{
			isin: "IE00B4X9L535",
		}, want: 0},
		{name: "When invalid ISIN - result should contain errors", args: args{
			isin: "DE00B4X9L53⅐",
		}, want: 1}, //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickValidations(getRuneSlice(tt.args.isin), tt.args.isin); !(len(got) == tt.want) {
				t.Errorf("quickValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getRuneSlice(isin string) *[]rune {
	tmp := []rune(isin)
	return &tmp
}

func Test_roundUpValue(t *testing.T) {
	type args struct {
		sum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "45 should be round up to 50", args: args{45}, want: 50},
		{name: "30 should be round up to 30", args: args{30}, want: 30},
		{name: "59 should be round up to 60", args: args{59}, want: 60},
		{name: "61 should be round up to 70", args: args{61}, want: 70},
		{name: "72 should be round up to 80", args: args{72}, want: 80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundUpValue(tt.args.sum); got != tt.want {
				t.Errorf("roundUpValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
