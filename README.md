# rr-lb-go

This is a simple pet project of mine. A simple toy Hash based Load balancer.

## How it works

The HTTP server receives a POST request which contains a random string which is to be written in the DATABASE (txt files act as databases). The server thus hashes the 
input and writes it to the database alongside the original data since hashes have a possibility of 16 characters: (0123456789ABCDEF).  


https://github.com/user-attachments/assets/ac7bca80-4e57-4605-8424-9d7daaae852f


## Prerequisites

- Golang installed

## Using it

Choose either the hash-based or the round-robin

- Open up two terminal tabs/instances
- Navigate into either of te two
- On tab/instance 1, run `go run load-balancer.go`
- On tab/instance 2, run `sh runme.sh`

Check the databases/ database partitions

## Contributions

I don't want them. Kindly, keep them to yourself
