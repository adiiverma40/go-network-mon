package network

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)



func PingCloudFlare() (float64 , error) {

	fmt.Println("Pinning cloudflare")
	ip := "1.1.1.1"
	var cmd *exec.Cmd

	if runtime.GOOS == "windows"{
		fmt.Errorf("Windows is not supported. Use ubuntu env inside windows using wsl")

	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "2", ip)
	}
	output, err := cmd.CombinedOutput()
	if err != nil{
		return 0, fmt.Errorf("ping to cloudflare %s failed: %v", ip , err)
	}
	re := regexp.MustCompile(`time[=<]([0-9\.]+)\s*ms`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
			return 0, fmt.Errorf("could not parse latency from ping output: %s", string(output))
		}

	latency, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert latency to number: %v", err)
	}

	return latency, nil
}
