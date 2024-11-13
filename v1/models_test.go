package gollama

import (
	"testing"
)

func TestGollama_ListModels(t *testing.T) {
	tests := []struct {
		name    string
		c       *Gollama
		wantErr bool
	}{
		{
			name:    "ListModels",
			c:       New("llama3.2"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ListModels()
			if (err != nil) != tt.wantErr {
				t.Errorf("Gollama.ListModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("Gollama.ListModels() without models")
			}
		})
	}
}
