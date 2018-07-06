# blazemeter

## Motivation
Simulate Blazemeter but within a firewall.

## Screenshots

## Technology/Framework Used

## Features

## Installation

## Cloud9 Installation
```
git clone https://github.com/craigwongva/blazemeter
cd /home/ec2-user/environment/blazemeter
./upsert-route53
export GOROOT=/usr/local/go
export GOPATH=`pwd`
export PATH=$PATH:$GOROOT/bin
go install blaze
./bin/blaze

## Tests

## Usage

## Teardown
Delete the c9 instance.

## Credits
