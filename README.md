# Oompa Loompa

![](http://i.imgur.com/6CvKzJ9.jpg)

Ooompa Loompa is a simple web service that listens to POST hooks from Bitbucket or GitHub. We use it for fully automatic deployment of static HTML sites.

## Usage

    $ oompaloompa --help
    Usage of oompaloompa:
      -config="conf.json": config file to load
      -logfile="": log to file
      -port=4000: listen on port


**config.json** contains a list of each project with the repo name and the path. When a payload is received for a given project Oompa Loompa will run a `git pull` in that directory.

    {"projects": [
      {
        "name": "example",
        "path": "/var/www/www.example.com/public_html"
      },
      {
        "name": "exshmample",
        "path": "/var/www/www.exshmample.io/public_html"
      }
    ]}

Now set up a new webhook for your GitHub or Bitbucket repo to fire at your Ooompa Loompa host, e.g.

![](http://i.imgur.com/sq5cN70.png)

That's it!
