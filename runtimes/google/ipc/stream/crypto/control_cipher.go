package crypto

// ControlCipher provides the ciphers and MAC for control channel encryption.
// Encryption and decryption are performed in place.
type ControlCipher interface {
	// MACSize returns the number of bytes in the MAC.
	MACSize() int

	// Seal replaces the message with an authenticated and encrypted version.
	// The trailing MACSize bytes of the data are used for the MAC; they are
	// discarded and overwritten.
	Seal(data []byte) error

	// Open authenticates and decrypts a box produced by Seal.  The trailing
	// MACSize bytes are not changed.  Returns true on success.
	Open(data []byte) bool

	// Encrypt encrypts the data in place.  No MAC is added.
	Encrypt(data []byte)

	// Decrypt decrypts the data in place.  No MAC is verified.
	Decrypt(data []byte)
}
