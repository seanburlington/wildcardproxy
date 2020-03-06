# Wildcard Reverse Proxy #

A wildcard reverse proxy to share all your containers on port 80 with easy dns

## Create a [wildcard DNS](https://en.wikipedia.org/wiki/Wildcard_DNS_record) entry pointing to my IP  ##

  e.g `*.local.tangiblebytes.co.uk. 3600 IN	A	192.168.1.117`

  This means that any domain name within that pattern all point to my local IP
  
You can use a local dns server if that works - but you can also use public DNS - even if (as in my case) you use a [private IP](https://en.wikipedia.org/wiki/Private_network) 

## Build this proxy (optional) ##

`docker build -t wildcardproxy  .`

## Create a docker network ##

`docker network create mynetwork`

## Start the proxy ##

either using the local image you just built

`docker run --name myproxy --network mynetwork -p 80:80 -d wildcardproxy`

Or run the pre-built image 

`docker run --name myproxy --network mynetwork -p 80:80 -d tangiblebytes/wildcardproxy`

This starts the proxy server, attaches port 80 on the container to port 80 on the host and attaches the prxy to your new network

## start any new containers ##

For example run the default apache container 

`docker run --name webserver --network mynetwork  httpd`

In your browser visit the site using the short docker name with your wildcard domain appended

eg

http://webserver.local.tangiblebytes.co.uk/

## Explanation ##

What the proxy does is look at the incoming hostname, takes the part before the first . and forwards the request to a container with taht short name on the local docker network.

This code was based on [Writing a Reverse Proxy in just one line with Go](https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b)

You might also want to look at [Automated Nginx Reverse Proxy for Docker](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/)

But I think the solution I have written is very simple and requires no changes when containers stop and start

Just bring up as mainy containers as you want - and access them all naturally using a simple domain name 

## Things that could be improved ##

It's only http - it would be better if the proxy handled SSL 


