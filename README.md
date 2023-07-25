![image](https://hub.steampipe.io/images/plugins/turbot/dockerhub-social-graphic.png)

# Docker Hub plugin for Steampipe

Use SQL to instantly query Docker Hub Repositories, Tags, Tokens and more. Open source CLI. No DB required.

- **[Get started ->](https://hub.steampipe.io/plugins/turbot/dockerhub)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/dockerhub/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-dockerhub/issues)

## Quick start

### Install

Download and install the latest Docker Hub plugin:

```shell
steampipe plugin install dockerhub
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/dockerhub#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/dockerhub#configuration).

Configure your account details in `~/.steampipe/config/dockerhub.spc`:

```hcl
connection "dockerhub" {
  plugin = "dockerhub"

  # Authentication information
  username = "turbot"
  password = "turbot@123"
}
```

Or through 2FA code:

```hcl
connection "dockerhub" {
  plugin = "dockerhub"

  # Authentication information
  username        = "turbot"
  password        = "turbot@123"
  two_factor_code = "123456"
}
```

Or through environment variables:

```sh
export DOCKER_HUB_USERNAME=turbot
export DOCKER_HUB_PASSWORD=turbot@123
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-dockerhub.git
cd steampipe-plugin-dockerhub
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/dockerhub.spc
```

Try it!

```
steampipe query
> .inspect dockerhub
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-dockerhub/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Docker Hub Plugin](https://github.com/turbot/steampipe-plugin-dockerhub/labels/help%20wanted)
