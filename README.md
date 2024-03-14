Tutorial: https://go.dev/doc/tutorial/create-module

1. start a new module using the
    >go mod init <module path: ex. example.com/greetings> 

2. For production use, you’d publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system.

To do that, use the go mod edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is).

From the command prompt in the hello directory, run the following command:
$ go mod edit -replace yauhen.example.com/greetings=../greetings

3. Ending a file's name with _test.go tells the go test command that this file contains test functions.
   >go test

4. While the go run command is a useful shortcut for compiling and running a program when you're making frequent changes, it doesn't generate a binary executable.

This topic introduces two additional commands for building code:

The go build command compiles the packages, along with their dependencies, but it doesn't install the results.
The go install command compiles and installs the packages.

5. From the command line in the hello directory, run the go build command to compile the code into an executable.
$ go build

6. Discover the Go install path, where the go command will install the current package.
You can discover the install path by running the go list command, as in the following example:

$ go list -f '{{.Target}}'

7. 
