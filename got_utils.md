Go utils for writing go code:
=====
[Docs from official website](http://golang.org.doc/code.html)  
Workspaces  
The go tool is designed to work with open source code maintained in public repositories.  
Although you don't need to publish your code, the model for how the environment is set up works the same whether you do or not.  
Go code must be kept inside a workspace. A workspace is a directory hierarchy with three directories at its root:  
   * src contains Go source files organized into packages (one package per directory),  
   * pkg contains package objects, and  
   * bin contains executable commands.   

The go tool builds source packages and installs the resulting binaries to the pkg and bin directories.  
The src subdirectory typically contains multiple version control repositories (such as for Git or Mercurial) that track the development of one or more source packages.  
To give you an idea of how a workspace looks in practice, here's an example: 
#### Directory structure
```
bin/
    hello                           # command executable
    outyet                          # comma executable
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutils.a           # package object
src/
    github.com/golang/example/
        .git/                       # git repo metadata
        hello/
            hello.go                # command source
        outyet/
            main.go                 # command source
```