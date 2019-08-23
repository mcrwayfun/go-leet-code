package _89_my_sqrt

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "test-case-1",
		//	args: args{4},
		//	want: 2,
		//},
		//{
		//	name: "test-case-2",
		//	args: args{8},
		//	want: 2,
		//},
		{
			name: "test-case-3",
			args: args{2147483647},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
