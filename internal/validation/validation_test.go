package validation

import "testing"

func TestValidateINN(t *testing.T) {
	type args struct {
		inn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid inn", args: args{"7802565953"}, wantErr: false},
		{name: "invalid inn", args: args{"1234567890"}, wantErr: true},
		{name: "less count digits", args: args{"780256595"}, wantErr: true},
		{name: "more count digits", args: args{"78025659530"}, wantErr: true},
		{name: "text", args: args{"foo"}, wantErr: true},
		{name: "mix lead valid inn text", args: args{"7802565953foo"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateINN(tt.args.inn); (err != nil) != tt.wantErr {
				t.Errorf("ValidateINN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
