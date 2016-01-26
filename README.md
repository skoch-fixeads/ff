# Find File #

Find File or text in file with go language.

-----------------------
### Source Find File ###

* Find text or file name 
* Version: 0.0.1
* License: ISC

-----------------------

## Parameters 

##### Helper: 

```
$ ff -h
Usage of ./ff:
  -a int
    	Range around of the word that I found. (default 10)
  -d	Show Debug Mode
  -f string
    	Filter by name of file or the file name
  -no-color
    	Disable color output
  -p string
    	The directory path (default "./")
  -r	Search by regex
  -t string
    	what I am searching
```

-----------------------

## Usage

Without parameters will show everything in subfolders.

	$ ff
	----------------------------------------------------------------------------------------------------
	Path : ./
	----------------------------------------------------------------------------------------------------
	[File] ./ 
	[File] .git 
	[File] .git/COMMIT_EDITMSG 
	[File] .git/FETCH_HEAD 
	[File] .git/HEAD 
	.....


-----------------------

### Parameter -f

Parameter `-f` of file, this parameter you will define the name of file  that you should like to filter


	$ ff -f css
	----------------------------------------------------------------------------------------------------
	Path : ./
	File : css
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css 
	[File] lookingFor/resources/css/angular-bootstrap-datepicker.css 
	[File] lookingFor/resources/css/bootstrap3.3.2.min.css 
	[File] lookingFor/resources/css/font-awesome.css 
	....


-----------------------
### Parameter -t

Parameter `-t` of text, this text that you will search inside of file.
	

	$ ff -t rlopes
	----------------------------------------------------------------------------------------------------
	Path : ./
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css/toastr.css 
		[144] .rlopes { lorem i
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------


-----------------------

### Parameter -r

With parameter `-r` as regex will defined that your `-t` is a regular expression


	$ ff -r -t "(rlopes)|(consectetur)"
	----------------------------------------------------------------------------------------------------
	Path : ./
	Regex: (rlopes)|(consectetur)
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css/none.txt 
		[1] sit amet, consectetur adipisici
		[4] sit amet, consectetur adipisici
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css/toastr.css 
		[144] .rlopes { lorem i
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------


-----------------------

### Parameter Combination

you can match both of the parameters as `-t` and `-f` to create more complex search


	$ ff -t "rlopes" -f ini
	----------------------------------------------------------------------------------------------------
	Path : ./
	File : ini
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------


-----------------------

### Parameter -d 

The `-d` parameter will active the debug mode, you can see the time of searching (more things coming soon)


	$ ff -t "rlopes" -f ini -d
	----------------------------------------------------------------------------------------------------
	Path : ./
	File : ini
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------
	final Execution took 2.031182ms


-----------------------

### Parameter -v

Show the version of application

-----------------------

### Parameter -h

Show the helper and parameters available

-----------------------

## Time Execution
	
The all tests was made in the same computer 

-----------------------
The `ff` running looking for word `rlopes`

	$ time ff -t rlopes
	----------------------------------------------------------------------------------------------------
	Path : ./
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] resources/css/toastr.css 
		[144] .rlopes { lorem i
	----------------------------------------------------------------------------------------------------
	[File] services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------
	ff -t rlopes  0,02s user 0,00s system 75% cpu 0,032 total


-----------------------
The `ack` running looking for word `rlopes`
	
	$ time ack rlopes
	resources/css/toastr.css
	144:.rlopes { lorem ipsum }

	services.dev.ini
	3:domain = ".rlopes.realestateid.fixe"
	6:baseUri        = 'http://rlopes.realestateid.fixe'
	15:baseUri        = 'http://rlopes.realestateid.fixe'
	ack .rlopes  0,06s user 0,02s system 93% cpu 0,077 total


-----------------------
The `grep` running looking for word `rlopes`

	$ time grep -i -E -r 'rlopes' ./*
	./resources/css/toastr.css:.rlopes { lorem ipsum }
	./services.dev.ini:domain = ".rlopes.realestateid.fixe"
	./services.dev.ini:baseUri        = 'http://rlopes.realestateid.fixe'
	./services.dev.ini:baseUri        = 'http://rlopes.realestateid.fixe'
	grep --color=auto --exclude-dir={.bzr,CVS,.git,.hg,.svn} -i -E -r 'rlopes' ./  0,30s user 0,00s system 98% cpu 0,308 total


-----------------------

## How to Compile it

 compile files from different platforms

### requirements 

* Go 

#### compiling to *Linux*

	$ env GOOS=linux GOARCH=arm GOARM=7 go build -o ff main.go


#### compiling to *MacOs*

	$ env GOOS=darwin GOARCH=386 go build -o ff main.go


#### compiling to *Windows*
	
	$ env GOOS=windows GOARCH=386 go build -o ff.exe main.go


