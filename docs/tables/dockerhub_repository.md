# Table: dockerhub_repository

A repository refers to a collection of Docker images with similar characteristics or versions. It is a centralized storage location where Docker images are stored and organized.

## Examples

### Basic info

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

### List private repositories

```sql
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

### List repositories with no pulls or downloads

```sql
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

```sql
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

```sql
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
