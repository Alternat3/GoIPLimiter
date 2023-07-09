/*
Rate limiter based on number of requests, time and IP address (port is stripped from inputted IP. If no port IPv6 addresses will be wrong. IPv4 OK.)
-Liam Smith
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

// Limiter struct holds all objects of IP Limiter where IP, number of requests is held. Limiter time is input with Time library.
type Limiter struct {
	ips        map[string]*IPLimiter
	rate       int
	limterTime time.Duration
}

type IPLimiter struct {
	numOfReqs    int
	timeLastWipe time.Time
}

// Creates new limiter to hold all IP limiters. Called when creating limiter in request handling function.
func newLimiter(reqPerMin int, limiterTime time.Duration) *Limiter {
	limiter := &Limiter{
		ips:  make(map[string]*IPLimiter),
		rate: reqPerMin,
	}

	return limiter
}

// Creates new IP limiter for inputted IP. Only called when given IP doesn't already exist. Hence reqs initalized to 1.
func (i *Limiter) AddIP(ip string) *IPLimiter {
	l := &IPLimiter{
		numOfReqs:    1,
		timeLastWipe: time.Now(),
	}

	i.ips[ip] = l

	return l
}

// Checks if IP limiter already exists for given IP after stripping port. If != exist, creates new IP Limiter.
func (i *Limiter) GetLimiter(ip string) *IPLimiter {
	splitIP := strings.Split(ip, ":")
	ip = strings.Join(splitIP[:len(splitIP)-1], ":") //Port stripping.
	limiter, exists := i.ips[ip]
	fmt.Println(exists)

	if !exists {
		return i.AddIP(ip)
	}

	return limiter
}

/*Middleware code for request handler + limiter declartion.

var ipLimiter = newLimiter(10, time.Minute)

func limiterMiddleWare(w http.ResponseWriter, r *http.Request) {
	limiter := ipLimiter.GetLimiter(r.RemoteAddr)

	fmt.Println("IP: ", r.RemoteAddr)
	fmt.Println("Requests: ", limiter.numOfReqs)
	fmt.Print("Time since last wipe: ", limiter.timeLastWipe, "\n")

	if time.Since(limiter.timeLastWipe) > ipLimiter.limterTime {
		limiter.timeLastWipe = time.Now()
		limiter.numOfReqs = 1
	} else {
		limiter.numOfReqs = limiter.numOfReqs + 1
		if limiter.numOfReqs > ipLimiter.rate {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
	}
}
*/
