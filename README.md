# fact
Factorial Program in GO language that can be executed:  
  
(a) on your Mac OS  
(b) on a Docker container on Mac OS  
(c) on Pivotal Cloud Foundry  
(d) on Pivotal Cloud Foundry using a Docker Image  
  
Start with option (a) to make sure you download the program from github  
  
# (a) fact on Mac OS  
  
Prerequisite: You'll need to install the GO Language on your Mac.  
Do you already have GO installed? Try $ go version  
You should see something like this:     go version go1.6.2 darwin/amd64  
If you need to install the GO Language, do this: $ brew install go  
And we're using a web server called martini, so: $ go get github.com/go-martini/martini  
  
Open a terminal window on your Mac and execute the following command:  
  
$ cd /work  
$ git clone https://github.com/rm511130/fact  
$ cd /work/fact  
$ ls -a  
.		..		.git		Dockerfile	Godeps		Procfile	README.md	fact.go  
$ go run fact.go  
  
You should see a message like this one:  [martini] listening on :3000 (development)  
You can then test it:        http://localhost:3000/5  
And you'll get as a reply:   Calculating Factorials: 5! = 120  
  
# (b) fact using Docker on Mac OS  
  
Prerequisite: You'll need to install Docker on your Mac.  
Do you already have Docker installed? Try $ docker version  
You should see both client and server version information: I see version 1.9.0 for both client and server.  
To install Docker on your Mac OS follow the instructions: https://docs.docker.com/engine/installation/mac/  
To run the Docker Server within VirtualBox on your Mac, perform a Mac Spotlight Search for "Docker Quickstart Terminal" and run it.  
  
On an open terminal window with the familiar Docker Whale displayed on your Mac, execute the following command:  
  
$ cd /work/fact  
$ docker build -t fact .  
$ docker-machine ip default  
192.168.99.100  
$ docker run --publish 6060:3000 --name fact --rm fact  
  
You should see a message like this one:  + exec /go/bin/fact  [martini] listening on :3000 (development)  
You can now test it:         http://192.168.99.100:6060/5  
And you'll get as a reply:   Calculating Factorials: 5! = 120  
  
# (c) fact on Pivotal Cloud Foundry  
  
Prerequisite: You'll need to install Godep on your Mac  
Do you already have Godep installed? Try $ godep version  
You should see something like this:        godep v29 (darwin/amd64/go1.6.2)  
To install Godep on your Mac, do this:   $ go get github.com/tools/godep  
  
Open a terminal window on your Mac and execute the following command:  
  
$ cd /work/fact  
$ godeps save  
$ ls -a  
.		..		.git		Dockerfile	Godeps		Procfile	README.md	fact.go  
$ cf push fact -b https://github.com/cloudfoundry/go-buildpack  
  
You should see the usual creating app, route, binding, uploading ... and: urls: fact.cfapps.io  
You can now test it:         http://fact.cfapps.io/6  
And you'll get as a reply:   Calculating Factorials: 6! = 720  
  







