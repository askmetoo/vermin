package info

import (
	"github.com/mhewedy/vermin/command"
	"strings"
)

func Get(vmName string) ([]string, error) {
	out, err := command.VBoxManage("showvminfo", vmName, "--machinereadable").Call()
	if err != nil {
		return nil, err
	}
	return strings.Fields(out), nil
}

func FindByPrefix(vmName string, prefix string) ([]string, error) {

	entries, err := Get(vmName)
	if err != nil {
		return nil, err
	}

	var values []string

	for i := range entries {
		entry := entries[i]
		if strings.HasPrefix(entry, prefix) {
			values = append(values, strings.Split(entry, "=")[1])
		}
	}

	return values, nil
}

func FindFirstByPrefix(vmName string, prefix string) (string, bool, error) {
	byPrefix, err := FindByPrefix(vmName, prefix)
	if err != nil {
		return "", false, err
	}

	if len(byPrefix) == 0 {
		return "", false, nil
	}

	return byPrefix[0], true, nil
}