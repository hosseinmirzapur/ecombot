package database

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
)

var pb *pocketbase.PocketBase

func RegisterDB() error {
	app := pocketbase.New()

	// initialized app and set it to global `pb`
	pb = app
	return app.Start()

}

func DB() dbx.Builder {
	return pb.Dao().DB()
}
