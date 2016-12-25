# ![gitql](https://rawgit.com/DiSiqueira/Gorganizer/master/gorganizer-logo.png)

# Gorganizer ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Dependencies Badge](https://img.shields.io/badge/Dependencies-BoltDB-brightgreen.svg) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

====

Gorganizer is a Go program inspired by [Bhrigu Srivastava][bhrigu123] [Classifier Project][classifier].

The Gorganizer goal is to be a perfect tool providing a stupidly easy to use and fast program to organize your files based on its extension.

[bhrigu123]: https://github.com/bhrigu123
[classifier]: https://github.com/bhrigu123/classifier

## Project Status

Gorganizer is on beta

Bolt is stable, more features will be added and full unit tests will be made.
Pull Requests [are welcome](https://github.com/DiSiqueira/Gorganizer#social-coding)

![](https://i.imgur.com/2rFfn9i.gif)

## Features

- It's perfect to organize your downloads folder
- Instantly organize your files
- Customize to your needs
- Easy to add rules
- Easy to delete default rules
- Stupidly [easy to use](https://github.com/DiSiqueira/Gorganizer#usage)
- Very fast start up and response time
- Uses natives libs
- Only one dependency - [BoltDB](https://github.com/boltdb/bolt)
- Option to organize your files

## Installation

### Option 1: Download binary

```bash
$
```

### Option 2: From source

```bash
$ go get github.com/boltdb/bolt/...
$ git clone https://github.com/DiSiqueira/Gorganizer.git
$ cd Gorganizer/src
$ go build *.go
```

## Usage

### Basic usage

```bash
# Organize your current directory
$ ./gorganizer
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
# Delete txt rule
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

![](http://image.prntscr.com/image/27c361f3891c461d83584577eb18ec72.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/Gorganizer/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ go get github.com/boltdb/bolt/...
$ git clone --recursive git@github.com:DiSiqueira/Gorganizer.git
$ cd Gorganizer/src/
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

Copyright (c) 2013-2015 Diego Siqueira

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