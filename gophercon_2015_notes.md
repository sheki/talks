#introduce the talk and yourself.

#what is this talk about.

#what is parse.com?
- backend as a service.
- we have Parse Core, Parse Push and Parse Analytics along with crash reporting, configuration managment etc.

#parse.com issues
- Parse had several problems affecting our uptime.
- Unicorn our ruby webserver worked with a limited number of workers. When our traffic spiked we would very quickly run out of unicorns. Each host could only handle limited number of requests.- Our deploy times were high. It would take us 25min to deploy, we could not manually respond to traffic spikes.
- The ruby codebase was unmanageable, hard to write unit tests against and full of optional parameters.

#why rewrite
* We decided to rewrite.
* the estimated performance wins of the rewrite were high.
* It was hard to understand the ruby codebase, we had unfortunately used a lot of gems, it was heard to understand what was going on.
We lost a few engineers who built the initial ruby stack.
* Reading a dynamic language is hard.
* Our bursty nature of traffic was not very well suited for ruby/unicorn.

#Why Go?
- Statically typed.
- Good concurrency support
- Easier to hire programmers then c++
- We tried running stuff on Java and we had trouble as most of our jems were not thread safe.
- We liked C# but it was not open.

#Rules of the rewrite
- Being a developer platform we did not want to break any backward compatibility.
- we wanted to do it live.
- my manager called it changing the engine of a running car.

#Progress of the rewrite
* we started with a new service parse hosting. Its a new service, needed to things which ruby was not particularly good with -> quick restarts.
* this was a great success we built a reliable service in quick time.
* we next went after our PPNS service. This is a service which opens a lot of network connections to push providers like Apple and Google. Our event-machine based ruby service was breaking at the seams. We rewrote this in go and very happy with the reliability and uptime.
* This convinced us that our biggest beast the API server should be in go and we earnestly started rewriting our stack.

#rewrite contd.
* we picked low traffic read endpoints and slowly moved on to read and write parts.
* through out the process we ran traffic through a shadow cluster and compared results with our prod ruby cluster.
* we found out that a lot of unexpected behaviour supported by our ruby stack and we had to go implement it as our customers relied on it.
* One example of this is how ruby would represent arrays in HTTP parameters.
* this process went on and we started making progress.

#go a young language.
- go was relatively young language - no best practices we had to build our own :)
- we had to build lots of libraries. it was not the gem world where we almost had a gem for everything.
- we had to write/maintain our own version of the cassandra library after it stopped being maintained.
- go did not have a good story around Stopping an HTTP server till 1.3.
- we had to build all of these.
- so we built a lot of libraries and tools.

#dependency injection problems
- as our go codebase started to grow. we needed to pass a lot of mock implementations of services in test.
- writing the code to build dependencies in each package was cumbersome.
- all our mocks in tests were globals and we could not write parallel tests because of that.
- also we needed to pass in components in a top down fashion and a missed component would cause a service to fail in production.
- this was a repeated pattern we saw at parse.

#start-stop
- the next problem we ran into was starting and stopping services in the order of the dependency.
- we needed to start the lowest component service before we started the upper layers. Doing this manually caused a few crashes. This was pretty uninteresting stuff to write too.

- we built a library to do this. its called start-stop

#error-reporting
- next problem we ran into was tracking errors and the components causing them
- we started adding stack traces to all our errors. and aggregated them in an in house called log view. This is great for us to track errors.

#proxy.
