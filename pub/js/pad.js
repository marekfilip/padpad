function Pad(canvas, sock, x, y) {
    this.sock = sock;
    this.length = canvas.width / 8;
    //this.x = canvas.width / 2 - this.length / 2;
    this.x = x;
    //this.y = canvas.height - 20;
    this.y = y;
    this.canvas = canvas;
    this.context = canvas.getContext('2d');
    this.canvasHeight = canvas.height;
    this.canvasWidth = canvas.width;
    this.draw = function() {
        this.context.fillStyle = '#FF0000';
        this.context.beginPath();
        this.context.fillRect(this.x - this.length/2, this.y, this.length, 5);
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