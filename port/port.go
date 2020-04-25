package port

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const rangeSep = "-"

// GetPortForwardArgs build string as param to ssh for port forward
//
// sample input: 3000 4040:40040 8080-8088:9080-9088
//
// where <vm port>[:<local port>]
//
// remember, vm port comes first
func GetPortForwardArgs(ports string) ([]string, error) {

	args := strings.Fields(ports)
	var result []string

	for _, arg := range args {
		mapping := strings.Split(arg, ":")
		vmPort := mapping[0]
		var localPort string
		if len(mapping) == 1 {
			localPort = vmPort
		} else {
			localPort = mapping[1]
		}

		portMap, err := getPortMapping(vmPort, localPort)
		if err != nil {
			return nil, err
		}

		for vm, local := range portMap {
			result = append(result, "0.0.0.0:"+local+":localhost:"+vm)
		}
	}

	return result, nil
}

func getPortMapping(vmPort string, localPort string) (map[string]string, error) {
	err := checkRangeFormat(vmPort, localPort)
	if err != nil {
		return nil, err
	}
	return doPortMapping(vmPort, localPort)
}

func doPortMapping(vmPort string, localPort string) (map[string]string, error) {

	if strings.Contains(vmPort, rangeSep) {
		vmPorts := strings.Split(vmPort, rangeSep)
		localPorts := strings.Split(localPort, rangeSep)

		firstVmPort, _ := strconv.Atoi(vmPorts[0])
		lastVmPort, _ := strconv.Atoi(vmPorts[1])

		firstLocalPort, _ := strconv.Atoi(localPorts[0])
		lastLocalPort, _ := strconv.Atoi(localPorts[1])

		if lastVmPort-firstVmPort != lastLocalPort-firstLocalPort {
			return nil, errors.New(fmt.Sprintf("number of ports not matched %s %s. %d ports vs %d ports",
				vmPort, localPort, lastVmPort-firstVmPort, lastLocalPort-firstLocalPort))
		}

		var ret = make(map[string]string)

		for i := 0; i <= lastVmPort-firstVmPort; i++ {
			ret[strconv.Itoa(firstVmPort+i)] = strconv.Itoa(firstLocalPort + i)
		}
		return ret, nil
	} else {
		return map[string]string{vmPort: localPort}, nil
	}
}

func checkRangeFormat(vmPort string, localPort string) error {
	if strings.Contains(vmPort, rangeSep) && !strings.Contains(localPort, rangeSep) {
		return errors.New(fmt.Sprintf("range ports not matched %s %s", vmPort, localPort))
	}
	return nil
}
