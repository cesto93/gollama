package gollama

import (
	"testing"
)

func TestGollama_Chat(t *testing.T) {
	type args struct {
		in ChatInput
	}
	tests := []struct {
		name    string
		c       *Gollama
		args    args
		want    *ChatResponse
		wantErr bool
	}{
		{
			name:    "Vision",
			c:       New("llama3.2-vision"),
			args:    args{in: ChatInput{Prompt: "what is on the road?", VisionImages: []string{"./test/road.png"}}},
			want:    &ChatResponse{Content: "There is a llama on the road."},
			wantErr: false,
		},
		{
			name:    "Math",
			c:       New("llama3.2"),
			args:    args{in: ChatInput{Prompt: "what is 2 + 2? only answer in number"}},
			want:    &ChatResponse{Content: "4"},
			wantErr: false,
		},
		{
			name:    "Invalid model",
			c:       New("invalid"),
			args:    args{in: ChatInput{Prompt: "hello"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Chat(tt.args.in)
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
