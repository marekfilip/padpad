function Ball(canvas) {
    this.x = 0;
    this.y = 0;
    this.dirX = 1;
    this.dirY = 1;
    this.angleX = 1;
    this.angleY = 1;
    this.context = canvas.getContext('2d');
    this.canvasHeight = canvas.height;
    this.canvasWidth = canvas.width;
    this.speed = 1;
    this.draw = function() {
        this.context.fillStyle = '#00FF00';
        this.context.beginPath();
        this.context.arc(this.x, this.y, 7, 0, 2 * Math.PI, false);
        this.context.fill();
    };
    this.updatePosition = function(x, y) {
        this.x = x;
        this.y = y;
    }
};
/*Ball.prototype.move = function(padpos) {
    if ((padpos.y >= (Math.round(this.y) + 4) && padpos.y <= (Math.round(this.y) + 7)) && this.x >= padpos.xLeft && this.x <= padpos.xRight) {
        this.dirY = -1;
        this.speed += 0.05;
    }
    if (this.x >= (this.canvasWidth - 7)) {
        this.dirX = -1;
        this.speed += 0.05;
    };
    if (this.x <= 7) {
        this.dirX = 1;
        this.speed += 0.05;
    }
    if (this.y <= 7) {
        this.dirY = 1;
        this.speed += 0.05;
    }
    this.x = this.x + this.angleX * this.dirX * this.speed;
    this.y = this.y + this.angleY * this.dirY * this.speed;
    if (this.x < 7) this.x = 7;
    if (this.y < 7) this.y = 7;
    if (this.x > (this.canvasWidth - 7)) this.x = (this.canvasWidth - 7);
};*/