package cryptdo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"hash"

	"golang.org/x/crypto/pbkdf2"

	"code.xoeb.us/cryptdo/cryptdopb"
)

type version interface {
	encrypt([]byte, string) (*cryptdopb.Message, error)
	decrypt(*cryptdopb.Message, string) ([]byte, error)
}

var (
	sha384 = sha512.New384

	versions map[int32]version = map[int32]version{
		1: &version1{
			iterations: 100000,
			hashAlg:    sha384,
			saltSize:   sha384().Size(),
			keySize:    32,
			nonceSize:  12,
		},
	}
)

func lookup(vers int32) (version, error) {
	v, ok := versions[vers]
	if !ok {
		return nil, &UnknownVersionError{version: int(vers)}
	}

	return v, nil
}

type version1 struct {
	// key derivation
	iterations int
	hashAlg    func() hash.Hash
	saltSize   int

	// encryption
	keySize   int
	nonceSize int
}

func (v *version1) encrypt(plaintext []byte, passphrase string) (*cryptdopb.Message, error) {
	salt, err := randomBytes(v.saltSize)
	if err != nil {
		return nil, err
	}

	key := v.key(passphrase, salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonce, err := randomBytes(v.nonceSize)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return &cryptdopb.Message{
		Version:    currentVersion,
		Salt:       salt,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}, nil
}

func (v *version1) decrypt(message *cryptdopb.Message, passphrase string) ([]byte, error) {
	key := v.key(passphrase, message.Salt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if aesgcm.NonceSize() != len(message.Nonce) {
		return nil, &InvalidNonceError{
			expected: aesgcm.NonceSize(),
			actual:   len(message.Nonce),
		}
	}

	return aesgcm.Open(nil, message.Nonce, message.Ciphertext, nil)
}

func (v *version1) key(passphrase string, salt []byte) []byte {
	return pbkdf2.Key([]byte(passphrase), salt, v.iterations, v.keySize, v.hashAlg)
}