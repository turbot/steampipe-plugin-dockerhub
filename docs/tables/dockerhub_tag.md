# Table: dockerhub_tag

A tag is a label or identifier attached to a specific version, variant, or configuration of a Docker image within a repository. It helps to differentiate and manage different versions or variations of the same image.

## Examples

### Basic info

```sql
select
  name,
  status,
  last_updater_user_name,
  last_pushed,
  last_pulled,
  full_size
from
  dockerhub_tag;
```

### List tags which are from a particular repository

```sql
select
  name,
  status,
  last_updater_user_name,
  last_pushed,
  last_pulled,
  full_size
from
  dockerhub_tag
where
  name like 'souravthe/test%';
```

### List tags which are never pulled

```sql
select
  name,
  status,
  last_updater_user_name,
  last_pushed,
  last_pulled,
  full_size
from
  dockerhub_tag
where
  last_pulled is null;
```

### List tags which are not active

```sql
select
  name,
  status,
  last_updater_user_name,
  last_pushed,
  last_pulled,
  full_size
from
  dockerhub_tag
where
  status <> 'active';
```

### List tags which are last updated by john

```sql
select
  name,
  status,
  last_updater_user_name,
  last_pushed,
  last_pulled,
  full_size
from
  dockerhub_tag
where
  last_updater_user_name <> 'john';
```

### List of images associated with a particular tag

```sql
select
  name,
  i ->> 'Architecture' as architecture,
  i ->> 'Digest' as digest,
  i ->> 'LastPulled' as last_pulled,
  i ->> 'LastPushed' as last_pushed,
  i ->> 'Os' as os,
  i ->> 'Size' as size,
  i ->> 'Status' as status,
  i ->> 'Variant' as variant
from
  dockerhub_tag,
  jsonb_array_elements(images) as i
where
  name like 'souravthe/test:latest';;
```
