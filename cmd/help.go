package cmd

var HelpStr string = UsageStr + FlagStr + CommandStr

var CommandStr string = `Commands: 
  help    list this help message
  list    list some things 
  serve   run the http server

`

var FlagStr string = `Flags: 
  -h --help	show this help message

`

var UsageStr string = `cannonical - example template for go apps

Usage: 
  cannonical [--help] <command> [<args>]

`
