package _40_search_a_2D_matrixII

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_case_1",
			args: args{
				matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
				target: 5,
			},
			want: true,
		},
		{
			name: "test_case_2",
			args: args{
				matrix: [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
				target: 20,
			},
			want: false,
		},
		{
			name: "test_case_3",
			args: args{
				matrix: [][]int{{1, 4, 5}},
				target: 1,
			},
			want: true,
		},
		{
			name: "test_case_4",
			args: args{
				matrix: [][]int{{1, 4, 5}},
				target: 6,
			},
			want: false,
		},
		{
			name: "test_case_5",
			args: args{
				matrix: [][]int{{1}, {2}, {3}},
				target: 1,
			},
			want: true,
		},
		{
			name: "test_case_6",
			args: args{
				matrix: [][]int{{1}, {2}, {3}},
				target: 6,
			},
			want: false,
		},
		{
			name: "test_case_7",
			args: args{
				matrix: [][]int{{}},
				target: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
