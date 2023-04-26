# Distributed-Systems

<!--
## Table of Contents

* [Introduction](#introduction)
* [Requirements](#requirements)
    * [Run Binary](#run-binary)
    * [Clone repo](#clone-this-repo)
* [How to use](#how-to-use)
    * [Examples](#Examples)
--> 

## Introduction

This project was developed as part of a college assignment at Universidade Federal do Rio de Janeiro (UFRJ), with the aim of exploring inter-process communication (IPC) in distributed systems.

It contains the source code for a program that implements several IPC mechanisms, including signals, pipes, and sockets. This program enables you to perform a variety of tasks, such as sending signals to processes, setting up socket servers, and connecting to socket servers as clients. It is completely written in Go, a programming language that is designed for building distributed systems.


## Requirements

### Run Binary
In case you are running a Linux x64 machine and just intend to run the program:

- Download the binary (trab1)
- Open your terminal and go to the folder where it is downloaded
- Run: `chmod +x trab1`

```bash
$ cd Downloads/
$ chmod +x trab1
$ ./trab1 [parameters]
```

### Clone this repo
If you are running on other settings or intend to alter the code

- Download a Go compiler (which you can download here: https://go.dev/dl/)
- Clone this repository
- Open its folder (still on terminal)
- Run: `go build -o trab1`

```bash
$ git clone https://github.com/ptk-trindade/Distributed-Systems.git
$ cd Distributed-Systems
$ go build -o trab1
$ ./trab1 [parameters]
```

## How to use

All the input parameters are given as arguments when executing the program.

Here is an image that exemplifies the parameters:

![image](https://cdn.discordapp.com/attachments/405585265055236128/1100917153454112868/Arquitetura_TP1_1.png)

The blue boxes are parameters that should be written exactly as you see them.

The red boxes should be substituted with a number.

And the yellow box is an optional parameter.

### Examples
#### Sending a signal 9 to process 1234:
```
$ ./trab1 signals send 1234 9
```

#### Creating socket server
```
$ ./trab1 sockets server
```

#### Creating a client that does not close the server
```
$ ./trab1 sockets client
```