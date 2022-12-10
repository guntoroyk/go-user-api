package converter

import "testing"

func TestToInt(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "string to int",
			args: args{
				s: "1",
			},
			want: 1,
		},
		{
			name: "int64 to int",
			args: args{
				s: int64(1),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.s); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
