# fact
Factorial Program in GO language that can be executed:  
  
(a) on your Mac OS  
(b) on a Docker container on Mac OS  
(c) on Pivotal Cloud Foundry  
(d) on Pivotal Cloud Foundry using a Docker Image  
  
Start with option (a) to make sure you download the program from github  
  
# (a) fact on Mac OS  
  
- Prerequisite: You'll need to install the GO Language on your Mac.  
- Do you already have GO installed? Try `Mac $ go version`   
- You should see something like this:   `go version go1.14 darwin/amd64` 
- If you need to install the GO Language, do this: `Mac $ brew install go` 
  
Open a terminal window on your Mac and execute the following command:  

```
Mac $ cd /work  
Mac $ git clone https://github.com/rm511130/fact  
Mac $ cd /work/fact  
Mac $ go run fact.go  
```
  
- You should see a message like this one:  `2020/03/02 16:19:05 Starting Factorial Application...`
- You can then test it using a browser:    `http://localhost:3000/5`
- And you'll get as a reply:               `Calculating Factorials: 5! = 120` 
  
# (b) fact using Docker on Mac OS  
  
- Prerequisite: You'll need to install Docker on your Mac.  
- Do you already have Docker installed? Try `Mac $ docker version`  
- You should see both client and server version information, e.g.: `version 19.03.5 for both client and server.`  
- To install Docker on your Mac OS follow the instructions: https://docs.docker.com/engine/installation/mac/  
- To run the Docker Server on your Mac, perform a Mac Spotlight Search for "Docker" and run it.  
  
On an open terminal window with the familiar Docker Whale icon displayed somewhere at the top of the screen of your Mac, execute the following command:  
  
```
Mac $ cd /work
Mac $ git clone https://github.com/rm511130/fact 
Mac $ cd /work/fact
Mac $ docker build -t fact .  
Mac $ docker run --publish 3000:3000 --name fact --rm fact  
```

- You should see a message like this one: `2020/03/02 21:39:23 Starting Factorial Application...`
- You can now test it using a browser:    `http://localhost:3000/35` 
- And you'll get as a reply:              `Calculating Factorial: 35! = 10333147966386144929666651337523200000000`

I also uploaded the Docker Image to the Docker Hub:  https://hub.docker.com/r/rmeira/fact
  
# (c) fact on Pivotal Cloud Foundry  
  
- Prerequisite: You need the CF CLI on your Mac  
- And you need to be pointing at a valid CF API, e.g.: `cf api api.run.pivotal.io`
  
Open a terminal window on your Mac and execute the following command:  
  
```  
Mac $ cd /work
Mac $ git clone https://github.com/rm511130/fact 
Mac $ cd /work/fact
Mac $ cf push fact -b go_buildpack
```
  
- You should see the usual creating app, route, binding, uploading ... and: urls: fact.cfapps.io  
- You can now test it by trying the following URLs:         
  - http://fact.cfapps.io/600  
  - http://fact.cfapps.io/header  
  - http://fact.cfapps.io/health
  - http://fact.cfapps.io/version

# (d) fact on Pivotal Cloud Foundry using a Docker Image
  
- Prerequisite: You need the CF CLI on your Mac  
- And you need to be pointing at a valid CF API, e.g.: `cf api api.run.pivotal.io`
  
Open a terminal window on your Mac and execute the following command:  
  
```  
Mac $ cf delete fact    # in case executed the step (c) described above
Mac $ cf push fact --docker-image rmeira/fact
```
  
- Once the `cf push` has completed, you can test the factorial program by trying the following URLs:         
  - http://fact.cfapps.io/600  
  - http://fact.cfapps.io/header  
  - http://fact.cfapps.io/health
  - http://fact.cfapps.io/version






