# Robot Simulator
This is an application to simulate toy robot moving on a square tabletop, The robot is free to roam around the surface of the table, but must be prevented from falling to destruction. Any movement that would result in the robot falling from the table must be prevented, however further valid movement commands must still be allowed.

Table of Contents
=================

  * [Preparations](#preparations)
    * [Install GoLang](#install-golang)
    * [Set Up ENV](#set-up-env)
    * [Go Version](#go-version)
    * [Working Directory](#working-directory)
  * [Running tests](#running-tests)
  * [Compiling Or Running the Application](#compiling-or-running-the-application)
    * [Compiling](#compiling)
    * [Running](#running)
      * [With Table Width and Length Option](#with-table-width-and-length-option)
      * [With File Commands Option](#with-file-commands-option)
      * [Help Option](#help-option)
  * [Commands](#commands)

## Preparations

### Install Golang

Please Follow the instruction [here](https://golang.org/doc/install)


### Set Up ENV

Make sure `$GOPATH` environment variable are correctly setup

    #this should be your workspace for all the go packages (result will be different from yours)
    $ echo $GOPATH
    $ /Users/macuser/Documents/gowork

Some comments start with !, TODO, *, and other such symbols. The better comments VSCode extension can help the developer experience by coloring those comments in special ways.

### GO version

At the time of writing, we are using `1.7.4` .
to check your version

    $ go version
    $ go version go1.7.4 darwin/amd64

### Working Directory

Ensure the structure of your working directory is `$GOPATH/src/robot-simulator/.`, copy or clone all the source code inside this working directory.

## Running tests

    $ go test -v -cover ./...

## Compiling Or Running the Application
This program has been pre-compiled and the executable is called main in the working folder.

### Compiling

To Compile the code first make sure you are in the working director

    $ pwd
    $ /Users/macuser/Document/gowork/src/robot-simulator

Then download all dependencies

    $ go get -t -v ./...

Then compile the code, which will create an executable under `$GOPATH/bin/robot-simulator`

    $ go build *.go

### Running

You can run the application from the executable file in the working directory

    $ /Users/macuser/Document/gowork/robot-simulator/main

Alternatively you can also compile and run the code from your working directory

    $ go run main.go

#### With Table Width and Length Option

You are able to change the width and length of the table before starting the program by providing it using flag `-width` and `-length` like below.

    $ ./main -width=10 -length=10

#### With File Commands Option

You are also able to feed commands using flat file in your working directory by providing the filename before starting the program by using flag `-f` like below

    $ ./main -f=test-data.Text

#### Help Option

Help is available on how to use the flag options by using `-h` like below

    $ ./main -h

## Commands

Below are available commands to simulate the robot:

`PLACE X,Y,F` will put the toy robot on the table in position X,Y,F.

`MOVE` will move the toy robot one unit forward in the direction it is
currently facing.

`LEFT` will rotate the robot 90 degrees to the left from the robot direction without changing the position of the robot.

`RIGHT` will rotate the robot 90 degrees to the right from the robot direction without changing the position of the robot.

`REPORT` will announce the X,Y and F of the robot. This will print out X,Y,F of the current state of the robot.
