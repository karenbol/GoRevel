#Install Go

Before you can use Revel, first need to install Go.

See official Install Go at https://golang.org

Ubuntu
Install package golang


sudo apt-get install golang

###Set up your GOPATH

If you did not create a GOPATH as part of installation, do so now.

The GOPATH is a directory tree where all of your Go code will live. Here are the steps to do that:

1-Make a directory: mkdir ~/gocode

2-Tell Go to use that as your GOPATH: export GOPATH=~/gocode

3-Save your GOPATH so that it will apply to all future shell sessions:

echo export GOPATH=$GOPATH >> ~/.bash_profile

Now your Go installation is complete.

###Install git and mercurial
Git and Mercurial are required to allow go get to clone various dependencies.

Installing Git

Installing Mercurial

sudo apt-get install mercurial meld

###Get the Revel framework

To get the Revel framework, run

go get github.com/revel/revel

###Get and Build the Revel command line tool

The revel command line tool is used to build, run, and package Revel applications.

Use go get to install:

go get github.com/revel/cmd/revel
# Welcome to Revel

## Getting Started

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Start the web server:

    revel run myapp

   Run with <tt>--help</tt> for options.

### Go to http://localhost:9000/ and you'll see:

"It works"

### Description of Contents

The default directory structure of a generated Revel application:

    myapp               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files

app

    The app directory contains the source code and templates for your application.

conf

    The conf directory contains the application’s configuration files. There are two main configuration files:

    * app.conf, the main configuration file for the application, which contains standard configuration parameters
    * routes, the routes definition file.


messages

    The messages directory contains all localized message files.

public

    Resources stored in the public directory are static assets that are served directly by the Web server. Typically it is split into three standard sub-directories for images, CSS stylesheets and JavaScript files.

    The names of these directories may be anything; the developer need only update the routes.

test

    Tests are kept in the tests directory. Revel provides a testing framework that makes it easy to write and run functional tests against your application.

### Follow the guidelines to start developing your application:

* The README file created within your application.
* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](http://revel.github.io/docs/godoc/index.html).

## Contributing
We encourage you to contribute to Revel! Please check out the [Contributing to Revel
guide](https://github.com/revel/revel/blob/master/CONTRIBUTING.md) for guidelines about how
to proceed. [Join us](https://groups.google.com/forum/#!forum/revel-framework)!
