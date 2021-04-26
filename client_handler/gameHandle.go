package client_handler

import (
	"landlords/client_proto"
	"landlords/helper/util"
	"landlords/manager"
	"landlords/registry"
	"landlords/session"
)

//开始游戏
func P_start_game_req(sess *session.Session, data []byte) (int16, interface{}) {
	tbl, _ := client_proto.PKT_auto_id(data)
	manager.AddPlayer2PvpPool(int(tbl.F_id), sess.User.Id, sess.User.Name)
	return Code["start_game_req"], nil
}

func P_cancel_match_req(sess *session.Session, data []byte) (int16, interface{}) {
	tbl, _ := client_proto.PKT_auto_id(data)
	manager.RemovePlayer4PvpPool(int(tbl.F_id), sess.User.Id)
	return Code["cancel_match_success_ack"], nil
}

//出牌
func P_out_of_the_card_req(sess *session.Session, data []byte) (int16, interface{}) {
	tbl, _ := client_proto.PKT_player_outof_card(data)
	if len(tbl.F_cards) == 0 {
		return Code["out_of_the_card_failed_ack"], client_proto.S_error_ack{"请选择正确的出牌数"}
	}
	room := manager.GetRoomManager(tbl.F_roomId)
	if room == nil {
		return Code["out_of_the_card_failed_ack"], client_proto.S_error_ack{"房间错误"}
	}
	if !room.CheckHandCards(sess.User.Id, tbl.F_cards) {
		return Code["out_of_the_card_failed_ack"], client_proto.S_error_ack{"出牌不符合规则"}
	}
	room.DeleteCards(sess.User.Id, tbl.F_cards)

	registry.PushRoom(sess.User.GetRoomId(), 2011, client_proto.S_out_of_cards{sess.User.Id,
		util.SortArrayStringBig2Small(room.GetUserCard4Array(sess.User.Id)), util.SortArrayStringSmall2Big(tbl.F_cards)})

	if len(room.GetUserCard4Array(sess.User.Id)) == 0 {
		registry.PushRoom(sess.User.GetRoomId(), 2017, client_proto.S_game_over{sess.User.Id})
		return 0, nil
	}
	return 0, nil
}

func P_give_up_card_req(sess *session.Session, data []byte) (int16, interface{}) {
	tbl, _ := client_proto.PKT_player_outof_card(data)
	room := manager.GetRoomManager(tbl.F_roomId)
	if room == nil {
		return Code["out_of_the_card_failed_ack"], client_proto.S_error_ack{"房间错误"}
	}
	registry.PushRoom(sess.User.GetRoomId(), 2011, client_proto.S_out_of_cards{sess.User.Id,
		util.SortArrayStringBig2Small(room.GetUserCard4Array(sess.User.Id)), util.SortArrayStringSmall2Big([]string{})})
	return 0, nil
}
