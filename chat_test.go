package gollama

import (
	"encoding/json"
	"testing"
)

func TestGollama_Chat(t *testing.T) {
	type args struct {
		Prompt  string
		Options interface{}
	}
	type outs struct {
		wantContent  *ChatOuput
		wantToolJson string
	}
	tests := []struct {
		name    string
		c       *Gollama
		args    args
		want    *outs
		wantErr bool
	}{
		{
			name:    "Vision",
			c:       New("llama3.2-vision"),
			args:    args{Prompt: "what is on the road?", Options: []string{"./test/road.png"}},
			want:    &outs{wantContent: &ChatOuput{Content: "There is a llama on the road."}},
			wantErr: false,
		},
		{
			name:    "Math",
			c:       New("llama3.2"),
			args:    args{Prompt: "what is 2 + 2? only answer in number"},
			want:    &outs{wantContent: &ChatOuput{Content: "4"}},
			wantErr: false,
		},
		{
			name: "Tool",
			c:    New("llama3.2"),
			args: args{Prompt: "what is the weather in New York?", Options: []Tool{
				{Type: "function", Function: ToolFunction{
					Name:        "get_current_weather",
					Description: "Get the current weather in a specific city",
					Parameters: ToolParameters{
						Type: "object",
						Properties: map[string]ToolProperty{
							"city": {
								Type:        "string",
								Description: "The name of the city",
							},
						},
						Required: []string{"city"},
					}},
				}},
			},
			want:    &outs{wantContent: &ChatOuput{Content: ""}, wantToolJson: `[{"function":{"name":"get_current_weather","arguments":{"city":"New York"}}}]`},
			wantErr: false,
		},
		{
			name:    "Invalid model",
			c:       New("invalid"),
			args:    args{Prompt: "hello"},
			want:    &outs{wantContent: &ChatOuput{Content: ""}},
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

			if got != nil && tt.want != nil &&
				got.Content != tt.want.wantContent.Content {
				t.Errorf("Gollama.Chat() = %v, want %v", got, tt.want)
			}

			if got != nil && tt.want != nil && tt.want.wantToolJson != "" {
				toolJson, _ := json.Marshal(got.ToolCalls)
				if string(toolJson) != tt.want.wantToolJson {
					t.Errorf("Gollama.Chat() tool calls = %v, want %v", string(toolJson), tt.want.wantToolJson)
				}
			}
		})
	}
}
