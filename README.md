At runtime supply a list of ports you want to listen on

Enter a comma-separated list of ports to listen on (default is 22):

The honeypot will log connections to the directory the application is run from the log format will be

honeypot-[port].log, where [port] is the number of the port that the honeypot is listening on. For example, if the honeypot is listening on port 22, the logs will be written to a file named honeypot-22.log.
