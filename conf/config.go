package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/naoina/toml"
)

type Work struct {
	Name     string
	Desc     string
	Excute   string
	Duration int
	Args     string
}

type Config struct {
	Server struct {
		Mode string
		Port string
	}

	DB map[string]map[string]interface{}

	Work []Work

	Log struct {
		Fpath   string
		Msize   int
		Mbackup int
		Mage    int
		Level   string
	}

	Contract struct {
		PrivateKey      string
		NetUrl          string
		OwnerAddress    string
		ContractAddress string
	}

	KeyStore struct {
		Path string
	}
}

func GetConfig(fpath string) *Config {
	c := new(Config)
	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			jsonBytes, err := ioutil.ReadFile(c.KeyStore.Path)
			if err != nil {
				panic(err)
			} else {
				password := ""
				fmt.Print("keyStore 해금을 위한 Password : ")
				fmt.Scanf("%s", &password)
				account, err := keystore.DecryptKey(jsonBytes, password)
				if err != nil {
					panic(err)
				} else {
					pData := crypto.FromECDSA(account.PrivateKey)
					// Encode시 0x가 접두어로 붙기때문에 제거
					c.Contract.PrivateKey = hexutil.Encode(pData)[2:]
					fmt.Println(c)
					return c
				}
			}
		}
	}
}
