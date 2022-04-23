# HTML Source

HTML source code extractor, written in Go.

## Build

Use Go compiler to build a executable.

```
$ go build
```

## Usage

There is one flag, you can use with the `HTML Source`.

```
-u URL of the web page
```

## Examples

Print the source code given web page in terminal.

```
$ html-source -u "https://enindu.com"
```

Print the source code of given web page to a file.

```
$ html-source -u "https://enindu.com" > enindu.com.html
```
