![Progress Badge](https://img.shields.io/badge/development-in%20progress-yellow)
<p align="center">
    <img width="128" height="128" src="images/doc.png" alt="centered image" />
</p>

<h1 align="center">Doc</h1>

This is a CLI tool written in Go to help you generate a design doc or tech spec and upload it to Google Docs. I created this tool after I wrote and article for the StackOverflow Blog about writing technical specs. You can read the article at this [link](https://stackoverflow.blog/2020/04/06/a-practical-guide-to-writing-technical-specs/). 

## Usage
> 🔺**Warning:** This tool is still under development and has not yet been fully tested. Use it at your own risk. 

To use this tool, first run:
```
git clone https://github.com/zaracooper/doc.git
```

I'm yet to release this tool and as such haven't integrated Google Auth yet, but if you're eager to try it out, you can create a new Cloud Platform project and automatically enable the Google Docs API [here](https://console.developers.google.com/projectcreate). Once done, download the client credentials and save them in a `credentials.json` file which you should put in project's root. 

Finally just run:
```
go run main.go
```

## TODO
- [ ] Write tests
- [ ] Set up CI
- [ ] Distribute this as an installable tool on various functions
