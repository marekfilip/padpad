function Ball(canvas, x, y) {
    this.x = x;
    this.y = y;
    this.context = canvas.getContext('2d');
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