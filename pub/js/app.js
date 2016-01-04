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
        var json = JSON.parse(e.data);
        if (typeof json.length != 'undefined'){
            for(var i=0; i < json.length; i++){
                if(json[i] !== null){
                    switch (json[i].t) {
                        case 3:
                            if(ball === null){
                                ball = new Ball(can, json[i].d.x, json[i].d.y)
                            } else {
                                ball.updatePosition(json[i].d.x, json[i].d.y);
                            }
                            break;
                        case 4:
                            console.log('Player: X: ' + json[i].d.x +  ' Y: ' + json[i].d.y);
                            if(player === null){
                                player = new Pad(can, sock, json[i].d.x, json[i].d.y);
                            } else {
                                player.setPos(json[i].d.x, json[i].d.y);
                            }
                            break;
                        case 5:
                            //console.log('Opponent: X: ' + json[i].d.x +  ' Y: ' + json[i].d.y);
                            if(opponent === null){
                                opponent = new Pad(can, null, json[i].d.x, json[i].d.y);
                            } else {
                                opponent.setPos(json[i].d.x);
                            }
                    }
                }
            }
        }
    }
    if (can.getContext) {
        setInterval(function() {
            if (player !== null && ball !== null /*&& opponent !== null*/) {
                can.getContext('2d').clearRect(0, 0, can.width, can.height);
                ball.draw();
                player.draw(can.height);
                //opponent.draw(can.height);
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