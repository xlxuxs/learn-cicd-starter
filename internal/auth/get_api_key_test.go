package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name	string 
		headers http.Header
		wantKey string 
		wantErr bool
}{
	{
		name: "missing header",
		headers: http.Header{},
		wantKey: "",
		wantErr: true,	
	},
	{
		name: "malinformed header",
		headers: http.Header{
		"Authorization": []string{"Bearer"},},
		wantKey: "",
		wantErr: true,
	},
	{
		name: "valid api key header",
		headers: http.Header{
			"Authorization": []string{"ApiKey my-secret-key"},
		},
		wantKey: "my-secret-key",
		wantErr: false,
	},
}

for _, tc := range tests {
	t.Run(tc.name, func(t *testing.T){
		gotKey, err := GetAPIKey(tc.headers)
		if (err != nil) != tc.wantErr {
			t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
		}
		if gotKey != tc.wantKey {
			t.Errorf("GetAPIKey() key = %q, want %q", gotKey, tc.wantKey)
		}
})
}
}
