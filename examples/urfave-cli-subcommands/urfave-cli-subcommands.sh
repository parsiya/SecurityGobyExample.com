# Running the program without any arguments will print usage.
# Same as running with -h or --help.
$ go run urfave-cli-subcommands.go
NAME:
   AppName - application usage

USAGE:
   urfave-cli-subcommands.exe [global options]
      command [command options] [arguments...]

DESCRIPTION:
   description two

# Note Authors slice is printed first and then Author
AUTHORS:
   Author2 <author2@example.com>
   Author3 <author3@example.com>
   Author1 <author1@example.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

# Subcommands are grouped based on their category
   Hacking:
     hack, haxorx, pwn  hack target
     crack, c           crack target

   Recon:
     scan, s  scan target

GLOBAL OPTIONS:
   --help, -h  show help
# Error message and status code fro cli.NewExitError
no commands provided
exit status 2

# Running the program with a subcommand and no arguments
# prints the scan help
$ go run urfave-cli-subcommands.go scan
NAME:
   urfave-cli-subcommands.exe scan - scan target

USAGE:
   urfave-cli-subcommands.exe scan [arguments...]

CATEGORY:
   Recon
no arguments for subcommand
exit status 3

# Calling any of hack aliases is the same as calling hack
$ go run urfave-cli-subcommands.go pwn
NAME:
   urfave-cli-subcommands.exe hack - hack target

USAGE:
   urfave-cli-subcommands.exe hack [arguments...]

CATEGORY:
   Hacking
no arguments for subcommand
exit status 3

# Calling hack with a target performs the action
$ go run urfave-cli-subcommands.go hack 127.0.0.1:22
hacking 127.0.0.1:22

# We can also use the hack alias (pwn)
$ go run urfave-cli-subcommands.go pwn 127.0.0.1:22
hacking 127.0.0.1:22
