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
			name: "1 аршин равен 0.71 метрам",
			args: args{n: 1},
			want: 0.71,
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