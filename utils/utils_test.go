package utils

import (
	"context"
	"reflect"
	"testing"

	"github.com/openzipkin/zipkin-go"
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

func TestContextCopy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	ctx := context.Background()
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "all is ok",
			args: args{
				ctx: ctx,
			},
			want: zipkin.NewContext(ctx, zipkin.SpanOrNoopFromContext(ctx)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContextCopy(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContextCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomIntList(t *testing.T) {
	type args struct {
		size  int
		edge1 int
		edge2 int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "all is ok",
			args: args{
				size:  3,
				edge1: 3,
				edge2: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GenerateRandomIntList(tt.args.size, tt.args.edge1, tt.args.edge2))
		})
	}
}

func TestMax(t *testing.T) {
	type args[T NUMBER] struct {
		m T
		n T
	}
	type testCase[T NUMBER] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "all is ok",
			args: args[int]{
				m: 1,
				n: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args[T NUMBER] struct {
		m T
		n T
	}
	type testCase[T NUMBER] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "all is ok",
			args: args[int]{
				m: 1,
				n: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
