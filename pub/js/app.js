var sock = null,
    button = null,
    can = null,
    ball = null,
    p1 = null;

(function() {
    sock = new WebSocket("ws://127.0.0.1:12345/handler");
    button = document.getElementById('start-game');
    can = document.getElementById('game');

    button.setAttribute('style', 'display: block');
    sock.onopen = function(e) {
        addMsg("Onopen: " + e.data);
        console.log(e.data);
    }
    sock.onclose = function(e) {
        addMsg("Onclose: " + e.data);
        console.log(e.data);
    }
    sock.onmessage = function(e) {
        //addMsg("Onmessage: " + e.data);
        var json = JSON.parse(e.data)
        switch (json.t) {
            case 2:
                ball.updatePosition(json.d.x, json.d.y);
                break;
        }
        console.log(json);
    }
    if (can.getContext) {
        setInterval(function() {
            if (p1 !== null && ball !== null) {
                can.getContext('2d').clearRect(0, 0, can.width, can.height);
                ball.draw();
                p1.draw();
            }
        }, 17)
        can.onmousemove = function(e) {
            if (p1 !== null) {
                p1.updatePos(e.clientX - can.offsetLeft);
            }
        }
    }
})();

function addMsg(msg) {
    var el = document.getElementById('msgBox'),
        elChild = document.createElement('div');
    elChild.innerHTML = msg;
    el.insertBefore(elChild, el.firstChild);
}

function startGame() {
    if (button !== null) {
        ball = new Ball(can);
        p1 = new Pad(can, sock);
        button.setAttribute('style', 'display: none');

        sock.send(JSON.stringify({
            't': 0,
            'd': {
                'cW': can.height,
                'cH': can.width
            }
        }));
    }
}