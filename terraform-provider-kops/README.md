
Terraform Provider
==================

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/honestbee/devops-tools`

```sh
$ mkdir -p $GOPATH/src/github.com/honestbee; cd $GOPATH/src/github.com/honestbee
$ git clone git@github.com:honestbee/devops-tools
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/honestbee/devops-tools/terraform-provider-kops
$ make build
```

Using the provider
----------------------

Follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.9+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-kops
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```
