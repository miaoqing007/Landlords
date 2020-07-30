package manager

import (
	"github.com/golang/glog"
	"sync"
)

var _player *Players

type Player struct {
	User *UserManager
}

type Players struct {
	idInMap sync.Map //map[id]*Player
}

func init() {
	_player = &Players{}
}

func (p *Players) getPlayer(id string) *Player {
	if v, ok := p.idInMap.Load(id); ok {
		return v.(*Player)
	}
	return nil
}

func (p *Players) addPlayer(id string, player *Player) {
	p.idInMap.Store(id, player)
}

func (p *Players) deletePlayer(id string) {
	p.idInMap.Delete(id)
}

func GetPlayer(id string) *Player {
	return _player.getPlayer(id)
}

func AddPlayer(id string, player *Player) {
	_player.addPlayer(id, player)
	glog.Infof("addPlayer id = %v", id)
}

func DeletePlayer(id string) {
	_player.deletePlayer(id)
	glog.Infof("deletePlayer id = %v", id)
}
