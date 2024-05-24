FROM nats-streaming:latest



# docker build  -t nats-streaming:latest .
# docker run --name nats-streaming-container -d -p 4222:4222 -p 8222:8222 nats-streaming:latest
# без  указания имени кластера сервер работает на дефолтном кластере "test-cluster"

# Run a NATS Streaming server
# Each server exposes multiple ports
# 4222 is for clients.
# 8222 is an HTTP management port for information reporting.
#
# To actually publish the ports when running the container, use the Docker port mapping
# flag "docker run -p <hostport>:<containerport>" to publish and map one or more ports,
# or the -P flag to publish all exposed ports and map them to high-order ports.
#
# This should not be confused with the NATS Streaming Server own -p parameter.
# For instance, to run the NATS Streaming Server and have it listen on port 4444,
# you would have to run like this:
#
#   docker run -p 4444:4444 nats-streaming -p 4444
#
# Or, if you want to publish the port 4444 as a different port, for example 5555:
#
#   docker run -p 5555:4444 nats-streaming -p 4444
#
# Check "docker run" for more information.
