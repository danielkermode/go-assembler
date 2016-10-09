package convert

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
)

const aCommand = "A"
const cCommand = "C"

// Convert ...
func Convert(line string, number int) (string, error) {
	ct := commandType(line)
	switch ct {
	case aCommand:
		return intToBinary(symbol(line)), nil
	case cCommand:
		res := getMems(line)
		if len(res) == 0 {
			return "", errors.New("Can't parse symbol(s)")
		}
		return "111" + res, nil
	//... etc
	default:
		return "", errors.New("Can't determine command")
	}
}

func commandType(line string) string {
	if string(line[0]) == "@" {
		return aCommand
	}
	return cCommand
}

func symbol(command string) string {
	return string(command[1:])
}

func intToBinary(n string) string {
	i, err := strconv.ParseInt(n, 10, 64)
	Check(err)
	str := strconv.FormatInt(i, 2)
	var buf bytes.Buffer

	for i := 0; i < 16-len(str); i++ {
		buf.WriteString("0")
	}
	buf.WriteString(str)

	return buf.String()
}

func getMems(s string) string {
	dest, comp, jump := split(s)
	var c CMap = CCommand{dest, comp, jump}
	return c.getComp() + c.getDest() + c.getJump()
}

func split(s string) (string, string, string) {
	splitter := regexp.MustCompile("(?:(\\w+)=)?([!&|\\w+-]+)(?:;(\\w+))?")
	res := splitter.FindStringSubmatch(s)
	return res[1], res[2], res[3]
}
