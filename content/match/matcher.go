package matcher

import (
	"regexp"
	"strconv"

	"ocelot/content/types"
)

func SeriesData(fullPath string) (types.Show, error) {
	var show types.Show

	seasonString, err := regexp.Compile(`(S[0-9]{2,})|(Season\s[0-9]{1,})`)
	num, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		return show, err
	}

	show.SeasonNumber, err = strconv.Atoi(num.FindString(seasonString.String()))
	if err != nil {
		return show, err
	}

	return show, nil
}
