# A light server for test purpose

The images are uploaded to Docker Hub -- https://hub.docker.com/r/mathnao/light-test-server/.

How to run:
```
$ docker run -P -d mathnao/light-test-server
```

Now, assuming we found out the IP address and the port that mapped to port 8080 on the container, in a browser we can make a request to the webserver and get a `200` status code response with a message in the page's body.