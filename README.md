# Crimson

Crimson is a lightweight web framework in Go.

## Config

Crimson can read `crimson-config.yml` as config file in the root directory of your project. The available options and
default values are:

```yaml
server:
  # Server (application) name
  name: Crimson
  # Port number of the server
  port: 8080
  # Background task wait time when shutting down server (in seconds)
  timeout: 60
browser:
  # If open browser automatically after starting server
  open: false
  # Default open page for browser auto open
  page: /
session:
  # Session provider name
  provider: Crimson
  # Name of cookie for session
  cookie: SESSION_ID
  # Session expires time (in seconds)
  timeout: 3600
```

To config MySQL data source, add the following lines:

```yaml
db:
  mysql:
    # Database driver name, usually "mysql". If you are using MariaDB, "mariadb" is also available
    driver: driver
    # Your MySQL username
    username: username
    # Your MySQL password
    password: password
    # Your MySQL URL in format of ip_address:port_number
    url: url
    # Database name
    name: name
    # Extra parameters for connecting MySQL, e.g., useSSL=false&charset=utf8
    extra: extra
```

## Usage

An example of starting server:

```go
package main

import (
	"github.com/AreSZerA/crimson"
	"net/http"
)

func main() {
	// Config handler for "/" and "/index", GET method is available
	crimson.AddRoute("(/)|(/index)", indexHandler)
	// Config handlers
	crimson.AddRoutes(
		// Config handler for "/login", POST method only
		*crimson.NewRoute("/login", updateHandler, "POST"),
		// Config handler for route begins with "/user/" and end with numbers and letters  
		*crimson.NewRoute("^/user/[0-9a-zA-Z]*$", userHandler, "GET", "POST"),
	)
	// Instantiate server and start
	crimson.NewServer().Start()
	// The Start() above will block the left code
	// Thus, the following code will not be executed until received system interrupt signal")
	// println("Hello, Crimson!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	//...
}
```

An example of using session:

```go
package view

import (
	"github.com/AreSZerA/crimson"
	"net/http"
)

var Manager = crimson.NewSessionManager()

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session := Manager.StartSession(w, r)
	if session.Get("info") == nil {
		session.Set("info", "foobar")
	} else {
		//...
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session := Manager.StartSession(w, r)
	if session.Get("info") != nil {
		session.Delete("info")
	} else {
		//...
	}
}
```

## Acknowledgements

- [Jetbrains GoLand](https://www.jetbrains.com/go/)
- [astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)