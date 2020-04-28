package funcs

import "testing"

func Test_countStrange(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	s1 := "ab"
	s2 := "aabbddcfg"
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "zz",
			args: args{
				s1: s1,
				s2: s2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStrange(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("countStrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
