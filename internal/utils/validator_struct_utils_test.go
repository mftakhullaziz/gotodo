package utils

import "testing"

func TestValidateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Valid email",
			args:    args{email: "test@example.com"},
			wantErr: false,
		},
		{
			name:    "Invalid email",
			args:    args{email: "invalid_email"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEmail(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateIntValue(t *testing.T) {
	type args struct {
		val []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Valid values",
			args:    args{val: []int{1, 2, 3}},
			wantErr: false,
		},
		{
			name:    "Invalid values",
			args:    args{val: []int{1, -2, 3}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIntValue(tt.args.val...); (err != nil) != tt.wantErr {
				t.Errorf("ValidateIntValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
