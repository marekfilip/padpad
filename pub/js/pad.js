function Pad(canvas, sock, x, y, l) {
    this.sock = sock;
    this.context = canvas.getContext('2d');

    this.x = x;
    this.y = y;
    this.length = l;

    this.draw = function(canW) {
        this.context.fillStyle = '#FF0000';
        this.context.beginPath();
        this.context.fillRect(this.x, this.y, this.length, 5);
        this.context.fill();
    };
    this.updatePos = function(x) {
        if(sock !== null){
            this.sock.send(JSON.stringify({
                't': 2,
                'd': {
                    'pX': x
                }
            }));
        }
    };
    this.setPos = function(x, y) {
        this.x = x;
        this.y = y;
    }
    this.getPos = function() {
        return {
            xLeft: this.x,
            xRight: (this.x + this.length),
            y: this.y
        }
    };
};