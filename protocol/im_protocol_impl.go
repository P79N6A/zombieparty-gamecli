/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-26 17:10:44
 */

package protocol

import (
	"bufio"
	"bytes"

	logging "github.com/chosen0ne/gologging"
	"github.com/chosen0ne/goutils"
)

const (
	VERSION    = 1
	HEADER_LEN = 4

	MSG_TYPE_PING  = 0
	MSG_TYPE_SAUTH = 1
	MSG_TYPE_CMD   = 2
)

var (
	log *logging.Logger
)

type ImProtocolImpl struct{}

func (p *ImProtocolImpl) Decode(r *bufio.Reader) (*MessagePacket, error) {

	if version, err := r.ReadByte(); err != nil {
		log.Exception(err, "failed to read version byte")
		return nil, err
	} else if version != VERSION {
		log.Error("unexpected version byte, need: %d, read: %d", VERSION, version)
		return nil, goutils.NewErr("unexpected version byte")
	}

	// pong packet
	if typeByte, err := r.ReadByte(); err != nil {
		log.Exception(err, "failed to read type byte")
		return nil, err
	} else if typeByte == MSG_TYPE_PING {
		return &MessagePacket{VERSION, nil, nil}, nil
	}

	if err := r.UnreadByte(); err != nil {
		log.Exception(err, "failed to unread type byte")
		return nil, err
	}
	h := make([]byte, HEADER_LEN)
	if err := readFixedLen(r, h); err != nil {
		log.Exception(err, "failed to read header")
		return nil, goutils.WrapErr(err)
	}

	bodyLen := bodyLen(h)
	b := make([]byte, bodyLen)
	if err := readFixedLen(r, b); err != nil {
		log.Exception(err, "failed to read body")
		return nil, goutils.WrapErr(err)
	}

	return &MessagePacket{VERSION, h, b}, nil
}

func (p *ImProtocolImpl) Encode(m *MessagePacket) ([]byte, error) {
	b := &bytes.Buffer{}
	if err := b.WriteByte(VERSION); err != nil {
		log.Exception(err, "failed to write version of header")
		return nil, err
	}

	// 'ping' packet
	if m.Header == nil && m.Body == nil {
		if err := b.WriteByte(MSG_TYPE_PING); err != nil {
			log.Exception(err, "failed to write ping type")
			return nil, err
		}

		return b.Bytes(), nil
	}

	// 'sauth' or 'command'

	if err := goutils.WriteBuffer(log, b, m.Header); err != nil {
		log.Exception(err, "failed to write header")
		return nil, err
	}

	if err := goutils.WriteBuffer(log, b, m.Body); err != nil {
		log.Exception(err, "failed to write body")
		return nil, err
	}

	return b.Bytes(), nil
}

func bodyLen(header []byte) int32 {
	return int32((header[2]&0x000000ff)<<8) + int32(header[3]&0x000000ff)
}

func msgType(header []byte) byte {
	return header[0]
}

func readFixedLen(r *bufio.Reader, b []byte) error {
	if n, err := r.Read(b); err != nil {
		log.Exception(err, "failed to read")
		return err
	} else if n != len(b) {
		log.Exception(err, "failed to read fixed length bytes, need: %d, read: %d",
			len(b), n)
		return goutils.NewErr("no enough bytes")
	}

	return nil
}

func init() {
	if err := logging.Load("logger.conf"); err != nil {
		logging.Exception(err, "failed load logger conf")
	}

	log = logging.GetLogger("protocol")
}
