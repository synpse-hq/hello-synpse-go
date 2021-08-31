# Build, Publish And Run A Go Application

In our hands-on section, we show how to deploy a deployable image file using synpse CLI. The question we are going to answer here is how do we do that from the original source. In this Getting Started article, we look at how to deploy a Go application on Synpse AMD64 and ARM64 architectures.

The main difference here will be building and publishing Docker images that can run on different hardware.

### Prerequisites

- Synpse account, register [here](https://cloud.synpse.net/)
- At least one connected device (your laptop, homelab server, Raspberry Pi, etc)

# The Hello Synpse Application

The hello-synpse application is, as you'd expect for an example, small. It's a Go application that uses the gin web framework. Here's all the code form main.go:

```golang
package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", handleIndex)
	r.GET("/:name", handleIndex)
	r.Run(":8080")
}

func handleIndex(c *gin.Context) {
	name := c.Param("name")
	if name != "" {
		name = strings.TrimPrefix(c.Param("name"), "/")
	}
	c.HTML(http.StatusOK, "hellosynpse.html", gin.H{"Name": name})
}

```

The main function sets up the server after loading in templates for pages to be output. Those templates live in ./templates/. When a request comes in, the handleIndex function looks for a name and feeds that name to a template. The template itself, hellosynpse.html, is very simple too:

```html
<!DOCTYPE html>
<html lang="en">
  <head> </head>
  <body>
    <h1>Hello from Synpse</h1>
    {{ if .Name }}
    <h2>and hello to {{.Name}}</h2>
    {{end}}
  </body>
</html>
```

We're using a template as it makes it easier to show what you should do with assets that aren't the actual application.

# Building the Application for AMD64 (x86)

We have created several targets in the `Makefile`, that can be helpful (just update the registry). To build it:

```
make build-all
```

And then, we need to push it to the Docker registry:

```
make push-all
```

Images will be available as:

```
quay.io/synpse/hello-synpse-go:arm64
quay.io/synpse/hello-synpse-go:arm32
quay.io/synpse/hello-synpse-go:amd64
```

# Install synpse CLI

We are ready to start working with Synpse and that means we need `synpse`, our CLI app for managing apps on Synpse. If you've already installed it, carry on. If not, hop over to our installation guide [LINK NEEDED], or, just use our installation script:

```
curl https://downloads.synpse.net/install-cli.sh | bash
```

Once the CLI is installed, vising [access keys page](https://cloud.synpse.net/access-keys) and click on "CREATE KEY" button to get your CLI authentication command.

> P.S. You can also use web UI :)

# Deploying the Application

In Synpse, we believe in declarative model of deploying and managing applications. Therefore, we require users to write a simple `yaml` format manifest that describes which containers should be deployed:

```yaml
# hello.yaml
name: hello-synpse # <- application name 
scheduling:
  type: AllDevices # <- deploys on all project devices
  selectors: {}
spec:
  containers:
    - name: hello
      image: quay.io/synpse/hello-synpse-go:latest # <- our container image
      ports:
        - 8060:8080
```

Save the contents as a `hello.yaml` file and create the application:

```
synpse application create -f hello.yaml
```

Once the application is deployed, you can open it in the browser by calling your device IP and port 8080 (if you are running the agent on localhost, then the URL will be http://localhost:8080)