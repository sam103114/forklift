package loadbalancer

import (
	"time"
)

//Checks if host responds to request
func checkHost(url string) bool {
	_, err := Bal.Client.Get(url)

	if err != nil {
		return false
	}
	return true
}

//UpdateAliveHosts : Called periodically to get a refined list of acceptable hosts
func UpdateAliveHosts(ticker <-chan time.Time) {
	for t := range ticker {
		_ = t
		aliveHosts := []string{}

		for _, host := range Bal.Hosts {
			if checkHost(host) {
				aliveHosts = append(aliveHosts, host)
			}
		}

		Bal.aliveHosts = aliveHosts
	}

}

//Rotate through alive hosts, select one that matches the current rotation
func selectRedirectURL() string {
	if Bal.nextHost >= len(Bal.aliveHosts) {
		Bal.nextHost = 0
	}

	for index, host := range Bal.aliveHosts {
		if index == Bal.nextHost {
			if checkHost(host) {
				Bal.nextHost++
				return host
			}
			Bal.nextHost++
		}
	}

	return ""
}
