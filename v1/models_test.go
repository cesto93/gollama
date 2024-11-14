package gollama

import (
	"log"
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

func TestGollama_HasModel(t *testing.T) {
	type args struct {
		model string
	}
	tests := []struct {
		name    string
		c       *Gollama
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "HasModel",
			c:       New("llama3.2"),
			args:    args{model: "llama3.2"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.HasModel(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gollama.HasModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Gollama.HasModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGollama_PullModel(t *testing.T) {
	tests := []struct {
		name    string
		c       *Gollama
		model   string
		wantErr bool
	}{
		{
			name:    "PullModel",
			c:       New("llama3.2"),
			model:   "opencoder:1.5b",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.PullModel(tt.model); (err != nil) != tt.wantErr {
				t.Errorf("Gollama.PullModel() error = %v, wantErr %v", err, tt.wantErr)
			}

			log.Fatalln(tt.c.HasModel(tt.model))
		})
	}
}
