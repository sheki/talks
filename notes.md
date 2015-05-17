Welcome every one to RootConf 2015.
I have never opened the batting before, usually come down the order.

I am Abhishek Kona, I am Software Engineer at parse.com, facebook.
Before that I worked on RocksDB at facebook, I was at Flipkart previously.

This talk is roughly arranged into three parts, it is more or less a 101 for backend team.

#PARSE CIRCA 2013.
60K apps
10 engineers
Ruby on Rails App like YC
3000 Req/s

#Beast list
* Unicorn our Ruby HTTP server was a resource hog.
  It is a forking process based webserver. Under
  traffic pressure we had to add new nodes, but not
  quickly enough thanks to our slow ruby build deploy  process.

#We decided to rewrite in GO.
Based on our beast list one of the
to rewrite our stack in GO.

#Why GO?
* Yc post on Ruby rails to go.

#Monoliths

#Metrics

- We measure every request response time over the network, primarily for debugging, not alerting.
- Tracing is useful, we have played around with a tool like Dapper (a nice to have).

#Shadow Live Traffic
* Shadow traffic for months for some endpoints before we released to 100% of users.
Works great for Read APIs.
Complicated setup for Write APIs with database snapshot and DB compare.

#Throttles
Return the amount of time in the error message.
Currently evolving into auto throttling.

