package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var (
	selfhost             = "http://127.0.0.1:18545"
	contractAddress      = "0xd0c696767a2053d2f4dDF89bA894973D2b026834"
	selfhostkovan        = "https://kovan.infura.io"
	contractAddresskovan = "0xa23bb04f42be5d00be77fb912f8553b76c51889b"

	///kovan
	userAddress1    = "0x7b6d393414d77cce0177ecd6883b08a0b5b881ea"
	userAddress2    = "0xe143c83f4e365928a1bc41034d3dc4ec8e5c47d9"
	userKeystore1   = `{"id":"16f2aa71-435d-66a9-8dc9-4c4889103772","version":3,"crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"11d6de57754581d92c701fcda0312e22"},"ciphertext":"8c5593f40055c9d686b5901036ef180c8f2e856f7b608baa56c45ce9f2371e5d","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"552e3fe91c2c6ac4241ca7ca1d44979081cc5a7890f4647d8e38b899816c08bc"},"mac":"c6fa315e304b7625d4bd2b28b26d8f3ca1774c0e2e3461b8ae9faec1156b068b"},"address":"7b6d393414d77cce0177ecd6883b08a0b5b881ea","name":"","meta":"{}"}`
	userPassphrase1 = "123456789"
	userKeystore2   = `{"id":"954ae627-bd39-6a38-bb0b-8f290dfe8dc9","version":3,"crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"92c2dd85471721fd7e0997ab9f5703c4"},"ciphertext":"cddc315e72d89a50228d4244e1e390413e12f39ed0fbd348312ebc39b1108175","kdf":"pbkdf2","kdfparams":{"c":10240,"dklen":32,"prf":"hmac-sha256","salt":"d8afc1723758d985273eb08979d2c9542757bd57f26e841e2372b507a80ec15f"},"mac":"fed2fdb5b538e65c9ee11c47b7d6f31f17eab1abb6776abcae1fe8d9573e959c"},"address":"e143c83f4e365928a1bc41034d3dc4ec8e5c47d9","name":"","meta":"{}"}`
	userPassphrase2 = "123456789"
	///////////////////ganache
	userAddress11   = "0xA6e2d4666b0c4B73bDbb54CA8209bfD0EC89Ea88"
	userAddress22   = "0x1eB5E96518f6DA9FFcf32a401F26e2a8CA9E10b7"
	userPrivateKey1 = "8b01fcf00efb07a877e57c69811b414586ec1e7d531443ab0d88a9093b05f015"
	userPrivateKey2 = "6a7af08e232aacd433158d5586d513be64d97044e83b7f5738bc345917ba6c17"
)

func LoadConfig(configFileName string, cfg interface{}) bool {

	var err error

	//configFileName := "api.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}

	configFileName, err = filepath.Abs(configFileName)
	if err != nil {
		log.Fatal("LoadConfig", "filepath.Abs err", err)
		return false
	}
	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("LoadConfig", "os.Open File error: ", err.Error())
		return false
	}
	defer configFile.Close()

	if _, err = toml.DecodeReader(configFile, cfg); err != nil {
		log.Fatal("LoadConfig", "toml.DecodeReader error: ", err.Error())
		return false
	}

	return true
}
