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
export GOROOT=/usr/local/go
export GOPATH=`pwd`
export PATH=$PATH:$GOROOT/bin
go install blaze
./bin/blaze
#-->Manually open security group so that yacback can talk to Blaze and vice versa.
./userdata/golang/install <dbpassword> $GOPATH

## Tests

## Usage

## Teardown
Delete the c9 instance.

## Credits
