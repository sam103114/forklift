# Forklift
A relatively simple http load-balancer written in Go.

Rotates through hosts and distributes requests based on uptime.
## Details
Forklift will probe all hosts to assemble an up-to-date list of available hosts for rotation periodically based on the "periodicHostCheckDelay" option.

The hosts assembled from this list are still checked before returning the next best host for reliability purposes, but this recovery is much slower (~1.3 second greater latency) than using the predetermined list. So still make sure to set the periodicHostCheckDelay to an appropriate value that gives frequent enough updates to the proxy.
