var sock = null,
    button = null,
    can = null,
    ball = null,
    player = null,
    opponent = null;

(function() {
    sock = new WebSocket("ws://127.0.0.1:12345/handler");
    button = document.getElementById('start-game');
    can = document.getElementById('game');

    button.setAttribute('style', 'display: block');
    sock.onopen = function(e) {
        addMsg("Połączono");
    }
    sock.onclose = function(e) {
        addMsg("Rozłączono");
    }
    sock.onmessage = function(e) {
        var json = JSON.parse(e.data)
        switch (json.t) {
            case 3:
                if(ball === null){
                    ball = new Ball(can, json.d.x, json.d.y)
                } else {
                    ball.updatePosition(json.d.x, json.d.y);
                }
                break;
            case 4:
                console.log('Player: X: ' + json.d.x +  ' Y: ' + json.d.y);
                if(player === null){
                    player = new Pad(can, sock, json.d.x, json.d.y);
                } else {
                    player.setPos(json.d.x, json.d.y);
                }
                break;
            case 5:
                //console.log('Opponent: X: ' + json.d.x +  ' Y: ' + json.d.y);
                if(opponent === null){
                    opponent = new Pad(can, null, json.d.x, json.d.y);
                } else {
                    opponent.setPos(json.d.x);
                }
        }
    }
    if (can.getContext) {
        setInterval(function() {
            if (player !== null && ball !== null /*&& opponent !== null*/) {
                can.getContext('2d').clearRect(0, 0, can.width, can.height);
                ball.draw();
                player.draw();
                //opponent.draw();
            }
        }, 17)
        can.onmousemove = function(e) {
            if (player !== null) {
                player.updatePos(e.clientX - can.offsetLeft);
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
        button.setAttribute('style', 'display: none');

        sock.send(JSON.stringify({
            't': 1,
            'd': {
                'cW': can.width,
                'cH': can.height
            }
        }));
    }
}