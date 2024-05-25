package database

import (
	"os"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models/settings"
)

var pb *pocketbase.PocketBase

func RegisterDB() error {
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: "/opt/pb_data",
	})

	s3Settings := settings.S3Config{
		Enabled:   true,
		Bucket:    os.Getenv("AWS_BUCKET"),
		Endpoint:  os.Getenv("AWS_ENDPOINT"),
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		Secret:    os.Getenv("AWS_SECRET_KEY"),
		Region:    os.Getenv("AWS_REGION"),
	}

	app.Settings().S3 = s3Settings
	app.Settings().Meta.AppName = "yazwalk"

	// initialized app and set it to global `pb`
	pb = app
	return app.Start()

}

func PB() *pocketbase.PocketBase {
	return pb
}

func DB() dbx.Builder {
	return pb.Dao().DB()
}
