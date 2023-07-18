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

### List the private repositories

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

### List the repositories which are never pulled

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

### List the repositories which are updated 90 days ago

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
  last_updated < now() - interval '90' day;
```
