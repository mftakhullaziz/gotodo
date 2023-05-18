package utils

import "testing"

func TestComparedPassword(t *testing.T) {
	type args struct {
		hashedPwd     string
		plainPassword []byte
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ComparedPassword(tt.args.hashedPwd, tt.args.plainPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComparedPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ComparedPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashPassword(tt.args.password); got != tt.want {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPasswordAndSalt(t *testing.T) {
	type args struct {
		pwd []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashPasswordAndSalt(tt.args.pwd); got != tt.want {
				t.Errorf("HashPasswordAndSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}
