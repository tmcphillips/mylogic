## 2019-07-23 Statically link Go app to XSB

### Background
- XSB will be the first logic engine that MyLogic supports.
- By statically linking to XSB, MyLogic executables will provide within a single executable all of the logic capabilities of XSB without depending on any other software or libraries being installed at runtime.

### XSB Documentation Excerpts
- Chapter 14 of the [XSB Programmer Manual](http://xsb.sourceforge.net/manual1/manual1.pdf) is titled *Embedding XSB in a Process*.
-   Figure 14.1 (page 484) provides an example program (slightly modified below) that uses the XSB C API to consult (load facts from) a file, query the facts, and output the results:

    ```C
    #include <stdio.h>
    #include <string.h>
    
    /* cinterf.h is necessary for the XSB API, as well as the path manipulation routines*/
    #include "cinterf.h"
    
    extern char *xsb_executable_full_path(char *);
    extern char *strip_names_from_path(char*, int);
    
    int main(int argc, char *argv[]) {
    
    	char init_string[1024];
    	int rc;
    	XSB_StrDefine(return_string);
    	
    	/* xsb_init_string() relies on the calling program to pass the absolute or relative
    		path name of the XSB installation directory. We assume that the current
    		program is sitting in the directory ../examples/c_calling_xsb/
    		To get the installation directory, we strip 3 file names from the path. */
    	
    	strcpy(init_string, strip_names_from_path(xsb_executable_full_path(argv[0]),3));
    	
    	if (xsb_init_string(init_string) == XSB_ERROR) {
    		fprintf(stderr, "++initializing XSB: %s/%s\n", 
    			xsb_get_init_error_type(), xsb_get_init_error_message());
    		exit(XSB_ERROR);
    	}
    	
    	/* Create command to consult a file: edb.P, and send it. */
    	if (xsb_command_string("consult(’edb.P’).") == XSB_ERROR) {
    		fprintf(stderr, "++Error consulting edb.P: %s/%s\n", 
    			xsb_get_error_type(), xsb_get_error_message());
    	}
    	
    	rc = xsb_query_string_string("p(X,Y,Z).", &return_string, "|");
    	while (rc == XSB_SUCCESS) {
    		printf("Return %s\n", (return_string.string));
    		rc = xsb_next_string(&return_string, "|");
    	}
    
    	if (rc == XSB_ERROR) {
    		fprintf(stderr, "++Query Error: %s/%s\n", 
    			xsb_get_error_type(), xsb_get_error_message());
    	}
    	
    	xsb_close();
    }
    ```
- Section 14.6 (page 507) provides instructions for building a C program that invokes XSB:

	> To create an executable that includes calls to the above C functions, these routines,
	and the XSB routines that they call, must be included in the link (ld) step.
	>
	>Unix instructions: You must link your C program, which should include the main
	procedure, with the XSB object file located in
	>
	> `$XSBDIR/config/<your-system-architecture>/saved.o/xsb.o`
	>
	> Your program should include the file `cinterf.h` located in the `XSB/emu` subdirectory,
	which defines the routines described earlier, which you will need to use in order to
	talk to XSB. It is therefore recommended to compile your program with the option
	`-I$XSB_DIR/XSB/emu`.
	>
	> The file `$XSB_DIR/config/your-system-architecture/modMakefile` is a makefile
	you can use to build your programs and link them with XSB. It is generated
	automatically and contains all the right settings for your architecture, but you will
	have to fill in the name of your program, etc.

- The `modMakefile` referenced above is located at `/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu`, and has the following contents, including suggested linker flags:

    ```makefile
    ##############################################################################
    #                                                                            #
    # Makefile for compiling C programs that call XSB as a module                #
    # and for linking with that module                                           #
    #                                                                            #
    # You will need to edit this file to adjust to your particular program       #
    #                                                                            #
    ##############################################################################
    
    ## File:      modMakefile.in
    ## Author(s): kifer
    ## Copyright (C) The Research Foundation of SUNY, 1998
    ##
    ## XSB is free software; you can redistribute it and/or modify it under the
    ## terms of the GNU Library General Public License as published by the Free
    ## Software Foundation; either version 2 of the License, or (at your option)
    ## any later version.
    ##
    ## XSB is distributed in the hope that it will be useful, but WITHOUT ANY
    ## WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
    ## FOR A PARTICULAR PURPOSE.  See the GNU Library General Public License for
    ## more details.
    ##
    ## You should have received a copy of the GNU Library General Public License
    ## along with XSB; if not, write to the Free Software Foundation,
    ## Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
    ##
    ## $Id: modMakefile.in,v 1.6 2010-08-19 15:03:35 spyrosh Exp $
    ##
    
    ## DON'T EDIT FROM HERE TIL NEXT NOTICE
    ############################################################################
    
    # Where the architectue specific XSB stuff is
    arch_dir=/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu
    
    CC=gcc
    CPP=gcc -E
    CFLAGS=  -O3 -fno-strict-aliasing   -fPIC -Wall -pipe
    # Flags for files requiring lower optimization level (emuloop.c)
    CPPFLAGS=
    LDFLAGS=  -lm -ldl -Wl,-export-dynamic -lpthread
    
    # where the xsb object module is found
    xsbmodule=$(arch_dir)/saved.o/$(emumake_goal)module.o
    
    # just to be sure that sh is used
    SHELL=/bin/sh
    
    ###### END NO-EDIT ZONE
    
    ##################### You may need to edit some of the files below!
    
    # Where to install your C program that calls XSB binaries
    bindir=$(arch_dir)/bin
    
    
    # You will most likely need to edit this
    your_program: your_program.o
            gcc -o $(bindir)/your_program $(xsbmodule) your_program.o $(LDFLAGS)
            @echo ""
            @echo "***************************************************"
            @echo "The executable is in:"
            @echo "     $(bindir)/your_program"
            @echo ""
            @echo ""
    
    
    your_program.o: your_program.c
            gcc -c $(CFLAGS) your_program.c
    ```

### Link first mimimal Go app to XSB

- Wrote minimal Go app that includes the required XSB C API header file `cinterf.h`.  Both include directories given as CFLAGS are required:

    ```go
    package main
    
    /*
    #cgo CFLAGS: -I/home/tmcphill/XSB/emu/
    #cgo CFLAGS: -I/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/
    
    #include "cinterf.h"
    */
    import "C"
    
    func main() {
    }
    ```
- The above program builds without error with the following command:
	```terminal
	$ go build  -ldflags "-linkmode external -extldflags -static"
	```

- Added a C function to initialize XSB:

    ```go
    package main
    
    /*
    #cgo CFLAGS: -I/home/tmcphill/XSB/emu/
    #cgo CFLAGS: -I/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/
    
    #include <stdio.h>
    #include <string.h>
    #include "cinterf.h"
    
    void initialize_xsb() {
    	char init_string[1024];
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
    ```
- As expected, building the new version of the app gives linker errors:
	```terminal
	$ go build minimal_xsb_app.go
	# command-line-arguments
	/tmp/go-build431834133/b001/_x002.o: In function `initialize_xsb':
	./minimal_xsb_app.go:13: undefined reference to `xsb_init_string'
	./minimal_xsb_app.go:14: undefined reference to `xsb_get_init_error_message'
	./minimal_xsb_app.go:14: undefined reference to `xsb_get_init_error_type'
	collect2: error: ld returned 1 exit status
	```

- Created an Unix library containing the XSB object file:
	```terminal
	tmcphill@circe-win10:~/XSB/config/x86_64-unknown-linux-gnu/saved.o$ ls -al xsb.o
	-rw-rw-rw- 1 tmcphill tmcphill 1902592 Jul 22 19:12 xsb.o

	tmcphill@circe-win10:~/XSB/config/x86_64-unknown-linux-gnu/saved.o$ ar cvq libxsb.a xsb.o
	a - xsb.o
	```
- Set the CGO_LDFLAGS variable to enable linking to the XSB library file:
	```terminal
	export CGO_LDFLAGS="-g -O2 -L /home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/saved.o -lxsb -lm -ldl" 
	```
- Build the Go app *without* requiring static linking to system libraries:
	```terminal
	minimal_xsb_app$ go build
	# github.com/tmcphillips/mylogic/go/minimal_xsb_app
	/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/saved.o/libxsb.a(xsb.o): In function `sys_system':
	(.text+0x86ffa): warning: the use of `tempnam' is dangerous, better use `mkstemp'
	
	minimal_xsb_app$ ls -al
	total 4612
	drwxrwxrwx 1 tmcphill tmcphill    4096 Jul 23 17:51 .
	drwxrwxrwx 1 tmcphill tmcphill    4096 Jul 23 13:58 ..
	-rwxrwxrwx 1 tmcphill tmcphill 2322928 Jul 23 17:50 minimal_xsb_app
	-rw-rw-rw- 1 tmcphill tmcphill     521 Jul 23 15:53 minimal_xsb_app.go
	
	minimal_xsb_app$ ldd minimal_xsb_app
	        linux-vdso.so.1 (0x00007fffe04d7000)
	        libm.so.6 => /lib/x86_64-linux-gnu/libm.so.6 (0x00007f35578f0000)
	        libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f35576e0000)
	        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f35574c0000)
	        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f3557120000)
	        /lib64/ld-linux-x86-64.so.2 (0x00007f3557e00000)	
	```
- All six of the dynamic dependencies are on shared system libraries.  The linking to XSB to static.

- Successfully ran the Go app that initializes XSB:
	```terminal
	$ ./minimal_xsb_app
	[xsb_configuration loaded]
	[sysinitrc loaded]
	[xsbbrat loaded]
	```

### Current limitations of statically linked executable

- Building a static binary succeeds but gives linker warnings: 
	```terminal
	$ go build -ldflags "-linkmode external -extldflags -static"
	# github.com/tmcphillips/mylogic/go/minimal_xsb_app

	/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/saved.o/libxsb.a(xsb.o): In function `load_obj_dyn.isra.0':
	dynload.c:(.text+0x296f0): warning: Using 'dlopen' in statically linked applications requires at runtime the shared libraries from the glibc version used for linking

	/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/saved.o/libxsb.a(xsb.o): In function `sys_system':
	(.text+0x86ffa): warning: the use of `tempnam' is dangerous, better use `mkstemp'

	/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/saved.o/libxsb.a(xsb.o): In function `tilde_expand_filename_norectify.part.0':
	pathname_xsb.c:(.text+0x79455): warning: Using 'getpwnam' in statically linked applications requires at runtime the shared libraries from the glibc version used for linking

	/home/tmcphill/XSB/config/x86_64-unknown-linux-gnu/saved.o/libxsb.a(xsb.o): In function `builtin_call':
	(.text+0x14636): warning: Using 'gethostbyname' in statically linked applications requires at runtime the shared libraries from the glibc version used for linking
	```
- Three of the warnings errors are due to XSB dependencies on `glibc` that in fact will require dynamic linking to the libraries at run time.
- Running the statically linked app does work on the same system where the binary was built and where XSB is installed:
	```
	$ ldd ./minimal_xsb_app
	        not a dynamic executable

	$ ./minimal_xsb_app
	[xsb_configuration loaded]
	[sysinitrc loaded]
	[xsbbrat loaded]
	```
- However, the app only works if the path to the XSB installation is provided to the XSB initialization function `xsb_init_string()`.  Passing an erroneous path to this function (e.g. `/XSB/` instead of `/home/tmcphill/XSB/` gives a run time error indicating that a configuration file is expected to be found in the XSB installation tree:
	```terminal
	$ ./minimal_xsb_app
	Error initializing XSB: init_error/XSB configuration file /XSB/config/x86_64-unknown-linux-gnu/lib/xsb_configuration.P does not exist or is not readable by you.
	```
- More research is needed to determine what other run-time dependencies on the XSB installation this implies.
