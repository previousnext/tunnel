Tunnel
======

Simple tunneller for passing ports from one point, to another.

## Usage

The primary use case for this project is to tunnel external connections from your work/home network to an internal host (VM / Container) on your workstation.

In this example we need to forward port 80, 3000 and 3001 to an internal vm with the internal DNS (project.dev).

```bash
$ sudo tunnel 80,3000,3001 project.dev
0.0.0.0:80 -> project.dev:80
0.0.0.0:3000 -> project.dev:3000
0.0.0.0:3001 -> project.dev:3001
```
