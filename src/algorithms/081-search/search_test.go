package _81_search

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
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
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "test_case_2",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			name: "test_case_3",
			args: args{
				nums:   []int{1, 1, 1, 1, 0, 1, 1},
				target: 0,
			},
			want: true,
		},
		{
			name: "test_case_4",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1},
				target: 1,
			},
			want: true,
		},
		{
			name: "test_case_5",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1},
				target: 0,
			},
			want: false,
		},
		{
			name: "test_case_6",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: true,
		},
		{
			name: "test_case_7",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
