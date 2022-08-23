package board

import (
	"reflect"
	"testing"
)

func Test_processRowLeft(t *testing.T) {
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
			if got := processRowLeft(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processRowLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_SumLeft(t *testing.T) {
	type fields struct {
		board [4][4]int
	}
	tests := []struct {
		name   string
		fields fields
		want   [4][4]int
	}{
		{
			name: "SumLeft",
			fields: fields{
				board: [4][4]int{
					{2, 0, 2, 0},
					{2, 4, 2, 0},
					{4, 2, 2, 0},
					{2, 0, 0, 2},
				},
			},
			want: [4][4]int{
				{4, 0, 0, 0},
				{2, 4, 2, 0},
				{4, 4, 0, 0},
				{4, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{
				board: tt.fields.board,
			}
			b.SumLeft()
			if !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("board.SumLeft() = %v, want %v", b.board, tt.want)
			}
		})
	}
}

func Test_board_SumRight(t *testing.T) {
	type fields struct {
		board [4][4]int
	}
	tests := []struct {
		name   string
		fields fields
		want   [4][4]int
	}{
		{
			name: "SumLeft",
			fields: fields{
				board: [4][4]int{
					{2, 0, 2, 0},
					{2, 4, 2, 0},
					{4, 2, 2, 0},
					{2, 0, 0, 2},
				},
			},
			want: [4][4]int{
				{0, 0, 0, 4},
				{0, 2, 4, 2},
				{0, 0, 4, 4},
				{0, 0, 0, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{
				board: tt.fields.board,
			}
			b.SumRight()
			if !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("board.SumRight() = %v, want %v", b.board, tt.want)
			}
		})
	}
}
