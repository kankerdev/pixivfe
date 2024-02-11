> # Important!
> ### This is an unofficial mirror of the original project on [CodeBerg](https://codeberg.org/vnpower/pixivfe).
> Upstream lacks a ready to use container image, so this fork will provide one hosted on [ghcr.io](https://github.com/kankerdev/searxng/pkgs/container/pixivfe). 
> There's little to no changes from upstream to the repo itself so it can be considered to be a mirror. 
>
> For bug-reports, contributions or other issues you should be contributing them to upstream. Not this repo!

# PixivFE

A privacy-respecting alternative front-end for Pixiv that doesn't suck

<p>
<a href="https://codeberg.org/vnpower/pixivfe">
<img alt="Get it on Codeberg" src="https://get-it-on.codeberg.org/get-it-on-blue-on-white.png" height="60">
</a>
</p>

![CI badge](https://ci.codeberg.org/api/badges/12556/status.svg)
[![Go Report Card](https://goreportcard.com/badge/codeberg.org/vnpower/pixivfe)](https://goreportcard.com/report/codeberg.org/vnpower/pixivfe)

Questions? Feedbacks? You can [PM me](https://matrix.to/#/@vnpower:eientei.org) on
Matrix!

You can keep track of this project's development
[here](https://codeberg.org/VnPower/PixivFE/wiki/Things-to-do).

## Features

- Lightweight - both the interface and the code
- Privacy-first - the server will do the work for you
- No bloat - we only serve HTML and CSS
- Open source - you can trust me!

## Hosting

You can use PixivFE for personal use! Assuming that you use an operating system that can run POSIX shell scripts, install `go`, clone this repository, modify the `run.sh` file, and profit!
I recommend self-hosting your own instance for personal use, instead of relying entirely on official instances.


Check out [this page](https://codeberg.org/VnPower/pixivfe/wiki/Hosting). We
currently have guides for Docker and Caddy.

## Development

```
# watch and compile styles with node-sass
pnpm i -g sass
sass --watch views/css

# run in development mode (auto reload templates)
PIXIVFE_DEV=1 ... go run .
```

## Instances

| Name               | Cloudflare? | URL                             |
|--------------------|-------------|---------------------------------|
| exozyme (Official) | No          | https://pixivfe.exozy.me        |
| dragongoose        | No          | https://pixivfe.drgns.space     |
| chaotic.ninja      | No          | https://pix.chaotic.ninja       |
| WhateverItWorks    | Yes         | https://art.whateveritworks.org |
| ducks.party        | No          | https://pixivfe.ducks.party     |

Hosted one yourself? Create a pull request to add it here!

## License & Attributions

License: [AGPL3](https://www.gnu.org/licenses/agpl-3.0.txt)

Special thanks:
- [huggy](https://huggy.moe): author of [ugoira.com](https://ugoira.com) for the ugoira API
- [dragongoose](https://drgns.space): writing guides
- Contributors, stargazers and users like you, as well! 
