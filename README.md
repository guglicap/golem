![golem](https://github.com/guglicap/golem/blob/master/pics/defaults.png)

## Golem

Go program to generate input for [lemonbar](https://github.com/LemonBoy/bar)

### Installation

* If you already have Go installed:
    `go get github.com/guglicap/golem`

* If you don't, install Go: 

    * Archlinux: `# pacman -Syu go gcc-go`
    * [Ubuntu](https://github.com/golang/go/wiki/Ubuntu) 
    * [Others](https://golang.org/dl/)

    * Then run the command in the section above.

You should now have a binary file `$GOPATH/bin/golem`

### Usage

`golem -config [config]`

Note that you need to pipe that into lemonbar. Here's an example from my setup:

`gobar -config ~/.config/golem.json | lemonbar -f "Source Code Pro-8" -f "FontAwesome-8" -B "#222222" -F "#a4dfdf" -g 1920x25+0+0 -d | sh`  

FontAwesome is recommmended, as most defaults use it. It is also recommmended to pipe the output to `sh`.

### Configuration 

Golem requires a json encoded config file. See the [examples](https://github.com/guglicap/golem/blob/master/examples/).  
All of the top-level fields are required, while you can edit the `Modules` array to customize the look of your bar. Every module is encoded in this form:

    {
         "Position", //If omitted it will be the same as the module above
         "Handler",  //Required.
         "Colors": { //Optional. Colors must be in hex format.
             "Background",
             "Foreground",
         },
         "Options":{
              //varies based on the module
         }
    }          

### Handlers

Here's a list of handlers.  
Most of them can be used with default options.  
For a list of options supported by the module, click on the link.


* [cpu](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#cpu)

* [date](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#date)

* [disk](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#disk)

* [launcher](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#launcher)

* [mem](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#mem)

* [text](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#text)

* [button](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#button)

* [mpd](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#mpd)

* [net](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#net)

* [power](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#power)

* [whoami](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#whoami)

* [ws](https://github.com/guglicap/golem/blob/master/modules/MODULES.md#ws)