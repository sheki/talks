#introduce the talk and yourself.
Hi every one. I hope you are having a good time at Gopher con. The next talk is titled Rewriting Parse.com in Go.
I am Abhishek Kona a Software Engineer at Parse, Facebook.

#what is this talk about.
We at Parse rewrote our API stack from ruby to Go. I
will talk about why we did it, how we did it. I will walkthrough a few libraries we built in the process.

#what is parse?
What is Parse? Parse is a Developer platform to build mobile apps. It is Backend-As-A-Service. If you are a mobile app developer you do not maintain any servers, you use parse. We have three main components - Parse Core to manage data, Parse Push to send notifications and Parse Analytics to track app metrics.

We support IoS, Android, PHP and a lot of other platforms.

Parse was acquired by Facebook in 2013.

# Parse - before Go.
How did Parse look before adopting Go. In 2013, parse was a Ruby on Rails app, it had 10 engineers working on it. It was quite popular we had around 60 thousand mobile apps. But we were starting to have issues.

#parse issues
Our biggest issue at Parse was one big app could impact our performance. This was because we used Unicorn - a process based ruby HTTP server.

Unicorn had a fixed number of workers on each api server. Under traffic spikes we would quickly run out of workers. It would happen too quickly for Auto Scaling groups to react.

Our ruby deploy process was slow - it took around 25 minutes to deploy. This meant changes or urgent fixes could still not be deployed quickly.

We were not in a great situation.

#why rewrite
So against popular wisdom - We decided to rewrite.

#why rewrite.
* Estimated performance and reliability wins of using go were huge.
* It was hard to understand the ruby codebase, we had unfortunately used a lot of gems, it was heard to understand what was going on.
We lost a few engineers who built the initial ruby stack.
* Reading a dynamic language is hard.
* Our stack did not look well suited for a 10x growth.
* Our bursty nature of traffic was not very well suited for ruby/unicorn.

#Why Go?
- Statically typed.
- Good concurrency support
- Easier to hire programmers then c++
- We tried running stuff on Java and we had trouble as most of our gems were not thread safe.
- We liked C# but it was not open.

#Rules of the rewrite
- Being a developer platform we did not want to break any backward compatibility.
- we wanted to do it live. Mind you doing it live was not easy. We were growing the number of customers at
a rapid pace. We were increasing the number of backends we were adding to our stack. So we were chasing a moving target.
- my manager called it changing the engine of a running car.

#Progress of the rewrite
So how did we go about the rewrite ->
Actually we started with a new service parse hosting. It was a new product, it needed to do things which ruby was not particularly good with -> quick restarts.
So we decided to give Go a try.
* this was a great success we built a reliable service in quick time.
* we next went after our PPNS service. This is a service which opens a lot of network connections to push providers like Apple and Google. Our event-machine based ruby service was breaking at the seams. We rewrote this in go and very happy with the reliability and uptime.
* This convinced us that our biggest beast the API server should be in go and we earnestly started rewriting our stack.

#rewrite contd.
* So how did we go about rewriting the API server.
* Parse has around 50 api end points.
* We picked low traffic read endpoints and slowly moved on to write end points.
* Through out the process we ran traffic through a shadow cluster and compared results with our prod ruby cluster.
* we found out a lot of unexpected behaviour supported by our ruby stack and had to implement it in go.
* One example of this is how ruby would represent arrays in HTTP parameters.
* this process went on and we started making progress.

# Some Comments in our Ruby Codebase.
.code code/ruby_comments.go

#go a young language.
- go was relatively young language
- we had to build lots of libraries. it was not the gem world where we almost had a gem for everything.
- we had to write/maintain our own version of the cassandra library after it stopped being maintained.
- go did not have a good story around Stopping an HTTP server till 1.3.
- we had to build all of these.
- in the next section I will talk about the libraries and tools we had to build in go.

#dependency injection problems
- as our go codebase started to grow. we needed to pass a lot of mock implementations of services in test.
- writing the code to build dependencies in each package was cumbersome.
- all our mocks in tests were globals and we could not write parallel tests because of that.
- also we needed to pass in components in a top down fashion and a missed component would cause a service to fail in production.
- this was a repeated pattern we saw at parse.

#DI code
A goode candiadate for inject is a struct which has only one instance through the lifecycle of the program. In this Example we provide the dependencies of a Haandler struct via inject.

#start-stop
- the next problem we ran into was starting and stopping services in the order of the dependency.
- we needed to start the lowest component service before we started the upper layers. Doing this manually caused a few crashes. This was pretty uninteresting code to write everytime.

- we built a library to do this. its called start-stop

#graceful restarts
- We had issues restarting our servers during deploys. We started with dropping a few user requests during deploys. We wanted to do better than this. Getting a load balancer involved would slow things down, so we built a library called grace. Grace restarts binaries and hands off the socket from the old process to the new process and waits for the old process to finish servicing its current requests before shut down.

#error-reporting
- next problem we ran into was tracking errors and the components causing them
- we started adding stack traces to all our errors. and aggregated them in an in house called log view. This is great for us to track errors.

#stackerr code
We wrap every call site where we return an error with a stackerror Wrap call. This attaches the stack at the point to the call site.
#proxy.
Our primary Database mongo has a hard limit on the number of database connections it can handle - 20000. As we started to add more services and apps we started hitting this limit. So we built a proxy for Mongo purely in Go. Go's IO shone through when we were building this. MongoProxy is called Dvara and is available on our github repo. So building in a proxy in Go not that hard.

# Muster

Another common pattern we saw through our codebase was batching operations together and fire them in a batch later. For example we wanted to batch our metrics send to our metrics collection system and fire one request instead of multiple. We saw the same pattern in our Billing Logger. So we built a library to move the common part of batching operations in a channel and firing them off when a TimeLimit or a BatchLimit is reached. Muster again is open-source and available at our github repo.
