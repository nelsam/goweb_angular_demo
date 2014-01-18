# This is a very simple demo, not intended to show anything particularly interesting

Most of the files in public/ are downloaded from various
sources.  Here is a list of outside libraries:

1. public/js/angular.min.js: Minified angular.js 1.2.9
2. public/js/less.min.js: Minified less compiler, because this is not
   intended to be used with any type of production code.  Since it's
   only intended to be run as a demo on a development server, it just
   uses the javascript client-side less compiler.
3. public/js/bootstrap*.js: Bootstrap javascript files.
4. public/less: Bootstrap less files.
5. public/fonts: Bootstrap fonts.
6. public/partials: Bootstrap html files (I wasn't paying too much
   attention to their contents, which might not have been a good idea ...
   I'll clean them up if I ever add real partials).

The main files that make up the demo are:

1. Any file that ends in .go: These are the Go source files for the API.
2. public/app.html: The main html file, which static.go maps all
   unknown paths to.
3. public/js/app.js: The angular controller that handles the demo.

#### Installing and Running

Pretty standard for a Go project:

1. Install the [Go tool](http://golang.org/doc/install#install)
2. Create a directory for Go development, e.g. `$ mkdir -p ~/dev/golang`
3. Set up the GOPATH environment variable to point to the above
   directory, e.g. `$ export GOPATH="~/dev/golang"`
4. Add `$GOPATH/bin` to your $PATH: `$ export PATH="$GOPATH/bin:$PATH"`
5. Run `go get`: `$ go get github.com/nelsam/goweb_angular_demo`
6. Run the application: `$ PORT=5000 goweb_angular_demo`
7. Visit `localhost:5000` in your web browser.
