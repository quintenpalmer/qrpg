
function ajax_request(command,callback,url){
	var xmlhttp = new XMLHttpRequest();
	xmlhttp.onreadystatechange=function(){
		if(xmlhttp.readyState == 4 && xmlhttp.status == 200){
			callback(xmlhttp.responseText);
		}
	}
	xmlhttp.open("POST",url,true);
	//var csrftoken = get_cookie('csrftoken');
	//xmlhttp.setRequestHeader("X-CSRFToken",csrftoken);
	xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded")
	xmlhttp.send(command);
}
