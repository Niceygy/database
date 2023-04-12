package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const dbPath = "/var/ndb"
const dbList = "/var/ndb/db.list"

func firstStart() {
	log.Println("First start detected, creating database directory...")
	app := "mkdir"
	//dbPath := "/var/ndb"

	cmd := exec.Command(app, dbPath)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(string(stdout))

}

func checkFirstStart() {
	// check if /var/ndb exists
	// if not, run firstStart()
	//dbPath := "/var/ndb"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		firstStart()
	} else {
		return
	}
}

func ReadDatabaseList() {
	// read the database list
	// return a list of databases
	//db list is in the root folder, under "db.list"
	//db.list is a list of database names, separated by newlines
	//each database is a folder in /var/ndb

	readFile, err := os.Open(dbList)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()

}

func init() {
	checkFirstStart()
}
