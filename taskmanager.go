package main

import (
	"flag"
	"fmt"
	"github.com/Abubakarr99/taskManager/server"
	"github.com/Abubakarr99/taskManager/storage/boltdb"
	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

var (
	addr = flag.String("addr", "0.0.0.0:6742", "The address to run the grpc server")
)

func main() {
	flag.Parse()
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(home, "tasks.db")
	db, err := boltdb.Init(dbPath)
	if err != nil {
		panic(err)
	}
	s, err := server.New(
		*addr,
		db,
		server.WithGRPCOpts(),
	)
	done := make(chan error, 1)
	fmt.Println("Starting server at: ", *addr)
	go func() {
		defer close(done)
		defer func(Db *bolt.DB) {
			err = Db.Close()
			if err != nil {
				panic(err)
			}
		}(db.Db)
		done <- s.Start()
	}()
	fmt.Println("server exited with error: ", <-done)
}
