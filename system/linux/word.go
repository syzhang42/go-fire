package linux

import (
	"bytes"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/syzhang42/go-fire/log"
	"golang.org/x/sys/unix"
)

// 你可以在这里获取linux的一些信息，包括如下：

// 机器唯一id
func MachineID() string {
	for _, p := range []string{"sys/devices/virtual/dmi/id/product_uuid", "etc/machine-id", "var/lib/dbus/machine-id"} {
		payload, err := os.ReadFile(path.Join("/proc/1/root", p))
		if err != nil {
			continue
		}
		id := strings.TrimSpace(strings.Replace(string(payload), "-", "", -1))
		return id
	}
	return ""
}

// hostname kernelVersion
func HostNameAndkernelVersion() (string, string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	f, err := os.Open("/proc/1/ns/uts")
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	self, err := os.Open("/proc/self/ns/uts")
	if err != nil {
		return "", "", err
	}
	defer self.Close()

	defer func() {
		unix.Setns(int(self.Fd()), unix.CLONE_NEWUTS)
	}()

	err = unix.Setns(int(f.Fd()), unix.CLONE_NEWUTS)
	if err != nil {
		return "", "", err
	}
	var utsname unix.Utsname
	if err := unix.Uname(&utsname); err != nil {
		return "", "", err
	}
	hostname := string(bytes.Split(utsname.Nodename[:], []byte{0})[0])
	kernelVersion := string(bytes.Split(utsname.Release[:], []byte{0})[0])
	return hostname, kernelVersion, nil
}

// 系统uuid
func SystemUUID() string {
	payload, err := os.ReadFile(path.Join("/proc/1/root", "/sys/devices/virtual/dmi/id/product_uuid"))
	if err != nil {
		log.Error("failed to read system-uuid:", err)
		return ""
	}
	return strings.TrimSpace(string(payload))
}
