package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalSequence(t *testing.T) {
	type args struct {
		mtx [][]int
		ua  []int
	}

	mtx1 := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}
	
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "mtx 5 verticals 100%",
			args: args{
				mtx: mtx1,
				ua:  []int{4, 1, 0, 2},
			},
			want: 100,
		},
		{
			name: "mtx 5 verticals 0%",
			args: args{
				mtx: mtx1,
				ua:  []int{},
			},
			want: 0,
		},
		{
			name: "mtx 5 verticals 50%",
			args: args{
				mtx: mtx1,
				ua:  []int{4, 1, 0},
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvalSequence(tt.args.mtx, tt.args.ua)
			assert.Equal(t, tt.want, got)
		})
	}
}
