package client_handle

import (
	"app/client_proto"
	"app/misc/packet"
	"app/session"
	"github.com/golang/glog"
)

func P_heart_beat_req(session *session.Session, reader *packet.Packet) [][]byte {
	tbl, _ := client_proto.PKT_auto_id(reader)
	glog.Info("tbl", tbl)
	return [][]byte{
		packet.Pack(Code["heart_beat_ack"], client_proto.S_auto_id{86}, nil),
	}
}
