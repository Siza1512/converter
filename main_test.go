package main

import "testing"

func Test_convertArshinToM(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 аршин равен 0 метрам",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "1 аршин равен 0.71 метрам",
			args: args{n: 1},
			want: 0.71,
		},
		{
			name: "2 аршина равно 1.42 метра",
			args: args{n: 2},
			want: 1.42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertArshinToM(tt.args.n); got != tt.want {
				t.Errorf("convertArshinToM() = %v, want %v", got, tt.want)
			}
		})
	}
}
