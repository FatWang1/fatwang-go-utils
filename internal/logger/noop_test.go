package logger

import (
	"reflect"
	"testing"
)

func TestNewNoop(t *testing.T) {
	tests := []struct {
		name string
		want *noop
	}{
		{
			name: "NewNoop",
			want: &noop{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noop_Info(t *testing.T) {
	type fields struct {
	}
	type args struct {
		in0 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "all is ok",
			fields: fields{},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &noop{}
			n.Info(tt.args.in0)
		})
	}
}
