package tmplt

import (
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/osutil"
)

type Configuration struct {
	ConfigPath    string
	ConfigDirPath string
}

func initializeDB() error {
	if err := osutil.FileExists(DBPath); !os.IsNotExist(err) {
		return err
	}

	db, err := bolt.Open(DBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	} else {
		defer db.Close()
	}

	return db.Update(func(tx *bolt.Tx) error {
		_ = tx.Bucket([]byte("MyBucket"))

		return nil
	})
}

func readConfig() (Configuration, error) {
	db, err := bolt.Open(DBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		exit.Error(err)
	} else {
		defer db.Close()
	}

	var conf Configuration
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))

		conf.ConfigPath = string(b.Get([]byte("configPath")))
		conf.ConfigDirPath = string(b.Get([]byte("configDirPath")))

		return nil
	})

	return conf, err
}
