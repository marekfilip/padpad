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
        this.context.fillRect(this.x, this.y, this.length, 5);
        this.context.fill();
    };
    this.updatePos = function(x) {
        this.x = x - this.length / 2;
        if (this.x < 0) this.x = 0;
        if ((this.x + this.length) > this.canvasWidth) this.x = this.canvasWidth - this.length;
        if(sock !== null){
            this.sock.send(JSON.stringify({
                't': 2,
                'd': {
                    'pX': this.x
                }
            }));
        }
    };
    this.getPos = function() {
        return {
            xLeft: this.x,
            xRight: (this.x + this.length),
            y: this.y
        }
    };
};