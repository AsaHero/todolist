package goenv

import (
	"bufio"
	"os"
	"strings"
)

type Variables struct {
	variable map[string]string
}

func Load(path string) (*Variables, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	vr := new(Variables)
	vr.variable = make(map[string]string, 1)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		splited_str := strings.Split(str, "=")
		vr.variable[splited_str[0]] = splited_str[1]
	}
	return vr, nil
}

func (v *Variables) Getenv(key string) string {
	return v.variable[key]
}
