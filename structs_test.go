package gollama

import (
	"reflect"
	"testing"
)

func TestStructToStructuredFormat(t *testing.T) {
	type args struct {
		s interface{}
	}

	type myStruct struct {
		Content string `json:"content" required:"true"`
		Ignored string `ignored:"true"`
		Value   bool   `json:"value" description:"test value"`
		List    []int  `json:"list" description:"test list"`
	}

	tests := []struct {
		name string
		args args
		want StructuredFormat
	}{
		{
			name: "StructToStructuredFormat",
			args: args{s: myStruct{}},
			want: StructuredFormat{Type: "object", Properties: map[string]FormatProperty{
				"content": {Type: "string", Description: ""},
				"value":   {Type: "boolean", Description: "test value"},
				"list":    {Type: "array", Description: "test list", Items: ItemProperty{Type: "integer"}},
			}, Required: []string{"content"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StructToStructuredFormat(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("GOT: %+v", got)
				t.Logf("WANT: %+v", tt.want)
				t.Errorf("StructToStructuredFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatOuput_DecodeContent(t *testing.T) {
	type wantResp struct {
		Content string `json:"content"`
	}
	tests := []struct {
		name     string
		o        ChatOuput
		wantResp wantResp
		wantErr  bool
	}{
		{
			name:     "DecodeContent invalid",
			o:        ChatOuput{Content: "hello"},
			wantResp: wantResp{},
			wantErr:  true,
		},
		{
			name:     "DecodeContent valid json",
			o:        ChatOuput{Content: "```\n{\"content\":\"valid\"}\n```"},
			wantResp: wantResp{Content: "valid"},
			wantErr:  false,
		},
		{
			name:     "DecodeContent valid json with type",
			o:        ChatOuput{Content: "```json\n{\"content\":\"valid\"}\n```"},
			wantResp: wantResp{Content: "valid"},
			wantErr:  false,
		},
		{
			name:     "DecodeContent valid json on text",
			o:        ChatOuput{Content: "random text before\n```\n{\"content\":\"hello\"}\n```\nrandom text after"},
			wantResp: wantResp{Content: "hello"},
			wantErr:  false,
		},
		{
			name:     "DecodeContent valid json on text get the last one",
			o:        ChatOuput{Content: "random text before\n```\n{\"content\":\"hello\"}\n```\nrandom text after ```\n{\"content\":\"last\"}\n```"},
			wantResp: wantResp{Content: "last"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &wantResp{}
			if err := tt.o.DecodeContent(resp); (err != nil) != tt.wantErr {
				t.Errorf("ChatOuput.DecodeContent() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(resp, &tt.wantResp) && !tt.wantErr {
				t.Logf("GOT: %+v", resp)
				t.Logf("WANT: %+v", tt.wantResp)
				t.Errorf("ChatOuput.DecodeContent() = %v, want %v", resp, tt.wantResp)
			}
		})
	}
}
