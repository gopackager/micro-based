package helper

import (
	"testing"
)

func TestEnv(t *testing.T) {
	type args struct {
		v        string
		fallback []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "should be success first params",
			args: args{
				v:        "LOGNAME",
				fallback: nil,
			},
			want: "macbook",
		},
		{
			name: "should be success second params on first params is not set",
			args: args{
				v:        "MY_NAME",
				fallback: []string{"POPO"},
			},
			want: "POPO",
		},
		{
			name: "should be failed second params",
			args: args{
				v:        "MY_NAME",
				fallback: []string{"POPO"},
			},
			want: "POPO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Env(tt.args.v, tt.args.fallback...); got != tt.want {
				t.Errorf("Env() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "should be success",
			want: uuid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uuid; got != tt.want {
				t.Errorf("GenerateUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
