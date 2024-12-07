package gollama

import (
	"testing"
)

func TestGollama_Chat(t *testing.T) {
	type args struct {
		Prompt  string
		Options interface{}
	}
	tests := []struct {
		name    string
		c       *Gollama
		args    args
		want    *ChatOuput
		wantErr bool
	}{
		{
			name:    "Vision",
			c:       New("llama3.2-vision"),
			args:    args{Prompt: "what is on the road?", Options: []string{"./test/road.png"}},
			want:    &ChatOuput{Content: "There is a llama on the road."},
			wantErr: false,
		},
		{
			name:    "Math",
			c:       New("llama3.2"),
			args:    args{Prompt: "what is 2 + 2? only answer in number"},
			want:    &ChatOuput{Content: "4"},
			wantErr: false,
		},
		{
			name:    "Invalid model",
			c:       New("invalid"),
			args:    args{Prompt: "hello"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Verbose = true
			got, err := tt.c.Chat(tt.args.Prompt, tt.args.Options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gollama.Chat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Content != tt.want.Content {
				t.Errorf("Gollama.Chat() = %v, want %v", got, tt.want)
			}
		})
	}
}
