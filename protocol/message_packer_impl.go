/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-30 16:16:22
 */

package protocol

import (
	"bytes"
	"fmt"

	"github.com/chosen0ne/goutils"
	pb "github.com/golang/protobuf/proto"
	"wowo.com/game-server-cli/cipher"
	protos "wowo.com/game-server-cli/protos"
)

const (
	HEADER_VERSION byte = 0
	PROTO_VERSION  byte = 0

	HEADER_SAUTH_LEN_MAX = 24
	HEADER_CMD_LEN       = 4

	HEADER_SAUTH_TYPE byte = 1
	HEADER_SYNC_TYPE  byte = 2
	HEADER_PULL_TYPE  byte = 3
	HEADER_START_TYPE byte = 4
	HEADER_RESP_TYPE  byte = 5

	HEADER_CLI_TYPE_LABEL    = "cli-type"
	HEADER_CLI_VERSION_LABEL = "cli-version"
	HEADER_WOWO_ID_LABEL     = "wowo-id"
)

type MessagePackerImpl struct {
	c cipher.ICipher
}

func (p *MessagePackerImpl) SetCipher(c cipher.ICipher) {
	p.c = c
}

func (p *MessagePackerImpl) PackSauth(m *protos.AuthRequest, attrs map[string]int) (*MessagePacket, error) {
	// wowo id
	wowoid_str := fmt.Sprintf("%d", attrs[HEADER_WOWO_ID_LABEL])
	wowoid_bytes := []byte(wowoid_str)
	log.Info("WOWOID: %s", wowoid_str)

	header := make([]byte, 7+len(wowoid_str))
	typeVal := HEADER_SAUTH_TYPE

	header[0] = typeVal
	header[1] = PROTO_VERSION

	// set body length
	body, err := p.packSauthBody(m)
	if err != nil {
		log.Exception(err, "failed to pack body")
		return nil, err
	}

	header[2] = byte(len(body) & (0x000000ff << 8) >> 8)
	header[3] = byte(len(body) & 0x000000ff)

	// cli type
	header[4] = byte(attrs[HEADER_CLI_TYPE_LABEL])

	// cli version
	header[5] = byte(attrs[HEADER_CLI_VERSION_LABEL] & (0x000000ff << 8) >> 8)
	header[6] = byte(attrs[HEADER_CLI_VERSION_LABEL] & 0x000000ff)

	for i := 0; i < len(wowoid_bytes); i++ {
		if i == len(wowoid_bytes)-1 {
			wowoid_bytes[i] = (wowoid_bytes[i] << 1 >> 1) + 0x00
		} else {
			wowoid_bytes[i] = (wowoid_bytes[i] << 1 >> 1) + 0x80
		}
		log.Info("WOWOID byte %d: %x", i, wowoid_bytes[i])
		header[7+i] = wowoid_bytes[i]
	}

	return &MessagePacket{HEADER_VERSION, header, body}, nil
}

func (p *MessagePackerImpl) PackCommand(m *protos.SyncRequest, attrs map[string]int) (*MessagePacket, error) {
	header := make([]byte, HEADER_CMD_LEN)

	header[0] = HEADER_SYNC_TYPE
	header[1] = PROTO_VERSION

	body, err := p.packCommandBody(m)
	if err != nil {
		log.Exception(err, "failed to pack body")
		return nil, err
	}

	header[2] = byte(len(body) & (0x000000ff << 8) >> 8)
	header[3] = byte(len(body) & 0x000000ff)

	return &MessagePacket{HEADER_VERSION, header, body}, nil
}

func (p *MessagePackerImpl) packSauthBody(m *protos.AuthRequest) ([]byte, error) {
	msgBytes, err := pb.Marshal(m)
	if err != nil {
		log.Exception(err, "failed to marshal msg in pb format")
		return nil, err
	}

	encryptBytes, err := p.c.Encrypt(msgBytes)
	if err != nil {
		log.Exception(err, "failed to encryp msg")
		return nil, err
	}

	// sauth message: need to pack cipher metadata
	metadata, err := p.c.Metadata(nil)
	if err != nil {
		log.Exception(err, "failed to fetch cipher metadata")
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := goutils.WriteBuffer(log, buf, metadata); err != nil {
		log.Exception(err, "failed to write cipher metadata")
		return nil, err
	}
	if err := goutils.WriteBuffer(log, buf, encryptBytes); err != nil {
		log.Exception(err, "failed to write encrypted bytes")
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *MessagePackerImpl) packCommandBody(m *protos.SyncRequest) ([]byte, error) {
	msgBytes, err := pb.Marshal(m)
	if err != nil {
		log.Exception(err, "failed to marshal msg in pb format")
		return nil, err
	}

	encryptBytes, err := p.c.Encrypt(msgBytes)
	if err != nil {
		log.Exception(err, "failed to encryp msg")
		return nil, err
	}

	return encryptBytes, nil
}
