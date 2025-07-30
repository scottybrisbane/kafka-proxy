package awsmskiamprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory_New(t *testing.T) {
	tests := []struct {
		name    string
		params  []string
		wantErr bool
	}{
		{
			name:    "missing region",
			params:  []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := &Factory{}
			provider, err := factory.New(tt.params)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, provider)
			}
		})
	}
}

func TestFactory_New_ValidParams(t *testing.T) {
	// Test that valid parameters are parsed correctly
	// Note: We don't test actual token generation as it requires AWS credentials
	factory := &Factory{}

	// Test with region only - this will fail due to missing AWS credentials
	// but we can test that the factory properly parses the parameters
	provider, err := factory.New([]string{"--region", "us-east-1"})
	// In test environment without AWS credentials, this will fail
	assert.Error(t, err)
	assert.Nil(t, provider)
	// But we can verify the error is related to token generation
	assert.Contains(t, err.Error(), "token")
}

func TestTokenProvider_GetToken(t *testing.T) {
	// This test would require AWS credentials and a real MSK cluster
	// For now, we'll just test the structure
	provider := &TokenProvider{
		timeout:     10,
		region:      "us-east-1",
		profile:     "",
		roleARN:     "",
		externalID:  "",
		sessionName: "test",
	}

	assert.NotNil(t, provider)
	assert.Equal(t, "us-east-1", provider.region)
	assert.Equal(t, 10, int(provider.timeout))
}

func TestTokenProvider_shouldRenew(t *testing.T) {
	provider := &TokenProvider{}

	// Test with no current token
	assert.True(t, provider.shouldRenew())

	// Test with current token but no expiry set
	provider.currentToken = "test-token"
	assert.True(t, provider.shouldRenew())
}
