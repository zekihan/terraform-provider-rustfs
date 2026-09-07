package provider

import "testing"

func TestMinioTransportInsecure(t *testing.T) {
	transport, err := minioTransport(true, true)
	if err != nil {
		t.Fatalf("creating transport: %v", err)
	}

	if transport.TLSClientConfig == nil || !transport.TLSClientConfig.InsecureSkipVerify {
		t.Fatal("expected insecure transport to skip TLS certificate verification")
	}
}

func TestMinioTransportSecure(t *testing.T) {
	transport, err := minioTransport(true, false)
	if err != nil {
		t.Fatalf("creating transport: %v", err)
	}

	if transport.TLSClientConfig != nil && transport.TLSClientConfig.InsecureSkipVerify {
		t.Fatal("expected secure transport to verify TLS certificates")
	}
}
