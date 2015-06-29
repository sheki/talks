api.parse.com slide 2
how why and tools and libraries.

What is parse? slide3
IoS to IOS
Just have React
Have a narrative around what we offer / storage - analytics - push

Parse - before Go. `
#problems parse had
Fixed unicorn pool.
Explain the unicorn problem with network time and thread not blocking.
Go - async model / network model._
Deploy time ~25mindeploy + ~25test run time.

Parse after go as the last slide.

#Why Rewrite
* Explain why ruby codebase is not readable
* Lost early engineer
* Rails had a lot of gems
* Estimated reliablity win was huge. Higher throughput
* No static.
* Bursty nature of traffic.

#What other options.
* JRUBY ?
* C# / C++

#Why GO.
* Statically typed.
* Easy to hire engineers.
* More engineers wanted to write Go > C# > C++

# Rules of the rewrite.
* Not break backward compatibility
* No downtime.

# Status of Rewrite  => Progression
* started with hosting
* get rid of event machine and built the PPNS service.
* went from 250k conns per node to 1.5m conns per node.
* Progressed to our beast => api.parse.com
* How did we pick endpoints
* talk about RollOut.

# Go the good parts -> SKIP (say it at the last).

# recap slide
* dropped capacity form ~1000nodes to ~100nodes

# pivot to libraries.
* Go was young
* We have- X number of libraries on facebookgo
* Highlight the ones.
* Cassandra was young

# DI
* lead with a problem
* Globals on tests and tests would run slowly.
* Parallel.
* We rely heavily on mongo and it is difficult.o
* Skip named inject.
* make all of them interfaces talk about how we provde concrete in prod mocked in test.

#Main for inject enabled service.go
* lot of repeated code
* starting and stopping services in the dependency order is error-prone

# startstop instead of ship
* graceful stopping services.
* interfaces implementing Start/Stop. (DONT SAY OPEN CLOSE)
* MAKE IT HTTP NOT THRIFT

# last-1 slide on popular libs we use
* mgo

# dont do harness. (after you check for time bring it back)

# Grac
# Testing
* maybe kill skip
* mgotest/mcskip

