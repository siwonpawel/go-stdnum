package pl

import (
	"fmt"
	"github.com/siwonpawel/go-stdnum/stdnum/validation"
	"reflect"
	"testing"
)

func Test_cleanse(t *testing.T) {
	type returnVals struct {
		cleaned  string
		warnings []string
	}
	tests := []struct {
		name   string
		number string
		want   returnVals
	}{
		{
			"Should trim spaces",
			"  7008045141  ",
			returnVals{
				cleaned:  "7008045141",
				warnings: []string{validation.MissingCountryPrefix},
			},
		},
		{
			"Should cleanse spaces",
			"70 08 04 51 41",
			returnVals{
				cleaned:  "7008045141",
				warnings: []string{validation.MissingCountryPrefix},
			},
		},
		{
			"Should remove country identifier",
			"PL70 08 04 51 41",
			returnVals{
				cleaned:  "7008045141",
				warnings: []string{},
			},
		},
		{
			"Should remove non-alphabetic or non-numeric chars",
			"70 08`~!@#$%^&*()+=,<./;'\\[]{}:¬∨¦¦|>\"|?>< 04 51 41",
			returnVals{
				cleaned:  "7008045141",
				warnings: []string{validation.MissingCountryPrefix},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, warnings := validation.Cleanse(tt.number, country); got != tt.want.cleaned || !reflect.DeepEqual(warnings, tt.want.warnings) {
				t.Errorf("cleanse() = [ %#v %#v ], want %#v", got, warnings, tt.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	rc := validation.NewResultCreator("PL", "NIP")

	tests := []struct {
		name   string
		number string
		want   *validation.Result
	}{
		{
			"Should be invalid on empty number",
			"",
			rc.Fail(
				"",
				[]string{validation.MissingCountryPrefix},
				validation.InvalidLength,
				validation.DebugInfo{},
			),
		},
		{
			"Should be invalid",
			"PL1234567890",
			rc.Fail(
				"PL1234567890",
				[]string{},
				validation.InvalidNumber,
				validation.DebugInfo{
					CleanedInput: "1234567890",
				},
			),
		},
		{
			"Should be valid",
			"PL0123456789",
			rc.Ok(
				"PL0123456789",
				[]string{},
				validation.DebugInfo{
					CleanedInput: "0123456789",
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateNIP(tt.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateNIP(%#v) =\n\t %#v\nwant %#v", tt.number, got, tt.want)
			}
		})
	}
}

func TestValidate_correctRecords(t *testing.T) {
	tests := []struct {
		number string
		want   bool
	}{
		{
			"0013599074",
			true,
		},
		{
			"6316798344",
			true,
		},
		{
			"0997308934",
			true,
		},
		{
			"0568679920",
			true,
		},
		{
			"8936667241",
			true,
		},
		{
			"6181170802",
			true,
		},
		{
			"1604761526",
			true,
		},
		{
			"8520413089",
			true,
		},
		{
			"1622892368",
			true,
		},
		{
			"8420373410",
			true,
		},
		{
			"2107612480",
			true,
		},
		{
			"4869292276",
			true,
		},
		{
			"7461345752",
			true,
		},
		{
			"4559870537",
			true,
		},
		{
			"6423269304",
			true,
		},
		{
			"4879335219",
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Should be valid %.2v", i), func(t *testing.T) {
			if got, err := validateNIP(tt.number); got != tt.want || err != nil {
				t.Errorf("ValidateNIP(%#v) =\n%#v, %v want\n%#v, %v", tt.number, got, err, tt.want, err)
			}
		})
	}
}
