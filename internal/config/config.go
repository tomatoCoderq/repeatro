package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitConfig(filename string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(filename)
	config.SetConfigType("toml")
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("No config found. Will create one.")
	}

	privKeyPath := config.GetString("keys.private_key_file")
	pubKeyPath := config.GetString("keys.public_key_file")

	if privKeyPath == "" || pubKeyPath == "" || !fileExists(privKeyPath) || !fileExists(pubKeyPath) {
		fmt.Println("Generating new ECDSA key pair...")

		privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			log.Fatalf("Failed to generate keys: %v", err)
		}

		privBytes, err := x509.MarshalECPrivateKey(privKey)
		if err != nil {
			log.Fatalf("Failed to marshal private key: %v", err)
		}
		pubBytes, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
		if err != nil {
			log.Fatalf("Failed to marshal public key: %v", err)
		}

		privKeyPath = "private.pem"
		pubKeyPath = "public.pem"

		writePEM(privKeyPath, "EC PRIVATE KEY", privBytes)
		writePEM(pubKeyPath, "PUBLIC KEY", pubBytes)

		config.Set("keys.private_key_file", privKeyPath)
		config.Set("keys.public_key_file", pubKeyPath)

		configFilePath := filename
		if filepath.Ext(configFilePath) != ".toml" {
			configFilePath += ".toml"
		}

		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			err := config.SafeWriteConfigAs(configFilePath)
			if err != nil {
				log.Fatalf("Failed to create config: %v", err)
			}
		} else {
			err := config.WriteConfigAs(configFilePath)
			if err != nil {
				log.Fatalf("Failed to update config: %v", err)
			}
		}

		fmt.Println("Key files and config written.")
	}

	return config
}

func writePEM(path string, typ string, data []byte) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Failed to create %s: %v", path, err)
	}
	defer f.Close()

	err = pem.Encode(f, &pem.Block{Type: typ, Bytes: data})
	if err != nil {
		log.Fatalf("Failed to encode PEM: %v", err)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
