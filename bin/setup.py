#!/usr/bin/env python

import os

def main():
	with open(os.path.join(os.environ['HOME'],".bashrc"),"a") as f:
		f.write("\nexport GOPATH=$GOPATH:%s\n"%os.getcwd())

if __name__=='__main__':
	main()
