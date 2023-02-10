package internal

import (
	"testing"
)

func TestExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "t01", args: args{path: "/tmp/doesnotexists"}, want: false},
		{name: "t01", args: args{path: "/tmp"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.path); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBinPath(t *testing.T) {
	type args struct {
		tool string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "t0", args: args{tool: "docker"}, want: "/opt/homebrew/bin/docker"},
		{name: "t1", args: args{tool: "kompose"}, want: "/usr/local/bin/kompose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBinPath(tt.args.tool); got != tt.want {
				t.Errorf("GetBinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
