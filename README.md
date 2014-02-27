# Oompa Loompa

Ooompa Loompa is a simple web service that listens to POST hooks from Bitbucket or GitHub. We use it for fully automatic deployment of static HTML sites.

**config.json** contains a list of each project with the repo name and the path. When a payload is received for a given project Oompa Loompa will run a `git pull` in that directory.

    {"projects": [
      {
        "name": "example",
        "path": "/var/www/www.example.com/public_html",
      },
      {
        "name": "exshmample",
        "path": "/var/www/www.exshmample.io/public_html",
      }
    ]}
