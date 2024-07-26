package linux

import (
	"fmt"
	"testing"
)

func Test_MachineID(t *testing.T) {
	fmt.Println(MachineID())
}
func Test_HostNameAndkernelVersion(t *testing.T) {
	hn, kv, err := HostNameAndkernelVersion()
	fmt.Println(hn, kv, err)
}
