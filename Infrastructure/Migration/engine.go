package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

var (
	flags = flag.NewFlagSet("hostgator-migration", flag.ExitOnError)
	dir   = flags.String("dir", "./migration", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()
	fmt.Println(args)

	if len(args) < 1 {
		flags.Usage()
		return
	}

	db, err := goose.OpenDBWithDriver("mysql", "root:123@/hostgator")
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	command := args[1]

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
