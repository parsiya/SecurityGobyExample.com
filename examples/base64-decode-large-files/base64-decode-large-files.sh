# Running the code on a 150 million line (2 characters on
# each line) took only 3 seconds on my machine.
# Measure-Command is the equivalent of time in PowerShell.
PS> Measure-Command{go run .\base64-decode.go -file input.txt}

Days              : 0
Hours             : 0
Minutes           : 0
Seconds           : 2
Milliseconds      : 960
Ticks             : 29604607
TotalDays         : 3.42645914351852E-05
TotalHours        : 0.000822350194444444
TotalMinutes      : 0.0493410116666667
TotalSeconds      : 2.9604607
TotalMilliseconds : 2960.4607