package main

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "given a positive number and positive n",
			args: args{
				x: 2,
				n: 10,
			},
			want: 1024,
		},
		{
			name: "given a positive number and negative n",
			args: args{
				x: 2,
				n: -2,
			},
			want: 0.25,
		},
		{
			name: "given a positive number and zero n",
			args: args{
				x: 2,
				n: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
