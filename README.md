Tutorial: https://go.dev/doc/tutorial/create-module

1. start a new module using the
    >go mod init <module path: ex. example.com/greetings> 

2. For production use, you’d publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system.

To do that, use the go mod edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is).

From the command prompt in the hello directory, run the following command:
$ go mod edit -replace yauhen.example.com/greetings=../greetings