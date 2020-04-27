package hash

import "golang.org/x/crypto/bcrypt"

func NewBCrypt(pepper string) BCrypt {
	return BCrypt{pepper, []byte(pepper)}
}

type BCrypt struct {
	pepper      string
	pepperBytes []byte
}

func (bc BCrypt) Equal(hash, raw string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(raw+bc.pepper))
	return err == nil
}

func (bc BCrypt) Bytes(input []byte) []byte {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		append(input, bc.pepperBytes...),
		bcrypt.DefaultCost)
	if err != nil {
		// This shouldn't happen in practice unless our code
		// is just very wrong, so we can panic on error
		panic(err)
	}
	return hashedBytes
}

func (bc BCrypt) String(input string) string {
	return string(bc.Bytes([]byte(input)))
}
