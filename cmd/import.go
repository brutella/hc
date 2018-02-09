// +build ignore

// Imports HomeKit metadata from a file and creates files for every characteristic and service.
// It finishes by running `go fmt` in the characterist and service packages.
//
// The metadata file is created by running the following command on OS X
//
//     plutil -convert json -r -o $GOPATH/src/github.com/brutella/hc/gen/metadata.json /Applications/HomeKit\ Accessory\ Simulator.app/Contents/Frameworks/HAPAccessoryKit.framework/Versions/A/Resources/default.metadata.plist
package main

import (
	"encoding/json"
	"github.com/brutella/hc/gen"
	"github.com/brutella/hc/gen/golang"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var LibPath = os.ExpandEnv("$GOPATH/src/github.com/brutella/hc")
var GenPath = filepath.Join(LibPath, "gen")
var SvcPkgPath = filepath.Join(LibPath, "service")
var AccPkgPath = filepath.Join(LibPath, "accessory")
var CharPkgPath = filepath.Join(LibPath, "characteristic")
var MetadataPath = filepath.Join(GenPath, "metadata.json")

func main() {

	log.Println("Import data from", MetadataPath)

	// Open metadata file
	f, err := os.Open(MetadataPath)
	if err != nil {
		log.Fatal(err)
	}

	// Read content
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// Import json
	metadata := gen.Metadata{}
	err = json.Unmarshal(b, &metadata)
	if err != nil {
		log.Fatal(err)
	}

	// Create characteristic files
	for _, char := range metadata.Characteristics {
		log.Printf("Processing %s Characteristic", char.Name)
		if b, err := golang.CharacteristicGoCode(char); err != nil {
			log.Println(err)
		} else {
			filePath := filepath.Join(CharPkgPath, golang.CharacteristicFileName(char))
			log.Println("Creating file", filePath)
			if f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
				log.Fatal(err)
			} else {
				if _, err := f.Write(b); err != nil {
					log.Fatal()
				}
			}
		}
	}

	// Create service files
	for _, svc := range metadata.Services {
		log.Printf("Processing %s Service", svc.Name)
		if b, err := golang.ServiceGoCode(svc, metadata.Characteristics); err != nil {
			log.Println(err)
		} else {
			filePath := filepath.Join(SvcPkgPath, golang.ServiceFileName(svc))
			log.Println("Creating file", filePath)
			if f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
				log.Fatal(err)
			} else {
				if _, err := f.Write(b); err != nil {
					log.Fatal()
				}
			}
		}
	}

	// Create an accessory categories file
	if b, err := golang.CategoriesGoCode(metadata.Categories); err != nil {
		log.Println(err)
	} else {
		filePath := filepath.Join(AccPkgPath, "constant.go")
		log.Println("Creating file", filePath)
		if f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
			log.Fatal(err)
		} else {
			if _, err := f.Write(b); err != nil {
				log.Fatal()
			}
		}
	}

	log.Println("Running go fmt")

	charCmd := exec.Command("go", "fmt")
	charCmd.Dir = os.ExpandEnv(CharPkgPath)
	if err := charCmd.Run(); err != nil {
		log.Fatal(err)
	}

	svcCmd := exec.Command("go", "fmt")
	svcCmd.Dir = os.ExpandEnv(SvcPkgPath)
	if err := svcCmd.Run(); err != nil {
		log.Fatal(err)
	}

	accCmd := exec.Command("go", "fmt")
	accCmd.Dir = os.ExpandEnv(AccPkgPath)
	if err := accCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
