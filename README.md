![](https://github.com/sam103114/forklift/blob/master/Forklift.png)
# Forklift
A multi-purpose reverse proxy written in Go.
## Current Features
### Load-balancer
Forklift will probe all hosts to assemble an up-to-date list of available hosts for rotation periodically based on the "periodicHostCheckDelay" option.

The hosts assembled from this list are still checked before returning the next best host for reliability purposes, but this recovery is much slower (~1.3 second greater latency) than using the predetermined list. So still make sure to set the periodicHostCheckDelay to an appropriate value that gives frequent enough updates to the proxy.
