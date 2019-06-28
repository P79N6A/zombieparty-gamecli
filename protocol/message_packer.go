/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-30 16:10:13
 */

package protocol

import (
	"git.wemomo.com/cosmos_server/zombieparty-gamecli/cipher"
	protos "git.wemomo.com/cosmos_server/zombieparty-gamecli/protos"
)

type MessagePacker interface {
	// ImMessage(in pb format) => MessagePacket
	// @param attrs, used to pass extral parameter. For sauth, use to pass
	//		more parameters.
	PackSauth(m *protos.AuthRequest, attrs map[string]int) (*MessagePacket, error)
	PackCommand(m *protos.SyncRequest, attrs map[string]int) (*MessagePacket, error)
	SetCipher(c cipher.ICipher)
}
