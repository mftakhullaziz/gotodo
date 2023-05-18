package utils

import "testing"

func TestHasValue(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValue(tt.args.s); got != tt.want {
				t.Errorf("HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasValueSlice(t *testing.T) {
	type args struct {
		handler interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValueSlice(tt.args.handler); got != tt.want {
				t.Errorf("HasValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
