/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-26 16:13:46
 */

package main

import (
	"flag"
	"time"

	logging "github.com/chosen0ne/gologging"
	"wowo.com/game-server-cli/cipher"
	cli "wowo.com/game-server-cli/cli"
	"wowo.com/game-server-cli/protocol"
)

func main() {
	host := flag.String("host", "127.0.0.1", "host to connect")
	port := flag.Int("port", 8888, "port to connect")
	id := flag.String("uid", "100", "user id")
	flag.Parse()

	logging.Info("try to connect %s:%d", *host, *port)

	loop := cli.NewLoop(*host, *port)
	loop.SetId(*id)
	loop.SetRoundId("1")
	loop.AddProtocol(&protocol.ImProtocolImpl{}, &protocol.MessagePackerImpl{})
	loop.SetCipher(&cipher.NoneCipher{})

	loop.Start()

	time.Sleep(time.Second * 1)
}
