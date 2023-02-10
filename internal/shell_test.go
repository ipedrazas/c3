package internal

import (
	"testing"
)

func TestExecute(t *testing.T) {
	script := "/usr/local/bin/kompose"
	cmd := []string{"kompose", "version"}
	type args struct {
		script  string
		command []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "t01", args: args{script: script, command: cmd}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Execute(tt.args.script, tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
