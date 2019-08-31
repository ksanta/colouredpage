# Coloured page web server
This is a very simple web server that will display the given colour as the background
colour.

The background colour is given as a command line argument and should be a standard CSS
colour name.

# Build the image
```shell script
docker build -t my-golang-app .
```

# Run a container
This runs server with default green background
```shell script
docker run -it --rm --name my-running-app -p 8080:8080 my-golang-app
```

This runs server with specified blue background
```shell script
docker run -it --rm --name my-running-app -p 8080:8080 my-golang-app app blue
```