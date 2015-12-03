(function() {
    var sock = new WebSocket("ws://127.0.0.1:12345/handler"),
        can = document.getElementById('game'),
        ball = new Ball(can),
        p1 = new Pad(can);
    sock.onopen = function(e) {
        addMsg("Onopen: " + e.data);
        console.log(e.data);
    }
    sock.onclose = function(e) {
        addMsg("Onclose: " + e.data);
        console.log(e.data);
    }
    sock.onmessage = function(e) {
        addMsg("Onmessage: " + e.data);
        var json = JSON.parse(e.data)
        switch (json.t) {
            case 2:
                ball.updatePosition(json.d.x, json.d.x);
                break;
        }
        console.log(json);
    }
    if (can.getContext) {
        setInterval(function() {
            can.getContext('2d').clearRect(0, 0, can.width, can.height);
            ball.draw();
            p1.draw();
        }, 17)
        can.onmousemove = function(e) {
            p1.updatePos(e.clientX - can.offsetLeft);
        }
    }
})();

function addMsg(msg) {
    var el = document.getElementById('msgBox'),
        elChild = document.createElement('div');
    elChild.innerHTML = msg;
    // Prepend it
    el.insertBefore(elChild, el.firstChild);
}