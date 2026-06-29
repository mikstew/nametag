package platform

import "testing"

func TestHandoffPID(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "no flag",
			args: []string{"nametag"},
			want: 0,
		},
		{
			name: "valid pid",
			args: []string{"nametag", "--wait-for-pid=4242"},
			want: 4242,
		},
		{
			name: "invalid pid",
			args: []string{"--wait-for-pid=abc"},
			want: 0,
		},
		{
			name: "zero pid",
			args: []string{"--wait-for-pid=0"},
			want: 0,
		},
		{
			name: "other args preserved",
			args: []string{"--wait-for-pid=99", "--other"},
			want: 99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandoffPID(tt.args); got != tt.want {
				t.Fatalf("HandoffPID() = %d, want %d", got, tt.want)
			}
		})
	}
}
