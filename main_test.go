package main

import "testing"

func Test_convertArshinToM(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "0 аршин равен 0 метрам",
			args:    args{n: 0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "1 аршин равен 0.71 метрам",
			args:    args{n: 1},
			want:    0.71,
			wantErr: false,
		},
		{
			name:    "2 аршина равно 1.42 метра",
			args:    args{n: 2},
			want:    1.42,
			wantErr: false,
		},
		{
			name:    "-1 аршин нельзя конвертировать",
			args:    args{n: -1},
			want:    0,
			wantErr: true,
		},
		{
			name:    "-2 аршин нельзя конвертировать",
			args:    args{n: -2},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertArshinToM(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertArshinToM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == true && err != negativeNumberNotSupport {
				t.Errorf("convertArshinToM() err = %v, want err %v", err, negativeNumberNotSupport)
				return
			}
			if got != tt.want {
				t.Errorf("convertArshinToM() got = %v, want %v", got, tt.want)
			}
		})
	}
}
