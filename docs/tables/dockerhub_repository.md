---
title: "Steampipe Table: dockerhub_repository - Query DockerHub Repositories using SQL"
description: "Allows users to query DockerHub Repositories, specifically the repository details including name, description, star count, pull count, and last updated date."
---

# Table: dockerhub_repository - Query DockerHub Repositories using SQL

DockerHub is a cloud-based registry service that allows you to link to code repositories, build your images and test them, stores manually pushed images, and links to Docker Cloud so you can deploy images to your hosts. It provides a centralized resource for container image discovery, distribution and change management, user and team collaboration, and workflow automation throughout the development pipeline.

## Table Usage Guide

The `dockerhub_repository` table provides insights into repositories within DockerHub. As a DevOps engineer, explore repository-specific details through this table, including repository name, description, star count, pull count, and last updated date. Utilize it to uncover information about repositories, such as those with high pull counts, the most starred repositories, and recently updated repositories.

## Examples

### Basic info
Explore the popularity and privacy settings of repositories on DockerHub. This query can help you identify popular repositories with high pull and star counts, and understand the balance between public and private repositories.

```sql+postgres
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository;
```

```sql+sqlite
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository;
```

### List private repositories
Explore which DockerHub repositories are set as private to gain insights into the level of exposure and access. This is useful for auditing and understanding the privacy settings of your repositories.

```sql+postgres
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  is_private;
```

```sql+sqlite
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  is_private = 1;
```

### List repositories with no pulls or downloads
Explore which Docker repositories have not been pulled or downloaded. This can help identify unused or less-popular repositories, enabling you to better manage your resources and focus on active repositories.

```sql+postgres
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  pull_count is null;
```

```sql+sqlite
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  pull_count is null;
```

### List repositories that have not received any stars or likes
Explore which repositories have not gained any popularity or recognition, helping you identify potential areas for improvement or promotion. This query can be useful in understanding which of your repositories may need more attention or enhancement to increase their visibility and user engagement.

```sql+postgres
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  star_count is null;
```

```sql+sqlite
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  star_count is null;
```

### List repositories which have not been updated in the last 7 days
Determine the areas in which repositories have remained inactive over the past week. This query is useful for identifying potentially outdated or unused repositories, aiding in the efficient management of resources.

```sql+postgres
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  last_updated > now() - interval '7' day;
```

```sql+sqlite
select
  name,
  pull_count,
  star_count,
  is_private,
  last_updated
from
  dockerhub_repository
where
  last_updated > datetime('now', '-7 day');
```