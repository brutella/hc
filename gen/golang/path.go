package golang

import (
	"github.com/brutella/hc/gen"
	"os"
	"path/filepath"
)

var (
	libPath                = os.ExpandEnv("$GOPATH/src/github.com/brutella/hc")
	ServiceLocalDir        = filepath.Join(libPath, "service")
	CharacteristicLocalDir = filepath.Join(libPath, "characteristic")
	ServiceDir             = "service"
	CharacteristicDir      = "characteristic"
)

// CharacteristicLocalFilePath returns the filepath to a characteristic
func CharacteristicLocalFilePath(char *gen.CharacteristicMetadata) string {
	return filepath.Join(CharacteristicLocalDir, CharacteristicFileName(char))
}

// CharacteristicRelativeFilePath returns the relative filepath to a characteristic
func CharacteristicRelativeFilePath(char *gen.CharacteristicMetadata) string {
	return filepath.Join(CharacteristicDir, CharacteristicFileName(char))
}

// ServiceLocalFilePath returns the filepath to a service
func ServiceLocalFilePath(sv *gen.ServiceMetadata) string {
	return filepath.Join(ServiceLocalDir, ServiceFileName(sv))
}

// ServiceRelativeFilePath returns the relative filepath to a service
func ServiceRelativeFilePath(sv *gen.ServiceMetadata) string {
	return filepath.Join(ServiceDir, ServiceFileName(sv))
}
