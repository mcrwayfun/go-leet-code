package _74_search_a_2D_matrix

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
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 3,
			},
			want: true,
		},
		{
			name: "test_case_2",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 13,
			},
			want: false,
		},
		{
			name: "test_case_3",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}},
				target: 3,
			},
			want: true,
		},
		{
			name: "test_case_4",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}},
				target: 13,
			},
			want: false,
		},
		{
			name: "test_case_5",
			args: args{
				matrix: [][]int{},
				target: 0,
			},
			want: false,
		},
		{
			name: "test_case_6",
			args: args{
				matrix: [][]int{{1}, {2}, {3}},
				target: 3,
			},
			want: true,
		},
		{
			name: "test_case_7",
			args: args{
				matrix: [][]int{{}},
				target: 1,
			},
			want: false,
		},
		{
			name: "test_case_7",
			args: args{
				matrix: [][]int{{1}, {2}, {3}},
				target: 4,
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

func Test_searchMatrix2(t *testing.T) {
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
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 3,
			},
			want: true,
		},
		{
			name: "test_case_2",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}},
				target: 13,
			},
			want: false,
		},
		{
			name: "test_case_3",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}},
				target: 3,
			},
			want: true,
		},
		{
			name: "test_case_4",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}},
				target: 13,
			},
			want: false,
		},
		{
			name: "test_case_5",
			args: args{
				matrix: [][]int{},
				target: 0,
			},
			want: false,
		},
		{
			name: "test_case_6",
			args: args{
				matrix: [][]int{{1}, {2}, {3}},
				target: 3,
			},
			want: true,
		},
		{
			name: "test_case_7",
			args: args{
				matrix: [][]int{{}},
				target: 1,
			},
			want: false,
		},
		{
			name: "test_case_7",
			args: args{
				matrix: [][]int{{1}, {2}, {3}},
				target: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix2(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}