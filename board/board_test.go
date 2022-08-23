package board

import (
	"reflect"
	"testing"
)

func Test_sumRowLeft(t *testing.T) {
	type args struct {
		row [4]int
	}
	tests := []struct {
		name string
		args args
		want [4]int
	}{
		{
			name: "{0, 0, 0, 0} -> {0, 0, 0, 0}",
			args: args{row: [4]int{0, 0, 0, 0}},
			want: [4]int{0, 0, 0, 0},
		},
		{
			name: "{2, 2, 0, 0} -> {4, 0, 0, 0}",
			args: args{row: [4]int{2, 2, 0, 0}},
			want: [4]int{4, 0, 0, 0},
		},
		{
			name: "{2, 4, 2, 0} -> {2, 4, 2, 0}",
			args: args{row: [4]int{2, 4, 2, 0}},
			want: [4]int{2, 4, 2, 0},
		},
		{
			name: "{2, 2, 2, 2} -> {4, 4, 0, 0}",
			args: args{row: [4]int{2, 2, 2, 2}},
			want: [4]int{4, 4, 0, 0},
		},
		{
			name: "{0, 2, 0, 2} -> {4, 0, 0, 0}",
			args: args{row: [4]int{0, 2, 0, 2}},
			want: [4]int{4, 0, 0, 0},
		}, {
			name: "{2, 0, 2, 0} -> {4, 0, 0, 0}",
			args: args{row: [4]int{0, 2, 0, 2}},
			want: [4]int{4, 0, 0, 0},
		},
		{
			name: "{4, 4, 2, 2} -> {8, 4, 0, 0}",
			args: args{row: [4]int{4, 4, 2, 2}},
			want: [4]int{8, 4, 0, 0},
		},
		{
			name: "{2, 0, 0, 2} -> {4, 0, 0, 0}",
			args: args{row: [4]int{2, 0, 0, 2}},
			want: [4]int{4, 0, 0, 0},
		},
		{
			name: "{2, 2, 2, 0} -> {4, 2, 0, 0}",
			args: args{row: [4]int{2, 2, 2, 0}},
			want: [4]int{4, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumRowLeft(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumRowLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
