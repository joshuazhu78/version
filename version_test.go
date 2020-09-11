package version_test

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/free5gc/version"
)

func TestVersion(t *testing.T) {
	t.Run("VERSION not specified", func(t *testing.T) {
		var expected = fmt.Sprintf(
			"\n\tNot specify ldflags (which link version) during go build\n\tgo version: %s %s/%s",
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH)
		assert.Equal(t, expected, version.GetVersion())
	})

	t.Run("VERSION constly specified", func(t *testing.T) {
		version.VERSION = "Release-v3.100.200"
		version.BUILD_TIME = "2020-09-11T07:05:04Z"
		version.COMMIT_HASH = "fb2481c2"
		version.COMMIT_TIME = "2020-09-11T07:00:29Z"

		var expected = fmt.Sprintf(
			"\n\tfree5GC version: %s"+
				"\n\tbuild time:      %s"+
				"\n\tcommit hash:     %s"+
				"\n\tcommit time:     %s"+
				"\n\tgo version:      %s %s/%s",
			version.VERSION,
			version.BUILD_TIME,
			version.COMMIT_HASH,
			version.COMMIT_TIME,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH)

		assert.Equal(t, expected, version.GetVersion())
		fmt.Println(string(version.VERSION))
	})

	t.Run("VERSION capture by system", func(t *testing.T) {
		var stdout []byte
		version.VERSION = "Release-v3.100.200" // VERSION using free5gc's version (git tag), we static set it here
		stdout, _ = exec.Command("bash", "-c", "date -u +\"%Y-%m-%dT%H:%M:%SZ\"").Output()
		version.BUILD_TIME = strings.TrimSuffix(string(stdout), "\n")
		stdout, _ = exec.Command("bash", "-c", "git log --pretty=\"%H\" -1 | cut -c1-8").Output()
		version.COMMIT_HASH = strings.TrimSuffix(string(stdout), "\n")
		stdout, _ = exec.Command("bash", "-c", "git log --pretty=\"%ai\" -1 | awk '{time=$1\"T\"$2\"Z\"; print time}'").Output()
		fmt.Println("Insert Data")
		version.COMMIT_TIME = strings.TrimSuffix(string(stdout), "\n")

		var expected = fmt.Sprintf(
			"\n\tfree5GC version: %s"+
				"\n\tbuild time:      %s"+
				"\n\tcommit hash:     %s"+
				"\n\tcommit time:     %s"+
				"\n\tgo version:      %s %s/%s",
			version.VERSION,
			version.BUILD_TIME,
			version.COMMIT_HASH,
			version.COMMIT_TIME,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH)

		assert.Equal(t, expected, version.GetVersion())
		fmt.Println(string(version.VERSION))
	})
}
