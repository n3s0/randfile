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
go build 
```

Move the application to wherever you wish. I'll work on making installation a
little easier at some point. Just need to learn how.

I generally put it in the /usr/local/bin using the following command.

```sh
sudo cp ./randfile /usr/local/bin/randfile
```

## Future Deployment Options

In the future I intend on either putting this into package repos for different
operating systems. This includes Debian/Ubuntu, CentOS/Fedora, Arch Linux, etc.
Although I don't think this would be entirely useful to everyone. Nor is it
popular enough to just deploy. I would like to learn how to package and deploy
open-source software. So, I thought this would be a good start to learning how
to do that. Considering how these different distros have their own flavor of
packaging and deploying software.

The following list contains a list of distrobutions I would like to attempt to
package this application for. This will also be done in various ARM and x86_64
architectures.

- [ ] Arch Linux
- [ ] Debian
- [ ] Ubuntu
- [ ] Fedora

I learn by doing.

## Usage

Running the application is pretty straight forward. Run randfile with either the
-p or --path flags and it will provide a random file from that path.

```sh
randfile -p /path/to/directory
```

OR (if it's not in the file path)

```sh
./randfile -p /path/to/directory
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

