package gollama

import (
	"testing"
)

func TestGollama_Embedding(t *testing.T) {
	type args struct {
		in ChatInput
	}
	tests := []struct {
		name    string
		c       *Gollama
		args    args
		wantLen int
		wantErr bool
	}{
		{
			name:    "Embedding",
			c:       New("llama3.2"),
			args:    args{in: ChatInput{Prompt: "hello"}},
			wantLen: 3072,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Embedding(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gollama.Embedding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Gollama.Embedding() = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}
