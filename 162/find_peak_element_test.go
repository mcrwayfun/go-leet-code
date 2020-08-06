package _62

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-case1",
			args: args{[]int{1, 2, 3, 1}},
			want: 2,
		},
		{
			name: "test-case2",
			args: args{[]int{1, 2, 1, 3, 5, 6, 4}},
			want: 5,
		},
		{
			name: "test-case3",
			args: args{[]int{0, 1}},
			want: 1,
		},
		{
			name: "test-case3",
			args: args{[]int{0, 1, 0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
