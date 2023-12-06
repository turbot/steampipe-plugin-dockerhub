---
title: "Steampipe Table: dockerhub_tag - Query DockerHub Tags using SQL"
description: "Allows users to query DockerHub Tags, providing detailed information about all the image tags within a DockerHub repository."
---

# Table: dockerhub_tag - Query DockerHub Tags using SQL

DockerHub is a cloud-based registry service that allows you to link to code repositories, build your images and test them, stores manually pushed images, and links to Docker Cloud so you can deploy images to your hosts. It provides a centralized resource for container image discovery, distribution and change management, user and team collaboration, and workflow automation throughout the development pipeline. 

## Table Usage Guide

The `dockerhub_tag` table provides insights into the tags within DockerHub repositories. As a DevOps engineer, you can explore tag-specific details through this table, including the associated DockerHub repository, the tag name, and its manifest. Utilize this table to manage and monitor your DockerHub repositories, ensuring that all tags are up-to-date and follow your organization's naming conventions.

## Examples

### Basic info
Explore the status and usage of Docker images by identifying when they were last updated, pushed, or pulled, and their size. This allows for efficient management and tracking of Docker images in use.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that are from a specific repository, allowing you to analyze the status, last update, and size of these segments. This can be useful for managing and optimizing your repository's resources.

```sql+postgres
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

```sql+sqlite
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

### List tags with no pulls or downloads
Discover the segments that contain tags with no pulls or downloads in order to identify potentially unused or unpopular resources. This can be useful in optimizing resource allocation and improving overall system efficiency.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that contain tags which are not currently active. This provides valuable insights to assess and manage inactive components within your system.

```sql+postgres
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

```sql+sqlite
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

### List tags which are last updated by a particular user
Explore tags updated by someone other than a specific user to gain insights into the status, size, and last activities. This can help in tracking user contributions and managing resources effectively.

```sql+postgres
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

```sql+sqlite
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
Explore the various attributes of images linked to a specific tag in a Docker Hub repository. This can help in understanding the characteristics of these images, such as their architecture, operating system, size, and status, as well as when they were last updated or accessed.

```sql+postgres
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
  name like 'souravthe/test:latest';
```

```sql+sqlite
select
  name,
  json_extract(i.value, '$.Architecture') as architecture,
  json_extract(i.value, '$.Digest') as digest,
  json_extract(i.value, '$.LastPulled') as last_pulled,
  json_extract(i.value, '$.LastPushed') as last_pushed,
  json_extract(i.value, '$.Os') as os,
  json_extract(i.value, '$.Size') as size,
  json_extract(i.value, '$.Status') as status,
  json_extract(i.value, '$.Variant') as variant
from
  dockerhub_tag,
  json_each(images) as i
where
  name like 'souravthe/test:latest';
```