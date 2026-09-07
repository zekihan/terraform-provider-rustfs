package rustfs

import (
	"net/http"
	"testing"
)

func TestNewInsecureTLS(t *testing.T) {
	client := New(&RustfsAdminConfig{Insecure: true, Ssl: true})

	transport, ok := client.httpClient.Transport.(*http.Transport)
	if !ok || transport.TLSClientConfig == nil || !transport.TLSClientConfig.InsecureSkipVerify {
		t.Fatal("expected insecure client to skip TLS certificate verification")
	}
}

func TestNewSecureTLS(t *testing.T) {
	client := New(&RustfsAdminConfig{Ssl: true})

	if client.httpClient.Transport != nil {
		t.Fatal("expected secure client to use the default HTTP transport")
	}
}
