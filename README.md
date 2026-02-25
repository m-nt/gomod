Overview ¶
Example
Gonew starts a new Go module by copying a template module.

Usage:

gonew srcmod[@version] [dstmod [dir]]
Gonew makes a copy of the srcmod module, changing its module path to dstmod. It writes that new module to a new directory named by dir. If dir already exists, it must be an empty directory. If dir is omitted, gonew uses ./elem where elem is the final path element of dstmod.

This command is highly experimental and subject to change.

Example ¶
To install gonew:

go install golang.org/x/tools/cmd/gonew@latest
To clone the basic command-line program template golang.org/x/example/hello as your.domain/myprog, in the directory ./myprog:

gonew golang.org/x/example/hello your.domain/myprog
To clone the latest copy of the rsc.io/quote module, keeping that module path, into ./quote:

gonew rsc.io/quote

### usage:
gonew github.com/m-nt/gomod users
