package gollama

import (
	"testing"
)

func TestGollama_Vision(t *testing.T) {
	type args struct {
		prompt string
		images []string
	}
	tests := []struct {
		name    string
		c       *Gollama
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Vision",
			c:       New("llama3.2-vision"),
			args:    args{prompt: "what is on the road?", images: []string{"./test/road.png"}},
			want:    "There is a llama on the road.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Vision(tt.args.prompt, tt.args.images)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gollama.Vision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Gollama.Vision() = %v, want %v", got, tt.want)
			}
		})
	}
}
