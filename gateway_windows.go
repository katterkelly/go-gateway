package gateway

import (
	"net"
	"os/exec"
	"syscall"
)

func DiscoverGateway() (ip net.IP, ip1 net.IP, err error) {
	routeCmd := exec.Command("route", "-4", "print", "0.0.0.0")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, nil, err
	}

	return parseWindowsRoutePrint(output)
}
