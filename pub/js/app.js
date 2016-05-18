var sock = null,
    button = null,
    can = null,
    ball = null,
    player = null,
    opponent = null;

function touchHandler(event)
{
    var touches = event.changedTouches,
        first = touches[0],
        type = "";
    switch(event.type)
    {
        case "touchstart": type = "mousedown"; break;
        case "touchmove":  type = "mousemove"; break;        
        case "touchend":   type = "mouseup";   break;
        default:           return;
    }

    // initMouseEvent(type, canBubble, cancelable, view, clickCount, 
    //                screenX, screenY, clientX, clientY, ctrlKey, 
    //                altKey, shiftKey, metaKey, button, relatedTarget);

    var simulatedEvent = document.createEvent("MouseEvent");
    simulatedEvent.initMouseEvent(type, true, true, window, 1, 
                                  first.screenX, first.screenY, 
                                  first.clientX, first.clientY, false, 
                                  false, false, false, 0/*left*/, null);

    first.target.dispatchEvent(simulatedEvent);
    event.preventDefault();
}

function init() 
{
    document.addEventListener("touchstart", touchHandler, true);
    document.addEventListener("touchmove", touchHandler, true);
    document.addEventListener("touchend", touchHandler, true);
    document.addEventListener("touchcancel", touchHandler, true);    
}

(function() {
    sock = new WebSocket("ws://"+document.location.hostname+":12345/handler");
    button = document.getElementById('start-game');
    can = document.getElementById('game');

    init();

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
                            if(player === null){
                                player = new Pad(can, sock, json[i].d.x, json[i].d.y, json[i].d.l);
                            } else {
                                player.setPos(json[i].d.x, json[i].d.y);
                                document.getElementById('player-points').innerHTML = json[i].d.p;
                            }
                            break;
                        case 5:
                            if(opponent === null){
                                opponent = new Pad(can, null, json[i].d.x, json[i].d.y, json[i].d.l);
                            } else {
                                opponent.setPos(json[i].d.x, json[i].d.y);
                                document.getElementById('opponent-points').innerHTML = json[i].d.p;
                            }
                    }
                }
            }
        }
    }
    if (can.getContext) {
        setInterval(function() {
            if (player !== null && ball !== null && opponent !== null) {
                can.getContext('2d').clearRect(0, 0, can.width, can.height);
                ball.draw();
                player.draw();
                opponent.draw();
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
    var el = document.getElementById('message-box'),
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