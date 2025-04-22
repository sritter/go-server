# My server

I started with the base server code from [this](https://dev.to/andyjessop/building-a-basic-http-server-in-go-a-step-by-step-tutorial-ma4) article, and the source was found [HERE](https://github.com/andyjessop/simple-go-server)

I then startd to work on Web api to see how this works.  Found "gin" which seems to make more sense that the original server, at least at this point.  
   * [THIS](https://dev.to/chefgs/develop-rest-api-using-go-and-test-using-various-methods-8e0) tutorial


## Building and deploying
 * Build docker image
```
docker build -t go-server:latest .
```

 * Running on local
```
docker run -p 8080:8080 go-server:latest
```