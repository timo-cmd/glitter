# glitter 
A leightweight (5kb) and cute CLI to manage git-users painless. Switch between multiple Git users to manage your projects. Simply grab the source code from the current release branch, install golang and fire up an administrative shell. Type this commands to install the project via the golang package manager

```shell
go get github.com/timo-cmd/glitter
```

#### Usage:
```
Usage: 
  glitter [flags]  

Flags:
  --global              Set user as global.
```

#### Project Content:

```
./github.com/timo-cmd/glitter
├───sources
|   ├───main.go
│   ├───config.go
│   └───user.go
├───binaries
├───.gitignore
├───go.mod
├───go.sum
├───LICENSE
└───REAMDE
```

#### License:

```
MIT License

Copyright (c) 2020 Timo S. 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

#### Thanks to:

- Golang.org for their absolutely awesome programming language
- promptui for their brilliant project
- Ansiterm for their cool and stylish go-module

##### Developed by Timo Sarkar v0.1.2dev 09.12.20
