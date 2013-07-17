function update_canvas(response_string) {
	clearCanvas();
	players = JSON.parse(response_string);
	people = players.people;
	for (var i = 0; i < people.length; i++) {
		draw_square(people[i].x,people[i].y);
	}
}

function view() {
	ajax_request("type=zone",update_canvas,"/ajax/");
}

function move(direction) {
	ajax_request("type=move&what="+direction,update_canvas,"/ajax/");
}
