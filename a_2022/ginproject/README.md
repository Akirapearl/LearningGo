# Needed steps to use & understand this specific exercise 

* Golang installed on local machine
* Basic knowledge of its usage

[create main.go file and add "package main" at the first line]

go mod init _name of the project_

add an import sentence inside the main.go file with "github.com/gin-gonic/gin" --> Creates a reference to the framework

go get github.com/gin-gonic/gin  --> Actually installs the framework at the current directory and generates a go.sum file


### Structure of the file

func hello --> Returns a basic hello world message when the request is done and received correctly.

	> See line 18, there is where using the http verb GET and the path where the request shall be asked for, the server will return
	the content of the hello function.

func json --> Same concept, but using gin.H to return a JSON formatted message.

	> Same as before, in line 19 is where the path to the json function is defined.

Line 20 --> "Activates" or "Starts" the server, by default it will listen at port 8080, changed to 8070 as a self reminder & test
