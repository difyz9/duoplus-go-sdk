package duoplus

import (
	"encoding/json"
	"testing"
)

func TestTextUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data string
		want Text
	}{
		{name: "string", data: `"abc"`, want: Text("abc")},
		{name: "number", data: `12345`, want: Text("12345")},
		{name: "bool", data: `true`, want: Text("true")},
		{name: "null", data: `null`, want: Text("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Text
			if err := json.Unmarshal([]byte(tt.data), &got); err != nil {
				t.Fatalf("Unmarshal() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("Unmarshal() got = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewClientRequiresAPIKey(t *testing.T) {
	if _, err := NewClient(""); err == nil {
		t.Fatal("expected error for empty api key")
	}
}

func TestNewClientInitializesServices(t *testing.T) {
	client, err := NewClient("test-key")
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if client.CloudPhones == nil || client.CloudNumbers == nil || client.Groups == nil {
		t.Fatal("expected core service fields to be initialized")
	}

	if client.Proxies == nil || client.SubscriptionStartups == nil || client.Apps == nil {
		t.Fatal("expected module service fields to be initialized")
	}

	if client.CloudDisk == nil || client.Automation == nil {
		t.Fatal("expected auxiliary service fields to be initialized")
	}
}
