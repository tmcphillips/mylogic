## 2019-07-22 Automate configuration of MyLogic development environment

### Defined an Ansible playbook for installing MyLogic dependencies
- First version of playbook `debian/playbooks/roles/mylogic-dev.yml`:
	```yaml
	---

	- name: install tools and dependencies for MyLogic system development
	  hosts: 127.0.0.1
	  connection: local
	  roles:
	    - git
	    - golang-1.12
	    - x11
	    - fyne
	    - xsb
	```
- Ran playbook:

    ```terminal
    (.venv-ansible-playbooks) tmcphill@circe-win10:~/GitRepos/ansible-playbooks/debian/playbooks$ ansible-playbook -K mylogic-dev.yml
    BECOME password:
     [WARNING]: No inventory was parsed, only implicit localhost is available
    
     [WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'
    
    
    PLAY [install tools and dependencies for MyLogic system development] ***********************************************************************************************************************
    
    TASK [Gathering Facts] *********************************************************************************************************************************************************************
    ok: [127.0.0.1]
    
    TASK [git : install git] *******************************************************************************************************************************************************************
     [WARNING]: Could not find aptitude. Using apt-get instead
    
    ok: [127.0.0.1]
    
    TASK [git : configure global git settings] *************************************************************************************************************************************************
    changed: [127.0.0.1]
    
    TASK [golang-1.12 : delete existing installation of Go] ************************************************************************************************************************************
    changed: [127.0.0.1]
    
    TASK [golang-1.12 : download and expand Go 1.12.7] *****************************************************************************************************************************************
    changed: [127.0.0.1]
    
    TASK [golang-1.12 : create and set contents of an initializer script to be run by bash at login] *******************************************************************************************
    ok: [127.0.0.1]
    
    TASK [x11 : install basic X11 applications] ************************************************************************************************************************************************
    ok: [127.0.0.1]
    
    TASK [x11 : create and set contents of an X11 initializer script to be run by bash at login] ***********************************************************************************************
    ok: [127.0.0.1]
    
    TASK [fyne : install X11 dependency] *******************************************************************************************************************************************************
    ok: [127.0.0.1]
    
    TASK [fyne : install OpenGL dependency] ****************************************************************************************************************************************************
    ok: [127.0.0.1]
    
    TASK [fyne : install the Fyne API for Go GUI development] **********************************************************************************************************************************
    changed: [127.0.0.1]
    
    TASK [xsb : download and expand XSB 3.8] ***************************************************************************************************************************************************
    changed: [127.0.0.1]
    
    TASK [xsb : configure and build XSB] *******************************************************************************************************************************************************
    changed: [127.0.0.1]
    
    TASK [xsb : create and set contents of script adding XSB bin directory to $PATH at loginn] *************************************************************************************************
    changed: [127.0.0.1]
    
    PLAY RECAP *********************************************************************************************************************************************************************************
    127.0.0.1                  : ok=14   changed=7    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    ```
- In new terminal confirmed that `Go 1.12` is installed and environment looks ok:
    ```terminal
	tmcphill@circe-win10:~$ which go
	/usr/local/go/bin/go
	
	tmcphill@circe-win10:~$ go version
    go version go1.12.7 linux/amd64
    
    tmcphill@circe-win10:~$ go env
    GOARCH="amd64"
    GOBIN=""
    GOCACHE="/home/tmcphill/.cache/go-build"
    GOEXE=""
    GOFLAGS=""
    GOHOSTARCH="amd64"
    GOHOSTOS="linux"
    GOOS="linux"
    GOPATH="/home/tmcphill/go"
    GOPROXY=""
    GORACE=""
    GOROOT="/usr/local/go"
    GOTMPDIR=""
    GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
    GCCGO="gccgo"
    CC="gcc"
    CXX="g++"
    CGO_ENABLED="1"
    GOMOD=""
    CGO_CFLAGS="-g -O2"
    CGO_CPPFLAGS=""
    CGO_CXXFLAGS="-g -O2"
    CGO_FFLAGS="-g -O2"
    CGO_LDFLAGS="-g -O2"
    PKG_CONFIG="pkg-config"
    GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build972424997=/tmp/go-build -gno-record-gcc-switches"
    ```
- Noted environment settings for `CGO` above.  Checked versions of `gcc` and `g++` installed:

	```terminal
	tmcphill@circe-win10:~$ gcc --version
	gcc (Debian 6.3.0-18+deb9u1) 6.3.0 20170516
	Copyright (C) 2016 Free Software Foundation, Inc.
	This is free software; see the source for copying conditions.  There is NO
	warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

	tmcphill@circe-win10:~$ g++ --version
	g++ (Debian 6.3.0-18+deb9u1) 6.3.0 20170516
	Copyright (C) 2016 Free Software Foundation, Inc.
	This is free software; see the source for copying conditions.  There is NO
	warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
	```
- Confirmed that `fyne` is installed in `~/go`:
	```terminal
	tmcphill@circe-win10:~$ tree -L 3 ~/go
	/home/tmcphill/go
	├── pkg
	│   └── linux_amd64
	│       └── fyne.io
	└── src
	    └── fyne.io
	        └── fyne

	6 directories, 0 files
	```
- Demonstrated that XSB 3.8 works:
	```terminal
	tmcphill@circe-win10:~$ xsb --version
	XSB Version 3.8.0 (Three-Buck Chuck) of October 28, 2017
	[x86_64-unknown-linux-gnu 64 bits; mode: optimal; engine: slg-wam; scheduling: local]
	[Build date: 2019-07-22]
	
	XSB is licensed under the GNU Library General Public License.
	You can change it and/or distribute copies of it under certain conditions.
	You should have received the License with this distribution of XSB.
	If not, see:  http://www.gnu.org/copyleft/lgpl.html
	XSB comes without warranty of any kind.

	tmcphill@circe-win10:~$ xsb
	[xsb_configuration loaded]
	[sysinitrc loaded]
	[xsbbrat loaded]

	XSB Version 3.8.0 (Three-Buck Chuck) of October 28, 2017
	[x86_64-unknown-linux-gnu 64 bits; mode: optimal; engine: slg-wam; scheduling: local]
	[Build date: 2019-07-22]

	| ?- halt.

	End XSB (cputime 0.02 secs, elapsetime 4.40 secs)
	```
- Checked that XSB source distribution location is defined by the `XSB_DIR`  environment variable:

    ```terminal
    tmcphill@circe-win10:~$ echo $XSB_DIR
    /home/tmcphill/XSB
    
    tmcphill@circe-win10:~$ tree -L 1 $XSB_DIR
    /home/tmcphill/XSB
    ├── admin
    ├── bin
    ├── build
    ├── cmplib
    ├── config
    ├── docs
    ├── emu
    ├── etc
    ├── examples
    ├── FAQ
    ├── gpp
    ├── installer
    ├── InstallXSB.jar
    ├── lib
    ├── LICENSE
    ├── Makefile
    ├── packages
    ├── prolog-commons
    ├── prolog_includes
    ├── README
    ├── site
    └── syslib
    ```

### Confirmed that Go demos work 

- Manually cloned tmcphillips/go-demos repo into `~/go/src/github.com/tmcphillips/go-demos/` so that SSH protocol is used:
	```
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips$ git clone git@github.com:tmcphillips/go-demos.git
	Cloning into 'go-demos'...
	remote: Enumerating objects: 8, done.
	remote: Counting objects: 100% (8/8), done.
	remote: Compressing objects: 100% (5/5), done.
	remote: Total 236 (delta 1), reused 8 (delta 1), pack-reused 228
	Receiving objects: 100% (236/236), 43.26 KiB | 0 bytes/s, done.
	Resolving deltas: 100% (103/103), done.
	
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips$ cd go-demos/
	
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips/go-demos$ ls -l
	total 0
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 00_hello_world
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 01_hello_with_args
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 02_hello_with_flags
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 03_hamming
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 04_hello_qt
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 05_hello_andlabs
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 06_hello_nk
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 07_hello_fyne
	drwxrwxrwx 1 tmcphill tmcphill 4096 Jul 22 21:07 notes	
	```

- Successfully ran first demo and tests:
	```terminal
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips/go-demos$ cd 00_hello_world/
	
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips/go-demos/00_hello_world$ go run hello.go
	Hello World
	
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips/go-demos/00_hello_world$ go test
	PASS
	ok      github.com/tmcphillips/go-demos/00_hello_world  0.004s
	```

- Successfully ran hamming demo:
	```terminal
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips/go-demos/00_hello_world$ cd ../03_hamming/
	
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips/go-demos/03_hamming$ go run hamming.go
	1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20
	```

- Successfully ran Fyne demo, saw Fyne demo window appear on desktop:
	```terminal
	tmcphill@circe-win10:~/go/src/github.com/tmcphillips/go-demos/07_hello_fyne$ go run hello_fyne.go
	```

