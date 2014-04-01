# Pathfinder

Pathfinder is a quick-and-dirty hack to attach a command-line shell to a
web browser. Pathfinder runs a static web server and provides a single
dynamic endpoint "/debug". If a POST request is made to the debug endpoint,
the body of the post request is printed on the console. Lines of text entered
on the console will be stored in a buffered channel, and GET requests on the
channel will be responded to with the next string in the channel. If the
channel is empty, the request will hang until a new input is available.

This program is designed to allow easier debugging of client-side applications
running in browser without good developer consoles, such as mobile browsers.
One example application provided by the files "index.html" and "debug.js" is
a remote Javascript REPL. You can try this out by running pathfinder in the
project directory and directing your browser to the root URL.

Pathfinder is of course named after the Mars Pathfinder lander, which was the
subject of an intense "remote debugging" session to fix a priority inversion
in its control software.
