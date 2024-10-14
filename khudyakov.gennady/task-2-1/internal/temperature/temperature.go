package temperature

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/KRYST4L614/task-2-1/internal/io"
)

type TemperatureConstraints struct {
	Min int
	Max int
}

func TemperatureRequestHandler(constraints TemperatureConstraints, reader *bufio.Reader, writer *bufio.Writer) error {
	currentMax := constraints.Max
	currentMin := constraints.Min

	k, err := io.ReadInt(reader)
	if err != nil {
		return err
	}

	for range k {
		term, temperature, err := ReadTerm(reader)
		if err != nil {
			return err
		}

		if temperature < 15 || temperature > 30 {
			return temperOutOfRangeErr(constraints)
		}

		switch term {
		case ">=":
			currentMin = temperature
		case "<=":
			currentMax = temperature
		default:
			return temperTermFormatErr(">=", "<=")
		}
		if currentMin > currentMax {
			writer.WriteString(fmt.Sprintln(-1))
			writer.Flush()
			return nil
		}
		writer.WriteString(fmt.Sprintln(currentMin))
		writer.Flush()
	}
	return nil
}

func ReadTerm(reader *bufio.Reader) (string, int, error) {
	inputString, err := reader.ReadString('\n')
	if err != nil {
		return "", 0, err
	}

	split := strings.Split(strings.TrimSpace(inputString), " ")
	if len(split) != 2 {
		return "", 0, temperFormatErr(inputString)
	}
	temperature, err := strconv.ParseInt(split[1], 10, 0)
	if err != nil {
		return "", 0, io.ParseIntError(split[1])
	}
	return split[0], int(temperature), nil
}