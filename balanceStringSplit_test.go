package balance_string

import "testing"

func Test_balanceStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				s: "RLRRLLRLRL",
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				s: "RLLLLRRRLR",
			},
			want: 3,
		},
		{
			name: "Example3",
			args: args{
				s: "LLLLRRRR",
			},
			want: 1,
		},
		{
			name: "Example4",
			args: args{
				s: "RLRRRLLRLL",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balanceStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
