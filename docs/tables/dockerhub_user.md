# Table: dockerhub_user

A Docker Hub user refers to an individual who has created an account on Docker Hub, which is a cloud-based service and repository for Docker images. Docker Hub users can utilize the platform to store, share, and discover Docker container images.

By creating a Docker Hub user account, individuals gain access to various features and functionalities. They can upload their own Docker images to their personal repositories, tag and version their images, and securely store them on Docker Hub. Users can also search for and pull public Docker images from other users or organizations, facilitating easy deployment and collaboration.

## Examples

### Basic info

```sql
select
  id,
  name,
  full_name,
  company,
  joined,
  location,
  private_repositories
from
  dockerhub_user;
```

### List users who do not have any private repositories

```sql
select
  id,
  name,
  full_name,
  company,
  joined,
  location,
  private_repositories
from
  dockerhub_user
where
  private_repositories is null;
```

### List users who are from the turbot company

```sql
select
  id,
  name,
  full_name,
  company,
  joined,
  location,
  private_repositories
from
  dockerhub_user
where
  company = 'turbot';
```

### List users who are joined in the last 30 days

```sql
select
  id,
  name,
  full_name,
  company,
  joined,
  location,
  private_repositories
from
  dockerhub_user
where
  joined >= now() - interval '30' day;
```

### List users who are not part of any teams

```sql
select
  id,
  name,
  full_name,
  company,
  joined,
  location,
  private_repositories
from
  dockerhub_user
where
  teams is null;
```