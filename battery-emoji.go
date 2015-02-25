package main

import "os/exec"
import "strings"
import "strconv"
import "fmt"

func main() {
	//	Shell out to 'ioreg', Apple's I/O Kit registry tool,
	//	to fetch properties of the AppleSmartBattery object
	cmd := exec.Command("ioreg", "-rc", "AppleSmartBattery")
	out, err := cmd.Output()
	data := strings.Split(string(out), "\n")

	if err != nil {
		println(err.Error())
		return
	}

	minutesRemaining := 0

	// Assume that the battery will never take longer
	// than 50 hours to charge or discharge
	minutesUpperLimit := 60 * 50

	charging := false
	pluggedIn := false

	for _, line := range data {
		if strings.Contains(line, "TimeRemaining") {
			minutesRemaining, _ = strconv.Atoi(strings.Split(line, " = ")[1])
		} else if strings.Contains(line, "IsCharging") {
			charging = strings.Split(line, " = ")[1] == "Yes"
		} else if strings.Contains(line, "ExternalConnected") {
			pluggedIn = strings.Split(line, " = ")[1] == "Yes"
		}
	}

	// We could use an emoji library to have access to every emoji
	// on most platforms (like https://github.com/kyokomi/emoji)
	// but this keeps the code faster and dependency-free.
	exclamationEmoji := "❕"
	questionMarkEmoji := "❔"
	plugEmoji := "🔌"
	batteryEmoji := "🔋"

	timeString := questionMarkEmoji
	if minutesRemaining < minutesUpperLimit {
		if minutesRemaining >= 60 {
			timeString = fmt.Sprintf("%dh%dm", int(minutesRemaining/60), minutesRemaining%60)
		} else {
			timeString = fmt.Sprintf("%dm", minutesRemaining)
		}
	}

	if pluggedIn {
		if charging {
			fmt.Printf("%s %s", plugEmoji, timeString)
		} else {
			fmt.Printf("%s ", plugEmoji)
		}
	} else {
		if minutesRemaining < 10 {
			fmt.Printf("%s%s %s", exclamationEmoji, batteryEmoji, timeString)
		} else {
			fmt.Printf("%s %s", batteryEmoji, timeString)
		}
	}
}
