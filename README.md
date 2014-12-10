boxstrapper
===========

Next Gen Boxstrap in Golang

Usage:

bs ap PACKAGE[:groups,...]

will:
* add package via your package manager
* list package with groups or default in boxstrap.d/packages.bss
* commit a revision with this request to boxstrap.d

Example:

bs ap vim:development

or

bs ap emacs


