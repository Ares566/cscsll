package cscsll

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		_storage interface{}
	}
	tests := []struct {
		name string
		args args
		want *CScsll
	}{
		{
			name: "Creating with Ints",
			args: args{_storage: make([]int, 0)},
			want: &CScsll{Storage: []int{}, Head: &node{Next: nil, Index: 0}},
		},
		{
			name: "Creating with Strings",
			args: args{_storage: make([]string, 1)},
			want: &CScsll{Storage: []string{""}, Head: &node{Next: nil, Index: 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args._storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCScsll_Add(t *testing.T) {
	type fields struct {
		Storage interface{}
		Head    *node
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CScsll{
				Storage: tt.fields.Storage,
				Head:    tt.fields.Head,
			}
			cs.Add(tt.args.val)
		})
	}
}
