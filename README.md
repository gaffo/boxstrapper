Boxstrap
===========
Simple, Repeatable Machine Configuration.

Install
===========
TODO: Figure out github pages && releases

Boxstrap means to make setting up your workstation / development machine as easy and repeatable as possible.

Boxstrapper: Next Gen Boxstrap
===========
This repo contains boxstrapper, a WIP replacement for Boxstrap. It will eventually be pulled into boxstrap.

Initally boxstrap was a set of shell scripts which were for managing and remembering what packages you had installed on your system. It worked under ubuntu variants and RHEL5.

Goals
===========
This new version aims to be:

Completely Self Contained
-----------
This new verison is written in go and will have no dependencies that aren't on a stock system. Anything that is required, like git (libgit2) will be statically linked into the executable and distributed that way.

In this way we aim to have a very minimal and quick bootstrap setp to get you up and going:

* Download bs executable
* Optionally recover from a previous configuration
* Start installing stuff

No more:

* Install scripting language
* Install 5 system dependencies
* Install scripting packages
* Polluting your machine with different revisions of scription languages you use for your apps (like ruby 1.8 for chef/puppet)

As Seamless as Possible
-----------
* Replace common system commands like apt-get / yum with simple boxstrap equivalents.
* Possibly eventually shim out these command completely so you don't even have to think about using boxstrap
* These command record into the ~/boxstrap.d directory in simple configuration.
* Make these commands even easier than the native ones (think adding a ppa which is 3 steps)

Desktop/Laptop Centric (as opposed to Server Centric)
-----------
All of the existing dev ops / configuration scripting packages I've seen began life as server configuration systems. They usually grow to include a single machine configuration mechanism, but this is really more arranged around configuring VMs / Containers through mechanisms like Vagrant. None of them are:
* designed to record what you set up on your dev machine and why.
* slimmed down to only handle the local use case

In fact they're all designed around managing a cluster of servers and daemon processes.

Packagages Packages Packages
-----------
The current idea is that boxstrap is about packages. That's it. Now I didn't say what type of packages. Currently to me packages to me mean:

* System Packages
* Ruby/Python/etc packages
* Downloadable Packages (wget, extract, symlink, add to path). Thinking this is really a url -> location, list symlinks.
* File packages (templates) like config files like zshrc, i3 config, etc
* ? The FUTURE!

So it looks like we'll need a few constructs:

* Groups (not required for above but I'm going with it): This is a tag of WHY you installed said thing. Packages can have 0 (default) to many groups. Groups can be added and removed later. We don't really care.
* Depdendency Chains: If you have to install maven first via system package before you can add another package from source control and build it... You need to be able to express this dependency. I think currently this would be like -> package.
* Simple Primitives: symlink, curl, git, directory creation.

Anyway these will get added as I need them.

Currently the idea is that packages.bss contains a list of packages like this
```
packagename: groups [-> dependencies...]
...
```

where packages are either a system package, or they refer to a file in the packages directory named the same, eg if you have an entry in packages for boxstrap, then if there is 
a file in boxstrap.d/packages/bs.bs

This file will contain the scripting in the form of:
```
dependencies
download(url, location)
symlink(from, to)
...
```

So generally you'll end up with a bunch of this.

Dunno, this needs thought, as you can tell by me saying Dunno.

Extensibility
-----------
For now boxstrap isn't going to allow extension. It's all packaged. All ofthe extension mechansims I've seen in other config management systems are way over complicated and super flexible. But they're also hard to learn. I don't want that. I want super simple. I have a few ideas in order of preference:

* Give back, you can extend boxstrap and submit back
* Compilable go scripts with a simple framework. You toss plugins in a dir, they get compiled as sub commands, and they get added to the system. This keeps us on one language. Not that it's a language everyone knows.
* Embedded scripting language: Looks like ruby && javascript are embeddable in golang. Providing a simple minimal plugin layer might be good, but it makes boxstrap into a 2 headed beast (well 3 if you count native). The only libraries allowed to these languages would be provided through golang.

Developing
============
Rule #1, Boxstrap is completely TDD. You want to add to it, you have to do this. Look at the history from the beginning. Othern than the ubntu and simple early filesystem drivers (which later got tested as it grew past a single file wrapper), the entire app is driven out with each commit being a test / feature pair. I'm continuing this. If you want to add to boxstrap, you'll do so too.

Anyway:

if you haven't setup a GOPATH yet:

```
mkdir -p ~/go
cd ~/go
export GOPATH=`pwd`
```

for everything:

```
cd $GOPATH
go get github.com/gaffo/boxstrapper
cd github.com/gaffo/boxstrapper
make
```

It'll take a bit to download the deps but once it does you should be good to go.