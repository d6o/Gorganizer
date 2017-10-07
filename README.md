# ![Gorganizer](https://rawgit.com/DiSiqueira/Gorganizer/master/gorganizer-logo-50.jpg)

# Gorganizer ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/Gorganizer) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

Gorganizer is a Go program inspired by [Bhrigu Srivastava][bhrigu123] [Classifier Project][classifier].

The Gorganizer's goal is to be a perfect tool providing a stupidly easy-to-use and fast program to organize your files based on its extension.

[bhrigu123]: https://github.com/bhrigu123
[classifier]: https://github.com/bhrigu123/classifier

## Project Status

Gorganizer is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/Gorganizer#social-coding)

![](https://i.imgur.com/2rFfn9i.gif)
![](https://i.imgur.com/AkgCeMx.jpg)

## Features

- MORE THAN 60 DEFAULT EXTENSIONS!!!
- It's perfect to organize your DOWNLOADS FOLDER
- Instantly organize your files
- CUSTOMIZE to your needs
- EASY to add rules
- Easy to delete default rules
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/Gorganizer#usage)
- Very fast start up and response time
- Uses native libs
- Option to organize your files
- Preview changes before moving
- Language support (English, Portuguese and Turkish)

## Installation

### Option 1: Go Get

```bash
$ go get github.com/DiSiqueira/Gorganizer
$ Gorganizer -h
```

### Option 2: From source

```bash
$ go get gopkg.in/ini.v1
$ git clone https://github.com/DiSiqueira/Gorganizer.git
$ cd Gorganizer/
$ go build *.go
```

## Usage

### Basic usage

```bash
# Organize your current directory
$ ./gorganizer
```

### Only preview, do not make change

```bash
# Prints a preview, but do not move
$ ./gorganizer -preview=true
```

### Recursive mode

```bash
$ ./gorganizer -recursive
```

### Specify language (Default: en)

```bash
# Set language to Turkish
$ ./gorganizer -language=tr
```

### Add new rule

```bash
# Add .py to Python folder
$ ./gorganizer -newrule=py:Python
```

### Delete existing rule

```bash
# Delete txt rule
$ ./gorganizer -delrule=txt
```

### Print all rules

```bash
# Print all rules
$ ./gorganizer -allrules=true
```

### Move organized files to another folder

```bash
# Run in current directory and move organized files to ~/Downloads
$ ./gorganizer -output=~/Downloads
```

### Run in other directory

```bash
# Run in ~/Downloads
$ ./gorganizer -directory=~/Downloads
```

### Run in other directory and send organized files to a organized one

```bash
# Run in ~/Downloads
$ ./gorganizer -directory=~/Downloads -output=~/Documents
```

### Show help

```bash
$ ./gorganizer -h
```

## Program Help

![](http://image.prntscr.com/image/a7f2e8071d3742cda44149ed9a7c2674.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/Gorganizer/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ go get gopkg.in/ini.v1
$ git clone --recursive git@github.com:DiSiqueira/Gorganizer.git
$ cd Gorganizer/
$ go run *.go
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/Gorganizer/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
