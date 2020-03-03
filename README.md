# Fact.go
Factorial Program in GO language that can be executed:  
  
(a) on your Mac OS
(b) in a Docker container on a Mac OS  
(c) on Pivotal Cloud Foundry using the Go Buildpack
(d) on Pivotal Cloud Foundry using a Docker Image  
(e) on a K8s environment using a Docker Image
  
Start with option (a) to make sure you download the program from github  
  
# (a) Fact.go on Mac OS  
  
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
  
# (b) Fact.go using Docker on Mac OS  
  
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
  
# (c) Fact.go on Pivotal Cloud Foundry using the Go Buildpack
  
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

# (d) Fact.go on Pivotal Cloud Foundry using a Docker Image
  
- Prerequisite: You need the CF CLI on your Mac  
- And you need to be pointing at a valid CF API, e.g.: `cf api api.run.pivotal.io`
  
Open a terminal window on your Mac and execute the following commands:  
  
```  
Mac $ cf delete fact    # in case executed the step (c) described above
Mac $ cf push fact --docker-image rmeira/fact
```
  
- Once the `cf push` has completed, you can test the factorial program by trying the following URLs:         
  - http://fact.cfapps.io/600  
  - http://fact.cfapps.io/header  
  - http://fact.cfapps.io/health
  - http://fact.cfapps.io/version
  
# (e) Fact.go on a K8s Cluster using a Docker Image

- Prerequisite: The kubectl CLI on your Mac and access to a K8s Cluster 
- In my case, the K8s Cluster is called `my-cluster`

Open a terminal window on your Mac and execute the following commands: 

```
Mac $ kubectl config use-context my-cluster
Switched to context "my-cluster".
```
```
Mac $ kubectl run --generator=run-pod/v1 factorial --image=rmeira/fact
pod/factorial created
```
Let's open a bash shell to access the Pod running the rmeira/fact image and test the Factorial program:

```
Mac $ kubectl exec -t -i factorial bash

root@factorial:/go/src/app# curl 127.0.0.1:3000/40; echo

Calculating Factorial: 40! = 815915283247897734345611269596115894272000000000
```

If we wish to give access to an external user, we'll need to expose the Factorial Pod appropriately:

```
Mac $ kubectl expose pod factorial --type=NodePort --port=80 --target-port=3000
service/factorial exposed
```
```
kubectl get pods -o wide
NAME        READY   STATUS    RESTARTS   AGE   IP             NODE                                      NOMINATED NODE   READINESS GATES
factorial   1/1     Running   0          11m   10.200.45.15   vm-6b355629-2699-4839-6ade-8235603c3110   <none>           <none>
```
```
Mac $ kubectl get services     
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
factorial    NodePort    10.100.200.77   <none>        80:31949/TCP   58s
kubernetes   ClusterIP   10.100.200.1    <none>        443/TCP        9h
```
```
Mac $ kubectl get nodes -o wide
NAME                                      STATUS   ROLES    AGE   VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION      CONTAINER-RUNTIME
vm-6b355629-2699-4839-6ade-8235603c3110   Ready    <none>   9h    v1.15.5   10.0.11.11                  Ubuntu 16.04.6 LTS   4.15.0-66-generic   docker://18.9.9
```

- What do we know about the rmeira/fact docker image?
   - It's running on a K8s Pod housed inside the vm-6b355629-2699-4839-6ade-8235603c3110 K8s Worker Node 
   - The vm-6b355629-2699-4839-6ade-8235603c3110 VM is located at 10.0.11.11
   - Port 31949 on the Worker Node is mapped to the Factorial Service at 10.100.200.77 on port 80
   - The Factorial NodePort Service is mapped to port 3000 of the Pod running the rmeira/factorial docker image
   - The Pod running the rmeira/factorial docker image is at location 10.200.45.15

- Assuming you have access to the 10.0.11.0/26 Network, you can now test the Factorial program:

```
VM@10.0.0.10:~$ curl http://10.0.11.11:31949/40; echo
Calculating Factorial: 40! = 815915283247897734345611269596115894272000000000
```
   
   





