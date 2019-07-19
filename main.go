/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-26 16:13:46
 */

package main

import (
	"flag"
	"time"

	"git.wemomo.com/cosmos_server/zombieparty-gamecli/cipher"
	cli "git.wemomo.com/cosmos_server/zombieparty-gamecli/cli"
	"git.wemomo.com/cosmos_server/zombieparty-gamecli/protocol"
	logging "github.com/chosen0ne/gologging"
)

func main() {
	host := flag.String("host", "47.93.192.139", "host to connect")
	port := flag.Int("port", 8080, "port to connect")
	uid := flag.String("uid", "100", "user id")
	rid := flag.String("rid", "100", "round id")
	transport := flag.String("transport", "tcp", "transport protocol, tcp or kcp")

	flag.Parse()

	if *transport != "tcp" && *transport != "kcp" {
		logging.Info("transport protocol isn't supported")
		return
	}

	logging.Info("try to connect %s:%d", *host, *port)

	loop := cli.NewLoop(*transport, *host, *port)
	loop.SetId(*uid)
	loop.SetRoundId(*rid)
	loop.AddProtocol(&protocol.ImProtocolImpl{}, &protocol.MessagePackerImpl{})
	loop.SetCipher(&cipher.NoneCipher{})

	loop.Start(*rid)

	time.Sleep(time.Second * 1)
}
