package format

import (
	"fmt"
	"regexp"
	"strconv"
)

const DEFAULT_VERSION float32 = 0.5

const version string = "0.5"
const AES_S4 string = "AES+S4"
const S4 string = "S4"

var spaceRegex = regexp.MustCompile(`\s+|={3,}`)

var travelFormatRegex = regexp.MustCompile(`(?msi)(?:BEGIN)?\[s4v(\d+\.\d+)\|*(AES\+S4|S4)\]?(.*?)(?:(?:END|\[).*?)?\]`)

// IsTravelValidFormat checks if the provided format is valid based on rules.
func IsTravelValidFormat(travelFormat string) bool {
	noSpaces := spaceRegex.ReplaceAllString(travelFormat, "")
	return travelFormatRegex.MatchString(noSpaces)
}

type Format struct {
	UseAES                    bool
	Version                   float32
	Data                      string
	OptimizedHumandReadbility bool
}

func CreateTravelFormat(f Format) string {
	opt := AES_S4
	if !f.UseAES {
		opt = S4
	}

	if !f.OptimizedHumandReadbility {
		return fmt.Sprintf("[s4 v%g %s %s]",
			f.Version, opt, f.Data)
	}

	dataInLines := "  "
	countInLine := 0
	for _, char := range f.Data {
		if countInLine == 20 {
			dataInLines += "\n  "
			countInLine = 0
		}
		if countInLine != 0 && countInLine%4 == 0 {
			dataInLines += " "
		}
		dataInLines += string(char)
		countInLine++
	}

	return fmt.Sprintf("[s4 v%g %s\n%s\n]",
		f.Version, opt, dataInLines)
}

func ParseTravelFormat(travelFormat string) (format *Format, err error) {
	noSpaces := spaceRegex.ReplaceAllString(travelFormat, "")

	// Parse via travelFormatRegex
	res := travelFormatRegex.FindStringSubmatch(noSpaces)
	if len(res) != 4 {
		return nil, fmt.Errorf("did not get 3 capture groups. Format broken")
	}

	// Parse float from string
	versionFloat, err := strconv.ParseFloat(res[1], 32)
	if err != nil {
		return nil, fmt.Errorf("could not parse version float: %w", err)
	}

	return &Format{
		UseAES:  res[2] == AES_S4,
		Version: float32(versionFloat),
		Data:    res[3],
	}, nil
}
