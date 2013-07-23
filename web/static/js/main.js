function general_drawer(response_string) {
	console.log(response_string);
	data = JSON.parse(response_string);
	console.log(data);
	if (data["type"] == "zone") {
		draw_zone(data);
	} else if (data["type"] == "battle") {
		draw_battle(data);
	}
}

function draw_zone(players) {
	clearCanvas();
	people = players.people;
	for (var i = 0; i < people.length; i++) {
		draw_square(people[i].x,people[i].y);
	}
}

function draw_battle(response_string) {
	clearCanvas();
	console.log("Drawing battle");
	draw_square(0,3);
	draw_square(8,0);
}

function zone_view() {
	ajax_request("type=zone&what=view",general_drawer,"/ajax/");
}

function move(direction) {
	ajax_request("type=zone&what=move&who=Quin&where="+direction,general_drawer,"/ajax/");
}

function battle_view() {
	ajax_request("type=battle&what=view&who=Quin",general_drawer,"/ajax/");
}

function finish() {
	ajax_request("type=battle&what=finish&who=Quin",general_drawer,"/ajax/");
}
