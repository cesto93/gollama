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
		name    string
		args    args
		want    StructuredFormat
		wantErr bool
	}{
		{
			name: "StructToStructuredFormat",
			args: args{s: myStruct{}},
			want: StructuredFormat{Type: "object", Properties: map[string]FormatProperty{
				"content": {Type: "string", Description: ""},
				"value":   {Type: "boolean", Description: "test value"},
				"list":    {Type: "array", Description: "test list", Items: ItemProperty{Type: "integer"}},
			}, Required: []string{"content"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToStructuredFormat(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToStructuredFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("GOT: %+v", got)
				t.Logf("WANT: %+v", tt.want)
				t.Errorf("StructToStructuredFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatOuput_DecodeContent(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		o       ChatOuput
		args    args
		wantErr bool
	}{
		{
			name:    "DecodeContent invalid",
			o:       ChatOuput{Content: "hello"},
			args:    args{v: &ChatOuput{}},
			wantErr: true,
		},
		{
			name:    "DecodeContent valid json",
			o:       ChatOuput{Content: "```\n{\"content\":\"hello\"}\n```"},
			args:    args{v: &ChatOuput{}},
			wantErr: false,
		},
		{
			name:    "DecodeContent valid json on text",
			o:       ChatOuput{Content: "random text before\n```\n{\"content\":\"hello\"}\n```\nrandom text after"},
			args:    args{v: &ChatOuput{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.o.DecodeContent(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ChatOuput.DecodeContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
