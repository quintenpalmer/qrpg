QRPG - Quinten's RPG Engine
===========================

Installation
------------
To install Go-1.1 read up on your Linux distro's package manager
or visit

	http://golang.org/doc/install

How to run
----------
From the root of this directory:

	./bin/setup.py

then:

	source ~/.bashrc

To run the backend:

	go run backend.go

To run the web server:

	go run frontend.go

Introduction
------------
There are three components:

Game Logic Server
-----------------
This is a Go server and does all of the backend logic.
It will have the world that the rpg takes place in
as well as all of the players and battles and zones
and monsters and all of that!

Web Server
----------
This is a Go server and deals with clients and
forwards their requests to the game logic server.
It is meant to be the only client facing part.

HTML5 Frontend
--------------
This is all HTML/CSS/Javascript and displays the game
to the user.


Package Structure
=================
bin/                Helper scripts  
src/                Go source code  
src/qrpg/           Game Logic Server code  
src/qrpgweb/       Web Server code  
web/                Web assets, static content, and templates  
web/templates/      Templates  
web/static/         Static html5 content  
web/static/js       Javascript code  
web/static/css      CSS  
web/static/images   Image assets  
backend.go          The "main" for the game logic server  
frontend.go         The "main" for the web server  
