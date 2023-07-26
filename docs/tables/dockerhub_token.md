# Table: dockerhub_token

A personal access token (PAT) refers to an authentication token that grants access to Docker services or APIs using personal credentials. It is used for authenticating and authorizing actions performed within the Docker ecosystem.

## Examples

### Basic info

```sql
select
  uuid,
  is_active,
  generated_by,
  creator_ua,
  creator_ip,
  created_at,
  client_id
from
  dockerhub_token;
```

### List inactive tokens

```sql
select
  uuid,
  is_active,
  generated_by,
  creator_ua,
  creator_ip,
  created_at,
  client_id
from
  dockerhub_token
where
  not is_active;
```

### List tokens which have never been used

```sql
select
  uuid,
  is_active,
  generated_by,
  creator_ua,
  creator_ip,
  created_at,
  client_id
from
  dockerhub_token
where
  last_used is null;
```

### List manually generated tokens

```sql
select
  uuid,
  is_active,
  generated_by,
  creator_ua,
  creator_ip,
  created_at,
  client_id
from
  dockerhub_token
where
  generated_by = 'manual';
```

### List tokens which are older than 90 days

```sql
select
  uuid,
  is_active,
  generated_by,
  creator_ua,
  creator_ip,
  created_at,
  client_id
from
  dockerhub_token
where
  created_at < now() - interval '90' day;
```