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
U - URL of the web page
```

## EXAMPLES

Download web page with assets. The downloaded files will be saved in `downloads` directory.

```
$ html-source -U "https://enindu.com"
```

## TODO

- Download assets from inline styles
- Download assets in style files
