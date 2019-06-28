/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-31 16:47:02
 */

package cipher

type ICipher interface {
	Encrypt(input []byte) ([]byte, error)
	Decrypt(input []byte) ([]byte, error)
	Metadata(attrs interface{}) ([]byte, error)
}
