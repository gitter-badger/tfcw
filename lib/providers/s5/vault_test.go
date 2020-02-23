package s5

import (
	"os"
	"testing"

	"github.com/mvisonneau/go-helpers/test"
	"github.com/mvisonneau/s5/cipher"
	cipherVault "github.com/mvisonneau/s5/cipher/vault"
	"github.com/mvisonneau/tfcw/lib/schemas"
)

const (
	testVaultTransitKey string = "foo"
)

func TestGetCipherEngineVault(t *testing.T) {
	cipherEngineType := schemas.S5CipherEngineTypeVault
	key := testVaultTransitKey

	// expected engine
	expectedEngine, err := cipher.NewVaultClient(key)
	test.Expect(t, err, nil)

	// all defined in client, empty variable config (default settings)
	v := &schemas.S5{}
	c := &Client{
		CipherEngineType: &cipherEngineType,
		CipherEngineVault: &schemas.S5CipherEngineVault{
			TransitKey: &key,
		},
	}

	cipherEngine, err := c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherVault.Client).Config, expectedEngine.Config)

	// all defined in variable, empty client config
	c = &Client{}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
		CipherEngineVault: &schemas.S5CipherEngineVault{
			TransitKey: &key,
		},
	}

	cipherEngine, err = c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherVault.Client).Config, expectedEngine.Config)

	// key defined in environment variable
	os.Setenv("S5_VAULT_TRANSIT_KEY", testVaultTransitKey)
	c = &Client{}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
	}

	cipherEngine, err = c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherVault.Client).Config, expectedEngine.Config)

	// other engine & key defined in client, overridden in variable
	otherCipherEngineType := schemas.S5CipherEngineTypeAES
	otherTransitKey := "bar"

	c = &Client{
		CipherEngineType: &otherCipherEngineType,
		CipherEngineVault: &schemas.S5CipherEngineVault{
			TransitKey: &otherTransitKey,
		},
	}
	v = &schemas.S5{
		CipherEngineType: &cipherEngineType,
		CipherEngineVault: &schemas.S5CipherEngineVault{
			TransitKey: &key,
		},
	}

	cipherEngine, err = c.getCipherEngine(v)
	test.Expect(t, err, nil)
	test.Expect(t, cipherEngine.(*cipherVault.Client).Config, expectedEngine.Config)
}