package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"br.com.techbh/pcc"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	host := ""
	user := ""
	pass := ""
	base := ""

	//command args
	hostPtr := flag.String("host", "127.0.0.1", "your data host")
	userPtr := flag.String("user", "user", "user")
	passPtr := flag.String("pass", "pass", "password")
	basePtr := flag.String("base", "schema", "database")
	portPtr := flag.Uint64("port", 3306, "")
	flag.Parse()

	//necessary to run container with env data
	//if has env vars, use them instead
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
		hostPtr = &host
	}
	if os.Getenv("DB_USER") != "" {
		user = os.Getenv("DB_USER")
		userPtr = &user
	}
	if os.Getenv("DB_PASS") != "" {
		pass = os.Getenv("DB_PASS")
		passPtr = &pass
	}
	if os.Getenv("DB_NAME") != "" {
		base = os.Getenv("DB_NAME")
		basePtr = &base
	}

	cAdmin := Connection{*hostPtr, *userPtr, *passPtr, *basePtr, *portPtr}
	fmt.Println("dsn = ", cAdmin.getUndisclosedDSN())
	dbAdmin, err := sql.Open("mysql", cAdmin.getDSN())
	defer dbAdmin.Close()
	if err != nil {
		log.Fatal("Erro ao parametrizar acesso à base principal")
		os.Exit(-1)
	}

	err = dbAdmin.Ping()
	if err != nil {
		log.Fatal("Acesso indisponível... ", err)
		os.Exit(-1)
	}

	fmt.Println("Now, you're ready to GO")

	//channel to receive active companies
	cEmp := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			cEmp <- i
		}
	}()

	//process channel with one or multiple go-routines
	for e := range cEmp {
		pcc.Process(e)
	}

}
