package update

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/creativeprojects/go-selfupdate"
)

func TestNewConfiguresChecksumValidator(t *testing.T) {
	svc, err := New("mikstew/nametag", "1.0.0")
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	if svc == nil || svc.updater == nil {
		t.Fatal("expected configured update service")
	}
}

func TestChecksumManifestFormat(t *testing.T) {
	binary := []byte("nametag-test-binary")
	sum := sha256.Sum256(binary)
	manifest := fmt.Sprintf("%x  nametag-darwin-arm64\n", sum)

	validator := &selfupdate.ChecksumValidator{UniqueFilename: checksumsFile}
	if err := validator.Validate("nametag-darwin-arm64", binary, []byte(manifest)); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}

	if validator.GetValidationAssetName("nametag-darwin-arm64") != checksumsFile {
		t.Fatalf("expected validation asset %q", checksumsFile)
	}
}

func TestChecksumManifestRejectsMismatch(t *testing.T) {
	binary := []byte("nametag-test-binary")
	manifest := "0000000000000000000000000000000000000000000000000000000000000000  nametag-darwin-arm64\n"

	validator := &selfupdate.ChecksumValidator{UniqueFilename: checksumsFile}
	if err := validator.Validate("nametag-darwin-arm64", binary, []byte(manifest)); err == nil {
		t.Fatal("expected checksum validation to fail")
	}
}
