package loadbalancer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//Balancer settings from settings.json
type balancer struct {
	Hosts                  []string    // List of hosts to check and rotate through
	Port                   string      `json:"port"`                   // Port to serve the proxy on
	HostCheckTimeout       int         `json:"hostCheckTimeout"`       // Timeout for the checkHost function to determine if host is online
	PeriodicHostCheckDelay int         `json:"periodicHostCheckDelay"` // Delay to schedule the UpdateAliveHosts function to probe all hosts for uptime
	Client                 http.Client // Instance of modified http.Client with settings applied from settings.json
	nextHost               int         // the next rotation used to select the next host
	hostCount              int         // Total number of hosts provided
	aliveHosts             []string    // Modified list of hosts that only contains ones acceptable for rotation
}

//Bal : balancer object
var Bal balancer

//ApplySettings : apply settings to Bal from settings.json
func ApplySettings() {
	settingsFile, err := os.Open("./settings.json")

	if err != nil {
		fmt.Println(err)
	}
	defer settingsFile.Close()

	fileBytes, _ := ioutil.ReadAll(settingsFile)
	json.Unmarshal([]byte(fileBytes), &Bal)

	timeout := time.Duration(Bal.HostCheckTimeout)
	Bal.Client = http.Client{Timeout: time.Duration(timeout * time.Second)}
	Bal.hostCount = len(Bal.Hosts)
	Bal.nextHost = 0
	Bal.aliveHosts = Bal.Hosts
}
