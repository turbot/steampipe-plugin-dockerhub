---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/dockerhub.svg"
brand_color: "#096BD4"
display_name: "Docker Hub"
short_name: "dockerhub"
description: "Steampipe plugin for querying Docker Hub Repositories, Tags and other resources."
og_description: Query Docker Hub with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/dockerhub-social-graphic.png"
---

# Docker Hub + Steampipe

[Docker Hub](https://hub.docker.com/) is a cloud-based repository and distribution service provided by Docker that allows developers to store and share container images.

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL from [Turbot](https://turbot.com/)

List your Docker Hub Repositories:

```sql
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository;
```

```
+------------------------+------------+------------+------------+---------------------------+
| name                   | pull_count | star_count | is_private | last_updated              |
+------------------------+------------+------------+------------+---------------------------+
| souravthe/test         | <null>     | <null>     | false      | 2023-07-17T18:26:25+05:30 |
| souravthe/test-private | <null>     | <null>     | true       | 2023-07-17T18:32:56+05:30 |
+------------------------+------------+------------+------------+---------------------------+
```

## Documentation

- [Table definitions / examples →](https://hub.steampipe.io/plugins/turbot/dockerhub/tables)

## Quick start

### Install

Download and install the latest Docker Hub plugin:

```sh
steampipe plugin install dockerhub
```

### Credentials

| Item | Description                                                                                                                                                                                              |
| ---- |----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | You will require the Docker `username`` and `password``.                                                                                               |
| Permissions | NA                                                              |
| Radius | Each connection represents one Docker Hub user. |                                                                    |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/dockerhub.spc`).<br />2. Credentials specified in environment variables, e.g., `DOCKER_HUB_USERNAME` and `DOCKER_HUB_PASSWORD`. |

### Configuration

Installing the latest Docker Hub plugin will create a config file (`~/.steampipe/config/dockerhub.spc`) with a single connection named `dockerhub`:

Configure your account details in `~/.steampipe/config/dockerhub.spc`:

```hcl
connection "dockerhub" {
  plugin = "dockerhub"

  # DockerHub Username. Required.
  # This can also be set via the 'DOCKER_HUB_USERNAME' environment variable.
  # username = "turbot"

  # DockerHub Password. Required.
  # This can also be set via the 'DOCKER_HUB_PASSWORD' environment variable.
  # password = "turbot@123"

  # DockerHub 2FA Code. Required when 2FA is enabled.
  # two_factor_code = "123456"
}
```

Alternatively, you can also use the standard Docker Hub environment variables to configure your credentials **only if other arguments (`username`, `password`) are not specified** in the connection:

```sh
export DOCKER_HUB_USERNAME=turbot
export DOCKER_HUB_PASSWORD=turbot@123
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-dockerhub
- Community: [Slack Channel](https://steampipe.io/community/join)
