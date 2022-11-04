package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {

	// reads tmp folder to get keystore filename
	dir := "./tmp"
	files, err := ioutil.ReadDir(dir)
	if err == nil {

		// check if to many keys
		if len(files) > 0 {
			// log.Println("you already have a keystore file. please remove the keystore folder tmp with rm -rf tmp and run go run keystore.go -create true")
			log.Fatal(errors.New("you already have a keystore file. please remove the keystore folder tmp with rm -rf tmp and run go run keystore.go -create true"))
		}
		// log.Fatal(err)
	}

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {

	// create tmporary directory to handle imported key session
	parentDir := os.TempDir()
	logsDir, err := ioutil.TempDir(parentDir, "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tmpdir for import", logsDir)
	defer os.RemoveAll(logsDir)

	// reads tmp folder to get keystore filename
	dir := "./tmp"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) > 1 {
		log.Println("too many keystore files. please remove the keystore folder tmp with rm -rf tmp and run go run keystore.go -create true to create a single new one")
		return
	}

	// for _, f := range files {
	// fmt.Println(f.Name())
	// }
	fileName := files[0].Name()

	// file := "./tmp/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	file := dir + "/" + fileName

	log.Println("reading in keystore:", file)
	// dirNew := "./" + directoryNew
	ks := keystore.NewKeyStore(logsDir, keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	// if err := os.Remove(file); err != nil {
	// log.Fatal(err)
	// }
}

var flagImport = flag.Bool("import", false, "set to true to import keystore")
var flagCreate = flag.Bool("create", false, "set to true to create new keystore")

// var flagDir = flag.String("dir", "tmp", "set to tmp to import old keystore")
// var flagDirNew = flag.String("dirnew", "tmp", "set to tmp1 to create new keystore")

func main() {

	flag.Parse()
	if *flagCreate {
		createKs()
	} else if *flagImport {
		importKs()
	} else {
		log.Println("please use the flag -create true, or -import true to specify if you want to create a keystore or import one")
	}

}
