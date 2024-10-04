# rr-lb-go

This is a simple pet project of mine. A simple toy Hash based Load balancer.

## How it works

The HTTP server receives a POST request which contains a random string which is to be written in the DATABASE (txt files act as databases). The server thus hashes the 
input and writes it to the database alongside the original data since hashes have a possibility of 16 characters: (0123456789ABCDEF).  

It uses round-robin

[rr-lb-go](recording-2024-10-04-16-30-38.mp4)


## Prerequisites

- Golang installed

## Using it

- Open up two terminal tabs/instances
- On tab/instance 1, run `go run load-balancer.go`
- On tab/instance 2, run `sh runme.sh`

Check the databases/ database partitions

## Contributions

I don't want them. Kindly, keep them to yourself
