## v0.1.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#11](https://github.com/turbot/steampipe-plugin-dockerhub/pull/11))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#9](https://github.com/turbot/steampipe-plugin-dockerhub/pull/9))
- Recompiled plugin with Go version `1.21`. ([#9](https://github.com/turbot/steampipe-plugin-dockerhub/pull/9))

## v0.0.1 [2023-07-26]

_What's new?_

- New tables added
  - [dockerhub_repository](https://hub.steampipe.io/plugins/turbot/dockerhub/tables/dockerhub_repository)
  - [dockerhub_tag](https://hub.steampipe.io/plugins/turbot/dockerhub/tables/dockerhub_tag)
  - [dockerhub_token](https://hub.steampipe.io/plugins/turbot/dockerhub/tables/dockerhub_token)
  - [dockerhub_user](https://hub.steampipe.io/plugins/turbot/dockerhub/tables/dockerhub_user)
