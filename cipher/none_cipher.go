/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-31 16:51:00
 */

package cipher

type NoneCipher struct{}

func (c *NoneCipher) Encrypt(input []byte) ([]byte, error) {
	return input, nil
}

func (c *NoneCipher) Decrypt(input []byte) ([]byte, error) {
	return input, nil
}

func (c *NoneCipher) Metadata(attrs interface{}) ([]byte, error) {
	return []byte{0}, nil
}
