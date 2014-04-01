var pathfinder = {}

pathfinder.log = function(message) {
    var request = new XMLHttpRequest();
    request.open("POST", "/debug");
    request.setRequestHeader("Content-Type", "text/plain");
    request.send(JSON.stringify(message));
}

pathfinder.start_webshell = function() {
    var request = new XMLHttpRequest();
    request.open("GET", "/debug");
    request.onreadystatechange = function() {
        if (request.readyState === 4) {
            if (request.status === 200) {
                if (request.responseText === "quit") {
                    return
                }
                console.log("CODE: " + request.responseText);
                try {
                    // HISSSSS - Don't ever run this in production
                    var result = eval(request.responseText);
                    if (result !== undefined) {
                        console.log("RESULT: " + result);
                        pathfinder.log(result);
                    }
                } catch (err) {
                    pathfinder.log(err.message);
                    console.log(err);
                }
            } else if (request.status === 400) {
                console.log("ERROR: " + request.responseText);
            } else {
                return
            }
            log.start_webshell();
        }
    }
    request.send(null);
}
