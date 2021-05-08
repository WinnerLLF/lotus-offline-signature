package secp

import (
	"fmt"
	"lotus-offline-signature/sigs"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"
	"github.com/minio/blake2b-simd"
)

type secpSigner struct{}

/**
 * @Description:
 * @receiver secpSigner
 * @return []byte
 * @return error
 */
func (secpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}

/**
 * @Description:
 * @receiver secpSigner
 * @param pk
 * @return []byte
 * @return error
 */
func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil
}

/**
 * @Description:
 * @receiver secpSigner
 * @param pk
 * @param msg
 * @return []byte
 * @return error
 */
func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {
		return nil, err
	}

	return sig, nil
}

/**
 * @Description:
 * @receiver secpSigner
 * @param sig
 * @param a
 * @param msg
 * @return error
 */
func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}

	return nil
}

/**
 * @Description:
 */
func init() {
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})
}
