# gateway

A simple library for discovering the IP address of the default gateway.


[![Build Status](https://travis-ci.org/jackpal/gateway.svg)](https://travis-ci.org/jackpal/gateway)

Provides implementations for:

+ FreeBSD
+ Linux
+ OS X (Darwin)
+ Solaris
+ Windows `gatewayIP,localAddr := gateway.DiscoverGateway()
`


Other platforms use an implementation that always returns an error.

Pull requests for other OSs happily considered!
