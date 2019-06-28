/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-26 16:48:06
 */

package protocol

import (
	"bufio"
)

type MessagePacket struct {
	Version byte
	Header  []byte
	Body    []byte
}

type ImProtocol interface {
	Decode(r *bufio.Reader) (*MessagePacket, error)
	Encode(m *MessagePacket) ([]byte, error)
}

func (p *MessagePacket) IsPong() bool {
	return p.Header == nil && p.Body == nil
}
