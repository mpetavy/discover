# DISCOVER documentation

## Description

DISCOVER is a service discovery tool by nwtwork broadcasting

## Connection roles and data roles

With two DISCOVER there must always be a "client" and a "server" to communicate.

DISCOVER can be starter in "client" (-c) or "server" (-s) connection role.

* As a "server" DISCOVER waits for incoming network broadcast messages
* As a "client" DISCOVER sends network broadcast messages and waits fro responses of running servers

## Architecture

Network boradcasts is a very commonway to to services recovers by using the UDP protocol.

Please keep in mind that such broadcasts are not transfered outside the current network via routers.

A service discovery is setup and find by an UID and a broadcast network port.

The DISOVER server "-s" waits for incoming network broadcasts via the UDP protocol. Only if the same UDP port for
broadcast sending and receiving of DISOVER server "-s" and client "-c" is used and the same UID "-uid" is defined on
both endpoints a successfull service discovery happens. The DISCOVER server then reports back to the DISCOVER client
its "-info" information.

Sample:

    discover -s :5000 -uid hydra -info "myinfo"
    discover -c :5000 -uid hydra

## GO development

DISCOVER is developed with the Google GO tooling.

Current used version 1.20.1

By using the GO programming language (https://golang.org) multiple architectures and OS are supported. You can find a
complete list of all supported architectures and OS at https://github.com/golang/go/blob/master/src/go/build/syslist.go

## Build

To build a binary executable for your preferred OS please do the following:

1. Install the GO programming language support (http://go.dev)
1. configure your OS env environment with the mandatory GO variables
    1. GOROOT (points to your <GO installation directory>)
    1. GOPATH (points to your <GO development direcotry root>)
    1. OS PATH (points to your <GO development direcotry root>/bin)
1. Open a terminal
1. CD into your GOPATH root directory
1. Create a "src" subdirectory
1. CD into the "src" subdirectory
1. Clone the DISCOVER repository
1. CD into the "DISCOVER" directory
1. Build:
    1. If you would like to cross compile to an other OS/architecture define the env variable GOOS and GOARCH along to
       the values defined here https://github.com/golang/go/blob/master/src/go/build/syslist.go
    1. Build DISCOVER by "go install". Multiple dependent modules will be downloaded during the build
    1. After a successful build you will find the DISCOVER executable in the "GOPATH\bin" directory

## Installation as application

Like all other GO based application there is only the file `DISCOVER.exe` or `DISCOVER` which contains the complete
application.

Just copy this executable into any installation directory you would like. Start the application by calling the
executable `DISCOVER.exe` or `./DISCOVER`

## Installation as OS service

Follow the instructions "Installation as application". To register DISCOVER as an OS service do the following steps.

1. Open a terminal
1. Switch to root/administrative rights
1. CD into your installation directory
1. Installation DISCOVER as an OS service:
    1. Windows: `discover -service install`
    1. Linux: `./discover -service install`
1. Uninstallation DISCOVER as an OS service:
    1. Windows: `discover -service uninstall`
    1. Linux: `./discover -service uninstall`
