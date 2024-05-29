package utils

import (
	"reflect"
	"testing"
)

func TestGetUuid(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "ok",
			want: GetUuid(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUuid(); len(got) != len(tt.want) {
				t.Errorf("GetUuid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTernaryOperator(t *testing.T) {
	type args[T any] struct {
		expr     bool
		trueVal  T
		falseVal T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[bool]{
		{
			name: "true",
			args: args[bool]{
				expr:     1 == 1,
				trueVal:  true,
				falseVal: false,
			},
			want: true,
		},
		{
			name: "false",
			args: args[bool]{
				expr:     1 != 1,
				trueVal:  true,
				falseVal: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TernaryOperator(tt.args.expr, tt.args.trueVal, tt.args.falseVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TernaryOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
