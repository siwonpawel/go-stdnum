package pl

import "testing"

func Test_validateRegon_9digit(t *testing.T) {
	type args struct {
	}
	tests := []string{
		"123456785",
		"292804060",
		"234191600",
		"175256672",
		"014148890",
		"917583463",
		"639145979",
		"590869583",
		"977339644",
		"554366290",
		"158365673",
		"192828833",
		"715027385",
		"198575557",
		"377427137",
		"554754572",
		"056957556",
		"011373534",
		"733779786",
		"893800074",
		"254561575",
		"351121520",
		"879121591",
		"814671336",
		"012239192",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			got, err := validateRegon(tt)
			if (err != nil) != false {
				t.Errorf("validateRegon() error = %v", err)
				return
			}
			if got != true {
				t.Errorf("validateRegon() got = %v, want %v", got, true)
			}
		})
	}
}

func Test_validateRegon_14digit(t *testing.T) {
	type args struct {
	}
	tests := []string{
		"67366764535896",
		"83559229617313",
		"23528305951638",
		"33116940691096",
		"35760040490290",
		"85066453258709",
		"79292862266767",
		"73878597502200",
		"97298381452805",
		"41131710917198",
		"01076794604438",
		"37593312674563",
		"13335909277021",
		"05221203252346",
		"59357910171293",
		"05906916425572",
		"63703526451674",
		"73114773497055",
		"03566178633032",
		"75876561074053",
		"39884417533204",
		"09046345407887",
		"09046345407887",
		"95781339295560",
		"95972130051440",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			got, err := validateRegon(tt)
			if (err != nil) != false {
				t.Errorf("validateRegon() error = %v", err)
				return
			}
			if got != true {
				t.Errorf("validateRegon() got = %v, want %v", got, true)
			}
		})
	}
}
