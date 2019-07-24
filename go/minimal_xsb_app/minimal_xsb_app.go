package main

/*
#cgo CFLAGS: -I/home/tmcphill/XSB/emu/
#cgo CFLAGS: -I/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/

#include <stdio.h>
#include <string.h>
#include "cinterf.h"

void initialize_xsb() {
	char init_string[1024];
	strcpy(init_string, "/home/tmcphill/XSB/");
	if (xsb_init_string(init_string) == XSB_ERROR) {
		fprintf(stderr, "Error initializing XSB: %s/%s\n",
			xsb_get_init_error_type(), xsb_get_init_error_message());
		exit(XSB_ERROR);
	}
}
*/
import "C"

func main() {
	C.initialize_xsb()
}
