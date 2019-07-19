/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-26 16:33:05
 */

package client

import (
	"bufio"
	"fmt"
	"net"

	"git.wemomo.com/cosmos_server/zombieparty-gamecli/protocol"
	logging "github.com/chosen0ne/gologging"
	"github.com/chosen0ne/goutils"
	"github.com/xtaci/kcp-go"
)

var (
	log *logging.Logger
)

type Client struct {
	net.Conn
	transport string
	host      string
	port      int
	in        *bufio.Reader
	out       *bufio.Writer
	protocol  protocol.ImProtocol
}

func NewClient(transport, host string, port int) *Client {
	return &Client{transport: transport, host: host, port: port}
}

func (cli *Client) setProtocol(p protocol.ImProtocol) {
	cli.protocol = p
}

func (cli *Client) Connect() (err error) {
	hostport := fmt.Sprintf("%s:%d", cli.host, cli.port)

	if cli.transport == "tcp" {
		cli.Conn, err = net.Dial("tcp", hostport)
	} else if cli.transport == "kcp" {
		cli.Conn, err = kcp.Dial(hostport)
	}

	if err != nil {
		log.Exception(err, "failed connect to %s:%d", cli.host, cli.port)
		return err
	}

	log.Info("connected to server %s:%d by %s", cli.host, cli.port, cli.transport)

	cli.in = bufio.NewReader(cli.Conn)
	cli.out = bufio.NewWriter(cli.Conn)

	return nil
}

func (cli *Client) NextMessageGroup() (*protocol.MessagePacket, error) {
	return cli.protocol.Decode(cli.in)
}

func (cli *Client) OutputMessage(m *protocol.MessagePacket) error {
	if msgBytes, err := cli.protocol.Encode(m); err != nil {
		log.Exception(err, "failed to encode msg")
		return err
	} else {
		if err := goutils.WriteBuffer(log, cli.out, msgBytes); err != nil {
			log.Exception(err, "failed to write to socket")
			return err
		}

		if err := cli.out.Flush(); err != nil {
			log.Exception(err, "failed to flush")
			return err
		}
	}

	return nil
}

func init() {
	log = logging.GetLogger("client")
}
