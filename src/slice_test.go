package slice

import (
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	a := []int{1, 2}
	b := Copy(a)
	if b[0] != 1 || b[1] != 2 {
		t.Fatal("error")
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		a []int
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{a: []int{1, 2, 3}, i: 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.a, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		a    []int
		keep func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{a: []int{1, 2, 3}, keep: func(i int) bool { return i > 1 }}, []int{2, 3}},
		{"1", args{a: []int{1, 2, 3}, keep: func(i int) bool { return i%3 == 0 }}, []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.a, tt.args.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	type args struct {
		a *[]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{&[]int{1, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pop(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
			if len(*tt.args.a) != 2 {
				t.Errorf("Pop() did not popped anything")
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		a *[]int
		i int
		x int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{&[]int{0, 1, 3}, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.a, tt.args.i, tt.args.x)
		})
		if (*tt.args.a)[2] != 2 && (*tt.args.a)[3] != 3 {
			t.Errorf("No")
		}
	}
}

func TestInsertList(t *testing.T) {
	type args struct {
		a *[]int
		i int
		x []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{&[]int{0, 1, 4}, 2, []int{2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.a, tt.args.i, tt.args.x)
		})
		if (*tt.args.a)[2] != 2 && (*tt.args.a)[3] != 3 && (*tt.args.a)[4] != 4 {
			t.Errorf("No")
		}
	}
}
