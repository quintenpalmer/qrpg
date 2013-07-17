var size = 50;
var canvasName = "mainCanvas";

function draw_square(x,y) {
	canvas = document.getElementById(canvasName);
	context = canvas.getContext("2d");
	context.font="22px arial";

	context.fillStyle="#996611";
	context.fillRect(x*size,y*size,size,size);
	/*
	context.lineWidth = 2;
	context.strokeStyle = "black";
	context.stroke();
	*/
}

function clearCanvas() {
	canvas = document.getElementById(canvasName);
	context = canvas.getContext("2d");
	context.save();
	context.setTransform(1, 0, 0, 1, 0, 0);
	context.clearRect(0, 0, canvas.width, canvas.height);
	context.restore();
}
