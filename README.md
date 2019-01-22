![](https://github.com/sam103114/forklift/blob/master/Forklift.png)
# Forklift
A multi-purpose reverse proxy written in Go.

## Current Features
### Load-balancer
Rotates through the list of hosts and selects the best host based on uptime.
### Details
Forklift will send out a simple get requests to all hosts in the array of hosts in settings.json to gather an updated list of selectable hosts. The frequency of this check is based on the PeriodicHostCheckDelay option. If the host matches the current rotation, it is selected as the next one to provide to the next client. Before a host from the list of alive hosts is selected, it is still checked once again for reliability. Although this may seem redundant, it prevents uneven distribution among hosts and makes sure a client doesn't get redirected to a dead host that might've went out in-between the intervals of checks.
