/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-30 17:17:20
 */

package client

import (
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"git.wemomo.com/cosmos_server/zombieparty-gamecli/cipher"
	"git.wemomo.com/cosmos_server/zombieparty-gamecli/protocol"
	protos "git.wemomo.com/cosmos_server/zombieparty-gamecli/protos"
	pb "github.com/golang/protobuf/proto"
)

var (
	pingPacket *protocol.MessagePacket
	sauthAttrs map[string]int
)

type Loop struct {
	cli     *Client
	packer  protocol.MessagePacker
	c       cipher.ICipher
	w       chan int
	wowoid  string
	roundId string
}

func NewLoop(transport, host string, port int) *Loop {
	cli := NewClient(transport, host, port)
	return &Loop{cli: cli, w: make(chan int, 1)}
}

func (l *Loop) SetId(id string) {
	l.wowoid = id
	sauthAttrs[protocol.HEADER_WOWO_ID_LABEL], _ = strconv.Atoi(id)
}

func (l *Loop) SetRoundId(rid string) {
	l.roundId = rid
}

func (l *Loop) AddProtocol(p protocol.ImProtocol, packer protocol.MessagePacker) {
	l.cli.setProtocol(p)
	l.packer = packer
}

func (l *Loop) SetCipher(c cipher.ICipher) {
	l.c = c
}

func (l *Loop) Start(rid string) {
	l.packer.SetCipher(l.c)

	if err := l.cli.Connect(); err != nil {
		log.Exception(err, "failed to connect")
		return
	}

	// auth
	authReq := l.newSauthMsg()
	msgPacket, err := l.packer.PackSauth(authReq, sauthAttrs)
	if err != nil {
		log.Exception(err, "failed to pack sauth")
		return
	}

	if err := l.cli.OutputMessage(msgPacket); err != nil {
		log.Exception(err, "failed to output msg")
		return
	}

	// read input
	go l.inputLoop()

	for _ = range time.Tick(time.Millisecond * 30) {
		l.randomOp(rid)
	}
}

func (l *Loop) randomOp(rid string) {
	input := &protos.InputData{}
	input.Type = rand.Int31() % 3
	input.UserId = l.wowoid
	input.Timestamp = time.Now().UnixNano() / 1000000

	if input.Type == 1 {
		move := &protos.InputData_Move{0.1, 0.1}
		input.Move = move
	}

	sync := &protos.SyncRequest{}
	sync.RoundId = rid
	sync.Data = append(sync.Data, input)

	msgPacket, err := l.packer.PackCommand(sync, nil)
	if err != nil {
		log.Exception(err, "failed to pack sync request")
		return
	}

	if err := l.cli.OutputMessage(msgPacket); err != nil {
		log.Exception(err, "failed to ouput input data")
		return
	}

	log.Info("send op: %s", input.String())
}

func (l *Loop) newSauthMsg() *protos.AuthRequest {
	auth := protos.AuthRequest{
		RoundId: l.roundId,
		UserId:  l.wowoid,
	}

	return &auth
}

func (l *Loop) pingLoop() {
	t := time.Tick(5 * time.Second)
	for _ = range t {
		if err := l.cli.OutputMessage(pingPacket); err != nil {
			log.Exception(err, "failed to output ping packet")
		}
	}
}

func (l *Loop) inputLoop() {
	for {
		if packet, err := l.cli.NextMessageGroup(); err != nil {
			if err == io.EOF {
				log.Error("server closed, exit...")
				os.Exit(-1)
				return
			}
			log.Exception(err, "failed to read next message")
			//l.w <- 1
		} else {
			log.Trace(" >> recv PACKET@{V=%d;H=%d;B=%d}", packet.Version,
				len(packet.Header), len(packet.Body))

			if packet.IsPong() {
				log.Info(" >> recv PONG")
				continue
			}

			switch packet.Header[0] {
			case protocol.HEADER_RESP_TYPE:
				l.recvResp(packet)
			case protocol.HEADER_SYNC_TYPE:
				l.recvSync(packet)
			case protocol.HEADER_PULL_TYPE:
				l.recvPull(packet)
			case protocol.HEADER_START_TYPE:
				l.recvStart(packet)
			default:
				log.Error("unknown msg type, type: %d", packet.Header[0])
			}
		}
	}
}

func (l *Loop) recvResp(packet *protocol.MessagePacket) {
	resp := &protos.Response{}
	if err := pb.Unmarshal(packet.Body, resp); err != nil {
		log.Exception(err, "failed to unmarshal response pb")
		return
	}

	if resp.GetSuccess() == 1 {
		log.Info("recv success response")
	} else {
		log.Info("recv failure response, msg: %s", resp.Errmsg)
	}
}

func (l *Loop) recvSync(packet *protocol.MessagePacket) {
	frame := &protos.Frame{}
	if err := pb.Unmarshal(packet.Body, frame); err != nil {
		log.Exception(err, "failed to unmarshal frame pb")
		return
	}

	log.Info("recv frame, frame: %s, ts: %d", frame, time.Now().UnixNano()/1000000)
}

func (l *Loop) recvPull(packet *protocol.MessagePacket) {

}

func (l *Loop) recvStart(packet *protocol.MessagePacket) {
	startMsg := &protos.GameStartResponse{}
	if err := pb.Unmarshal(packet.Body, startMsg); err != nil {
		log.Exception(err, "failed to unmarshal game start")
		return
	}

	log.Info("recv game start, %s", startMsg)
}

func init() {
	pingPacket = &protocol.MessagePacket{protocol.HEADER_VERSION, nil, nil}

	sauthAttrs = make(map[string]int)
	sauthAttrs[protocol.HEADER_CLI_TYPE_LABEL] = 1
	sauthAttrs[protocol.HEADER_CLI_VERSION_LABEL] = 1000
	sauthAttrs[protocol.HEADER_WOWO_ID_LABEL] = 380118068
}
