package utils

import (
	"context"
	"github.com/openzipkin/zipkin-go"
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
