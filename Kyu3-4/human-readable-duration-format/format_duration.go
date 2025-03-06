package kata

import (
	"fmt"
	"strings"
)

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	sliceOfReadableString := []string{}

	// Days = seconds / 86400
	// seconds - (86400 * days)
	days := seconds / 86400
	daysString := new(string)
	yearsString := new(string)
	if days != 0 {
		seconds -= (86400 * days)
		if days > 364 {
			years := days / 365

			if years > 1 {
				*yearsString = fmt.Sprintf("%v years", years)
				sliceOfReadableString = append(sliceOfReadableString, *yearsString)
			} else {
				*yearsString = fmt.Sprintf("%v year", years)
				sliceOfReadableString = append(sliceOfReadableString, *yearsString)
			}

			remainderDays := days % 365
			if days > 364 {
				if remainderDays != 0 {
					if remainderDays > 1 {
						*daysString = fmt.Sprintf("%v days", remainderDays)
						sliceOfReadableString = append(sliceOfReadableString, *daysString)
					} else {
						*daysString = fmt.Sprintf("%v day", remainderDays)
						sliceOfReadableString = append(sliceOfReadableString, *daysString)
					}
				}
			}
		} else {
			if days > 1 {
				*daysString = fmt.Sprintf("%v days", days)
				sliceOfReadableString = append(sliceOfReadableString, *daysString)
			} else {
				*daysString = fmt.Sprintf("%v day", days)
				sliceOfReadableString = append(sliceOfReadableString, *daysString)
			}
		}
	}

	// Hours = seconds / 3600
	// seconds - (3600 * hours)
	hours := seconds / 3600
	hoursString := new(string)
	if hours != 0 {
		seconds -= (3600 * hours)
		if hours > 1 {
			*hoursString = fmt.Sprintf("%v hours", hours)
			sliceOfReadableString = append(sliceOfReadableString, *hoursString)
		} else {
			*hoursString = fmt.Sprintf("%v hour", hours)
			sliceOfReadableString = append(sliceOfReadableString, *hoursString)
		}
	}

	// Minutes = seconds / 60
	// seconds - (60 * minutes)
	minutes := seconds / 60
	minutesString := new(string)
	if minutes != 0 {
		seconds -= (60 * minutes)
		if minutes > 1 {
			*minutesString = fmt.Sprintf("%v minutes", minutes)
			sliceOfReadableString = append(sliceOfReadableString, *minutesString)
		} else {
			*minutesString = fmt.Sprintf("%v minute", minutes)
			sliceOfReadableString = append(sliceOfReadableString, *minutesString)
		}
	}

	// Seconds = seconds % 60
	secondsString := new(string)
	if seconds != 0 {
		if seconds > 1 {
			*secondsString = fmt.Sprintf("%v seconds", seconds%60)
			sliceOfReadableString = append(sliceOfReadableString, *secondsString)
		} else {
			*secondsString = fmt.Sprintf("%v second", seconds%60)
			sliceOfReadableString = append(sliceOfReadableString, *secondsString)
		}
	}

	n := len(sliceOfReadableString)
	if n == 1 {
		return sliceOfReadableString[0]
	} else if n == 2 {
		return sliceOfReadableString[0] + " and " + sliceOfReadableString[1]
	} else {
		return strings.Join(sliceOfReadableString[:n-1], ", ") + " and " + sliceOfReadableString[n-1]
	}
}
