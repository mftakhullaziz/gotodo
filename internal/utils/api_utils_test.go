package utils

import "testing"

func TestListEndpoints(t *testing.T) {
	type args struct {
		endpoints []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListEndpoints(tt.args.endpoints)
		})
	}
}
