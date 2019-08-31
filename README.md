# Coloured page web server
This is a very simple web server that will display the given colour as the background
colour.

The background colour is given as a command line argument and should be a standard CSS
colour name.

# Build the image
```shell script
cd ${GOPATH}/src/github.com/ksanta/colouredpage
docker login -u ksanta
docker build -t ksanta/colouredpage .
```

# Run a container
This runs server with default green background
```shell script
docker run -it --rm --name my-colouredpage -p 8080:8080 ksanta/colouredpage
```

This runs server with specified blue background
```shell script
docker run -it --rm --name my-colouredpage -p 8080:8080 ksanta/colouredpage app blue
```

# Publish the image to Docker Hub
```shell script
docker push ksanta/colouredpage
```

# Running on Minikube
```shell script
kubectl run colouredpage --image=ksanta/colouredpage
kubectl expose deployment colouredpage --type=LoadBalancer --port=8080
minikube service colouredpage
```