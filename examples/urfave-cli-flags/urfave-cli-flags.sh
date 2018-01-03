# Running the program without any arguments will print usage.
# Same as running with -h or --help.
$ go run urfave-cli-flags.go
NAME:
   AppName - application usage

USAGE:
   urfave-cli-flags.exe [global options]
      command [command options] [arguments...]

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -p PORT, --port PORT                 target PORT (default: 0)
   -t HOST, --target HOST, 
      --host HOST  hacking HOST (default: "127.0.0.1")
   --help, -h                           show help
   --version, -v                        print the version
# Running without flags will return an error message
please set both flags
exit status 2

# Running with both flags will hack the target
$ go run urfave-cli-flags.go -host 127.0.0.1 -port 22
hacking 127.0.0.1:22

# Calling one flag will use the default value for the other
$ go run urfave-cli-flags.go -port 1234
hacking 127.0.0.1:1234

# If a default value is not set, the variable's default 
# value will be used
$ go run urfave-cli-flags.go -host 127.0.0.1
hacking 127.0.0.1:0