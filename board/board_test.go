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
			b := &board{}
			if got := b.processRowLeft(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processRowLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_sumLeft(t *testing.T) {
	type fields struct {
		board [4][4]int
	}
	tests := []struct {
		name   string
		fields fields
		want   [4][4]int
	}{
		{
			name: "sumLeft",
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
			b.sumLeft()
			if !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("board.sumLeft() = %v, want %v", b.board, tt.want)
			}
		})
	}
}

func Test_board_sumRight(t *testing.T) {
	type fields struct {
		board [4][4]int
	}
	tests := []struct {
		name   string
		fields fields
		want   [4][4]int
	}{
		{
			name: "sumRight",
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
			b.sumRight()
			if !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("board.sumRight() = %v, want %v", b.board, tt.want)
			}
		})
	}
}

func Test_board_sumUp(t *testing.T) {
	type fields struct {
		board [4][4]int
	}
	tests := []struct {
		name   string
		fields fields
		want   [4][4]int
	}{
		{
			name: "sumUp",
			fields: fields{
				board: [4][4]int{
					{2, 0, 2, 2},
					{2, 4, 2, 0},
					{4, 2, 2, 0},
					{2, 0, 0, 2},
				},
			},
			want: [4][4]int{
				{4, 4, 4, 4},
				{4, 2, 2, 0},
				{2, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{
				board: tt.fields.board,
			}
			b.sumUp()
			if !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("board.sumUp() = %v, want %v", b.board, tt.want)
			}
		})
	}
}

func Test_board_sumDown(t *testing.T) {
	type fields struct {
		board [4][4]int
	}
	tests := []struct {
		name   string
		fields fields
		want   [4][4]int
	}{
		{
			name: "sumDown",
			fields: fields{
				board: [4][4]int{
					{2, 0, 2, 2},
					{2, 4, 2, 0},
					{4, 2, 2, 0},
					{2, 0, 0, 2},
				},
			},
			want: [4][4]int{
				{0, 0, 0, 0},
				{4, 0, 0, 0},
				{4, 4, 2, 0},
				{2, 2, 4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{
				board: tt.fields.board,
			}
			b.sumDown()
			if !reflect.DeepEqual(b.board, tt.want) {
				t.Errorf("board.sumDown() = %v, want %v", b.board, tt.want)
			}
		})
	}
}

func Test_board_IsGameOver(t *testing.T) {
	type fields struct {
		board [4][4]int
		score int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "IsGameOver - full board stuck",
			fields: fields{
				board: [4][4]int{
					{32, 256, 2, 16},
					{8, 32, 256, 4},
					{2, 16, 8, 2},
					{4, 2, 16, 4},
				},
			},
			want: true,
		},
		{
			name: "IsGameOver - full board not stuck",
			fields: fields{
				board: [4][4]int{
					{8, 2, 16, 32},
					{4, 32, 4, 16},
					{2, 8, 2, 16},
					{2, 4, 2, 4},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{
				board: tt.fields.board,
				score: tt.fields.score,
			}
			if got := b.IsGameOver(); got != tt.want {
				t.Errorf("IsGameOver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStuck(t *testing.T) {
	type args struct {
		row [4]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isStuck - can add 2",
			args: args{
				row: [4]int{32, 16, 16, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStuck(tt.args.row); got != tt.want {
				t.Errorf("isStuck() = %v, want %v", got, tt.want)
			}
		})
	}
}
