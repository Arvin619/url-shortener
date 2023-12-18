package usecase

import (
	"testing"
)

func TestAddShortUrlCommand_validate(t *testing.T) {
	type fields struct {
		OriginalUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "pass-1 (https)",
			fields: fields{
				OriginalUrl: "https://google.com",
			},
			wantErr: false,
		},
		{
			name: "pass-2 (http)",
			fields: fields{
				OriginalUrl: "http://google.com",
			},
			wantErr: false,
		},
		{
			name: "pass-3 (query string)",
			fields: fields{
				OriginalUrl: "https://google.com?hello=world",
			},
			wantErr: false,
		},
		{
			name: "pass-4 (path)",
			fields: fields{
				OriginalUrl: "https://google.com/abc123",
			},
			wantErr: false,
		},
		{
			name: "pass-5 (path and query string)",
			fields: fields{
				OriginalUrl: "https://google.com/abc123?hello=world&xyz=456",
			},
			wantErr: false,
		},
		{
			name: "fail-1 (少一個 /)",
			fields: fields{
				OriginalUrl: "https:/google.com",
			},
			wantErr: true,
		},
		{
			name: "fail-2 (少兩個 /)",
			fields: fields{
				OriginalUrl: "https:google.com",
			},
			wantErr: true,
		},
		{
			name: "fail-3 (少 :)",
			fields: fields{
				OriginalUrl: "https//google.com",
			},
			wantErr: true,
		},
		{
			name: "fail-4 (少 com)",
			fields: fields{
				OriginalUrl: "https://google",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &AddShortUrlCommand{
				OriginalUrl: tt.fields.OriginalUrl,
			}
			if err := cmd.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
