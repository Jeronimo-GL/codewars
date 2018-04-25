package coordinates

import (
	"strings"
	"strconv"
	"regexp"
	"errors"
)


func parseCoord(coord string) (float64, error) {
	re :=regexp.MustCompile("\\A[+-]?([0-9]*[.])?[0-9]+\\z")
	if re.MatchString(coord) {
		numb, err := strconv.ParseFloat(coord, 64)
		if err != nil {
			return 0.0, err
		}
		return numb, nil
	}else{
		return 0, errors.New("Wrong format")
	}
}

func IsValidCoordinates(coordinates string) bool {
	parts := strings.Split(coordinates, ",")
	if len(parts) == 2 {
		latitude, err := parseCoord(strings.TrimSpace(parts[0]))
		if (err != nil || latitude > 90.0 || latitude < -90.0) {
			return false
		}
		longitude, err := parseCoord(strings.TrimSpace(parts[1]))
		if (err != nil || longitude > 180.0 || longitude < -180.0) {
			return false
		}
		return true
	} else {
		return false
	}
}
