package gateway

import (
	"net"
	"os/exec"
)

func DiscoverGateway() (ip net.IP, err error) {
	ip, err = discoverGatewayUsingRoute()
	if err != nil {
		ip, err = discoverGatewayUsingIpRouteShow()
	}
	if err != nil {
		ip, err = discoverGatewayUsingIpRouteGet()
	}
	return
}

func discoverGatewayUsingIpRouteShow() (net.IP, error) {
	routeCmd := exec.Command("ip", "route", "show")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return parseLinuxIPRouteShow(output)
}

func discoverGatewayUsingIpRouteGet() (net.IP, error) {
	routeCmd := exec.Command("ip", "route", "get", "8.8.8.8")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return parseLinuxIPRouteGet(output)
}

func discoverGatewayUsingRoute() (net.IP, error) {
	routeCmd := exec.Command("route", "-n")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return parseLinuxRoute(output)
}
