package pl

import (
	"github.com/siwonpawel/go-stdnum/stdnum/validation"
	"reflect"
	"testing"
)

func Test_validatePESEL_validOnly(t *testing.T) {
	tests := []string{
		"44042203232",
		"13102504611",
		"08031105628",
		"03062702686",
		"03092303439",
		"78080205600",
		"47082804667",
		"14030200332",
		"88080400356",
		"42121105600",
		"54111402457",
		"16021206200",
		"43020502822",
		"20022408493",
		"60080503178",
		"00280201436",
		"99060403308",
		"16110100859",
		"03111300104",
		"96052102548",
		"54111603834",
		"17092006421",
		"74021105875",
		"60022605858",
		"50022304710",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			got, err := validatePESEL(tt)
			if err != nil {
				t.Errorf("validatePESEL() error = %v, wantErr %v", err, nil)
				return
			}
			if got != true {
				t.Errorf("validatePESEL() got = %v, want %v", got, true)
			}
		})
	}
}

func Test_validatePESEL_invalidOnly(t *testing.T) {
	tests := []string{
		"43042203232",
		"14102504611",
		"05031105628",
		"05062702686",
		"05092303439",
		"75080205600",
		"45082804667",
		"15030200332",
		"85080400356",
		"45121105600",
		"55111402457",
		"15021206200",
		"45020502822",
		"25022408493",
		"65080503178",
		"05280201436",
		"95060403308",
		"15110100859",
		"05111300104",
		"95052102548",
		"55111603834",
		"15092006421",
		"75021105875",
		"65022605858",
		"55022304710",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			got, err := validatePESEL(tt)
			if err != nil {
				t.Errorf("validatePESEL() error = %v, wantErr %v", err, nil)
				return
			}
			if got != false {
				t.Errorf("validatePESEL() got = %v, want %v", got, true)
			}
		})
	}
}

func TestValidatePESEL(t *testing.T) {
	rc := validation.NewResultCreator(country, peselIdentifierName)

	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want *validation.Result
	}{
		{
			"Should validate prefixed PESEL",
			args{number: "PL20022408493"},
			rc.Ok("PL20022408493", []string{}, validation.DebugInfo{CleanedInput: "20022408493"}),
		},
		{
			"Should validate PESEL with no country prefix",
			args{number: "20022408493"},
			rc.Ok("20022408493", []string{"Missing country identifier at beginning."}, validation.DebugInfo{CleanedInput: "20022408493"}),
		},
		{
			"Should validate invalid PESEL",
			args{number: "PL21022408493"},
			rc.Fail("PL21022408493", []string{}, validation.InvalidNumber, validation.DebugInfo{CleanedInput: "21022408493"}),
		},
		{
			"Should validate invalid PESEL with no country prefix",
			args{number: "21022408493"},
			rc.Fail("21022408493", []string{"Missing country identifier at beginning."}, validation.InvalidNumber, validation.DebugInfo{CleanedInput: "21022408493"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePESEL(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidatePESEL() = \n%v,\nwant\n%v", got, tt.want)
			}
		})
	}
}
