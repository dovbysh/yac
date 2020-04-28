package d

import (
	"bytes"
	"testing"
)

func TestGen(t *testing.T) {
	type args struct {
		len uint8
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "1",
			args: args{
				len: 1,
			},
			wantW: "()\n",
		},
		{
			name: "2",
			args: args{
				len: 2,
			},
			wantW: "(())\n()()\n",
		},
		{
			name: "3",
			args: args{
				len: 3,
			},
			wantW: "((()))\n(()())\n(())()\n()(())\n()()()\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Gen(tt.args.len, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Gen() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
