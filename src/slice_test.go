package slice

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		g    *GoSlice
		keep func(v any) bool
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{"1", args{&GoSlice{Elements: []any{1, 2, 3}}, func(v any) bool { return v == 2 }}, []any{2}},
		{"1", args{&GoSlice{Elements: []any{1, 2, 3, 6}}, func(v any) bool { return v.(int)%3 == 0 }}, []any{3, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.g, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
