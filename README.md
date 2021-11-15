# Build, Publish And Run A Go Application

In our hands-on section, we show how to deploy a deployable image file using synpse CLI. The question we are going to answer here is how do we do that from the original source. In this Getting Started article, we look at how to deploy a Go application on Synpse AMD64 and ARM64 architectures.

The main difference here will be building and publishing Docker images that can run on different hardware.

### Prerequisites

- Synpse account, register [here](https://cloud.synpse.net/)
- At least one connected device (your laptop, homelab server, Raspberry Pi, etc)

# The Hello Synpse Application

The hello-synpse application is, as you'd expect for an example, small. It's a Go application that uses the native http libraries. Here's all the code form main.go:

```golang
	// override port if provided
	port := ":8080"
	if val, ok := os.LookupEnv("PORT"); ok {
		port = val
	}

	// create filesystem router
	fs := http.FileServer(http.Dir("./ui"))
	http.Handle("/", fs)

	// read certificate variables and if set - run in https mode.
	// else - run in plain http mode.
	certFile, certSet := os.LookupEnv("TLS_CRT")
	keyFile, keySet := os.LookupEnv("TLS_KEY")
	if certSet || keySet {
		log.Println("Listening TLS on " + port)
		http.ListenAndServeTLS(port, certFile, keyFile, nil)
	} else {
		log.Println("Listening on " + port)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
```
# Building the Application for AMD64 (x86) and ARM64 (ARM)

We use docker `buildx` command to build the application image for multiple architectures. 
More on this [here](https://synpse.net/blog/images/multiarch-images/).

We have created build target in the `Makefile`, that can be helpful (just update the registry). To build it:

```
make image
```

Images will be available as:
```
quay.io/synpse/hello-synpse-go
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
      image: quay.io/synpse/hello-synpse-go # <- our container image
      ports:
        - 8080:8080
```

Save the contents as a `hello.yaml` file and create the application:

```
synpse application create -f hello.yaml
```

Once the application is deployed, you can open it in the browser by calling your device IP and port 8080 (if you are running the agent on localhost, then the URL will be http://localhost:8080)