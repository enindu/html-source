# HTML SOURCE

HTML source code extractor, written in Go.

## BUILD

Use Go compiler to build a executable.

```
$ go build
```

## USAGE

There is one flag, you can use with the HTML Source.

```
-U - URL of the web page
```

## EXAMPLES

Print the source code given web page in terminal.

```
$ html-source -U "https://enindu.com"
```

Print the source code of given web page to a file.

```
$ html-source -U "https://enindu.com" > enindu.com.html
```

## TODO

- Download assets from inline styles
- Download assets in style files
