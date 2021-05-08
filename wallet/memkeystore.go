package wallet

import (
	"lotus-offline-signature/types"
)

type MemKeyStore struct {
	m map[string]types.KeyInfo
}

func NewMemKeyStore() *MemKeyStore {
	return &MemKeyStore{
		make(map[string]types.KeyInfo),
	}
}

/**
 * @Description: List lists all the keys stored in the KeyStore
 * @receiver mks
 * @return []string
 * @return error
 */
func (mks *MemKeyStore) List() ([]string, error) {
	var out []string
	for k := range mks.m {
		out = append(out, k)
	}
	return out, nil
}

/**
 * @Description: Get gets a key out of keystore and returns KeyInfo corresponding to named key
 * @receiver mks
 * @param k
 * @return types.KeyInfo
 * @return error
 */
func (mks *MemKeyStore) Get(k string) (types.KeyInfo, error) {
	ki, ok := mks.m[k]
	if !ok {
		return types.KeyInfo{}, types.ErrKeyInfoNotFound
	}

	return ki, nil
}

/**
 * @Description: Put saves a key info under given name
 * @receiver mks
 * @param k
 * @param ki
 * @return error
 */
func (mks *MemKeyStore) Put(k string, ki types.KeyInfo) error {
	mks.m[k] = ki
	return nil
}

/**
 * @Description: Delete removes a key from keystore
 * @receiver mks
 * @param k
 * @return error
 */
func (mks *MemKeyStore) Delete(k string) error {
	delete(mks.m, k)
	return nil
}

var _ (types.KeyStore) = (*MemKeyStore)(nil)
