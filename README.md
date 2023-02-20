
Installing Go
```
brew install golang
```
Verify Go version and path 
```
go version
go version go1.19.6 darwin/amd64
which go
/usr/local/bin/go
go env GOROOT
/usr/local/Cellar/go/1.19.6/libexec
```
Create Go Workspace and Module
```
mkdir -p /Users/coderprabhu/gospace/go-programs/hello
cd /Users/coderprabhu/gospace/go-programs/hello
go mod init github.com/CoderPraBhu/go-programs/hello
```

Open in VSCode
```
code /Users/coderprabhu/gospace/go-programs/hello
```
Create hello.go file. You will be prompted to install go plugin followed by gopls (Go Please) Go language server for IDE features.
Refers: https://go.dev/doc/tutorial/getting-started

```
Tools environment: GOPATH=/Users/coderprabhu/go
Installing 1 tool at /Users/coderprabhu/go/bin in module mode.
  gopls
Installing golang.org/x/tools/gopls@latest (/Users/coderprabhu/go/bin/gopls) SUCCEEDED  

go install -v github.com/go-delve/delve/cmd/dlv@latest
```

Add Hello World code to hello.go
```
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```

Run Go Program
```
go run .
Hello, World!
```

Do git stuff after creating empty repo in GitHub
```
cd .. 
git init
git status
git checkout -b main
git add .
git commit -m "Initialize go hello world program"
git remote add origin https://github.com/CoderPraBhu/go-programs.git
git push -u origin main
git pull 
git add .
git commit --amend -m "Initialize go hello world program"
git push --force
git pull
```

Add Reference to External Module
```
package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
```

Install the module
```
cd hello
go mod tidy
```

Run
```
go run .
Don't communicate by sharing memory, share memory by communicating.
```

Create Module
```
mkdir -p /Users/coderprabhu/gospace/go-programs/greetings
cd /Users/coderprabhu/gospace/go-programs/greetings
go mod init github.com/CoderPraBhu/go-programs/greetings

```

Add code in greetings.go 
```
package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}
```

Redirect Go tools to use module from local file system instead of module path
```
cd /Users/coderprabhu/gospace/go-programs/hello
go mod edit -replace github.com/CoderPraBhu/go-programs/greetings=../greetings
```

Tidy up dependencies
```
go mod tidy
```