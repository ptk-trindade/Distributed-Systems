# Distributed-Systems

This repository about Distributed Systems was developed for a college project at Universidade Federeal do Rio de Janeiro (UFRJ). This assignment is a study about IPC (Inter-process communication).


## Requirements

### In case you are running a Linux x64 machine and just intend to run the program:

- Download the binary (trab1)
- Open your terminal and go to the folder where it is downloaded
- Run: `chmod +x trab1`

```bash
$ cd Downloads/
$ chmod +x trab1
$ ./trab1 [parameters]
```

### If are running in another settings or intend to alter the code

- Go compiler (which you can download here: https://go.dev/dl/)
- Clone this repository
- Open your terminal and go to the folder where it is downloaded
- Run: `go build -o trab1`

```bash
$ cd Downloads/
$ go build -o trab1
$ ./trab1 [parameters]
```


## How to use

All the input parameters are given as arguments when executing the program.

Here is a image that exemplifies the parameters:

![image](https://cdn.discordapp.com/attachments/405585265055236128/1100917153454112868/Arquitetura_TP1_1.png)

The blue boxes are parameters that should be written exaclty how you see them.

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

#### Creating client that does not close the server
```
$ ./trab1 sockets client
```


