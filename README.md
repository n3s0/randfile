# Randfile

## Summary

Randfile is a Go commandline application that ingests a directory path and picks
a random file. This is a small project. But, it has its place in my setup.

For my purposes this is for learning and randomly picking a wallpaper PNG or
JPEG and outputting its full path for Sway or i3wm.

## Build from Source

Clone the randfile repository.

```sh
git clone https://github.com/n3s0/randfile.git
```

Build the application.

```sh
go build .
```

Move the application to wherever you wish. I'll work on making installation a
little easier at some point. Just need to learn how.

## Usage

Running the application is pretty straight forward. Run randfile with either the
-p or --path flags and it will provide a random file from that path.

```sh
randfile -p /path/to/directory
```

Provided is the output of the command.

```sh
/path/to/directory/file.png
```

## Issues

If there are issues. Please feel free to submit an issue describing the problem
and I'll work on a solution.

## Pull Requests

This is kind of a learning pet project for me. But, if you would like to submit
a pull request. I'm happy to review it.

