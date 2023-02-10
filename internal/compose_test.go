package internal

import (
	"os"
	"reflect"
	"testing"

	"github.com/compose-spec/compose-go/types"
)

func TestReadCompose(t *testing.T) {
	content, _ := os.ReadFile("../../data/docker-compose.yaml")
	srvs := &types.Services{}
	type args struct {
		filename string
		content  []byte
	}
	tests := []struct {
		name string
		args args
		want *types.Services
	}{
		{name: "t0", args: args{filename: "docker-compose.yaml", content: content}, want: srvs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadCompose(tt.args.filename, tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCompose() = %v, want %v", got, tt.want)
			}
		})
	}
}
