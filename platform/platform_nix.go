//go:build linux || darwin
// +build linux darwin

package platform

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// GetArchInfo() returns basic host architecture information
func GetArchInfo() (archInfo map[string]string, err error) {
	archInfo = map[string]string{}

	out, err := exec.Command("uname", unameOptions...).Output()
	if err != nil {
		return nil, err
	}
	line := fmt.Sprintf("%s", out)
	values := regexp.MustCompile(" +").Split(line, 7)
	updateArchInfo(archInfo, values)

	out, err = exec.Command("uname", "-v").Output()
	if err != nil {
		return nil, err
	}
	archInfo["kernel_version"] = strings.Trim(string(out), "\n")

	osReleasePath := "/etc"
	if os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		osReleasePath = "/host/etc"
	}

	archInfo["release_version"] = getOsFromOsReleaseFile(osReleasePath)

	return
}

func getOsFromOsReleaseFile(osReleasePath string) string {
	release_version := ""
	bytesRead, err := ioutil.ReadFile(fmt.Sprintf("%s/os-release", osReleasePath))
	if err != nil {
		fmt.Println("could not read os-release file")
		return ""
	}
	regExp := regexp.MustCompile(`PRETTY_NAME="(.*)"`)
	result := regExp.FindAllStringSubmatch(string(bytesRead), -1)
	if len(result) == 1 { //result=[PRETTY_NAME="CentOS Linux 7 (Core)" CentOS Linux 7 (Core)]
		s := result[0]
		if len(s) == 2 {
			release_version = s[1]
		}
	}
	return release_version
}
