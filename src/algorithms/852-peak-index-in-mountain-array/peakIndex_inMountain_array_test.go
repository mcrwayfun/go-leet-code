package _52_peak_index_in_mountain_array

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-case1",
			args: args{[]int{0, 1, 0}},
			want: 1,
		},
		{
			name: "test-case2",
			args: args{[]int{0, 2, 1, 0}},
			want: 1,
		},
		{
			name: "test-case2",
			args: args{[]int{0, 2, 3, 4}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.A); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_peakIndexInMountainArray2(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-case1",
			args: args{[]int{0, 1, 0}},
			want: 1,
		},
		{
			name: "test-case2",
			args: args{[]int{0, 2, 1, 0}},
			want: 1,
		},
		{
			name: "test-case2",
			args: args{[]int{0, 2, 3, 4}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray2(tt.args.A); got != tt.want {
				t.Errorf("peakIndexInMountainArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}