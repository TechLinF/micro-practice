package infrastructure

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	"git.rd.rijin.com/micro/utils/database"
	isolator_xorm "github.com/gogap/isolator/extension/xorm"
)

//define a variable
var Engine *xorm.Engine

var (
	XormEngines = make(isolator_xorm.XORMEngines)
)

//initDatabase init mysql database engines
func initDatabase(filename string, conns int) (err error) {
	engines, err := database.InitalXORMEngineFromConfigWithConns(filename, conns)
	if err != nil {
		return
	}

	for _, v := range engines {
		v.Logger().SetLevel(core.LOG_INFO)
	}
	Engine = engines["todo"]
	return
}
