package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Mod in Golang")
	greeter()
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", router))
	// we can do classic comma ok , but in case of web we use log package. It has fatal
	// which automatically logs error if something fails.

	log.Fatal()

}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// It take response writer and pointer of request.
	// r contains the request somebodu sends like the params and other things.
	// w sends the response back.
	w.Write([]byte("<h1>Welcome to Golang Series on Youtube</h1>"))

}

/*
//Read semantic versioning docs to understand on how versioning happens.
//Mux is used in go for routing
// go get is used to install dependencies

Example:
=> go get -u github.com/gorilla/mux
*AS soon as we run this in terminal. It puts a new line in go.mod file.
The line looks like :
 => require github.com/gorilla/mux v1.8.0 // indirect
indirect means as of now code is not using the mux. As soon as we use it goes away.


To avoid attacks like man in middle attack. A go.sum file is also created.
So I am pulling a repo of a particular version here 1.8.0. So a hash is created and
any change in the repo of mux will lead to change in version and hash.
But if due to any attack the files of mux repo changes and version is not changed .It detects the attack
Because the hash changes and version isnot updated. (Every version has its own hash)


* go env is also a command that gives lot of things . One such thing is GOPATH. If we go to that path
in terminal and do ls, then we get bin , pkg, src.
If we go in pkg and do ls we get mod and subdb
If we go to mod and do ls we get lot of things
One of them is cache. This is where all the dependency we install from outside is stored . And not stored
in current working directory. The reason is next time when we request the module it is cached
and not downloaded again.
Then we go to cache -> download -> github.com -> mux , we can see the mux.


* Once we do go mod tidy
-It removes the indirect that we had and makes sure all used packages are made direct.
-It removes all unused packages
-It installs all the used packages that by mistake might have been uninstalled.
*
 Also when we do " go build ." it creates a file with the name that we gave in go mod init <name>

*/

/*
* "go mod verify" verifies if all modules ok and there is no conflicts. By looking in the code
and go.sum file.

"go list" tells our main package/module.
"go list all" tells all packages/modules installed
"go list -m all" tells all those packages on which our main package is dependent.
"go list -m -versions github.com/gorilla/mux", tells all the version of gorilla mux available
publicly on the internet.


"go mod why github.com/gorilla/mux"
Tells why I am dependent on this module.

"go mod graph" pulls up graph of all dependenvcies.


"go mod edit" useful in case when we need to edit the go mod file.
Flags :
-go <version>: changes version number
-module <name> : changes module name.
Ex. go mod edit -module 1.16

Everything we are bringing in goes into cache and not in prject directory.
*/

/*

Go.sum file contains the version of installed modules and hashes unique to those versions.

*/
