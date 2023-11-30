---
title: "Steampipe Table: dockerhub_user - Query Docker Hub Users using SQL"
description: "Allows users to query Docker Hub Users, specifically their details, providing insights into user profiles and their associated metadata."
---

# Table: dockerhub_user - Query Docker Hub Users using SQL

Docker Hub is a cloud-based registry service that allows you to link to code repositories, build your images and test them, stores manually pushed images, and links to Docker Cloud so you can deploy images to your hosts. It provides a centralized resource for container image discovery, distribution and change management, user and team collaboration, and workflow automation throughout the development pipeline.

## Table Usage Guide

The `dockerhub_user` table provides insights into user profiles within Docker Hub. As a DevOps engineer, explore user-specific details through this table, including their Docker ID, type of account, company details, and associated metadata. Utilize it to uncover information about users, such as their profile details, company affiliations, and the verification of user profiles.

## Examples

### Basic info
Gain insights into the profile details of DockerHub users, such as their company affiliation and location. This query is useful for understanding user demographics and their private repository usage.

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
Discover the segments of DockerHub users who have not created any private repositories. This is useful for understanding user behavior and identifying potential opportunities for promoting the use of private repositories.

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

### List users who are from a particular company
Explore which users are associated with a specific company. This can be particularly useful for gaining insights into the distribution of users across different companies.

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

### List users who have joined in the last 30 days
Identify the recent additions to your Dockerhub user base by pinpointing those who have joined within the past month. This allows you to keep track of your growing community and understand the pace of your user acquisition.

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
Discover the segments of users who have not joined any teams, providing insights into potential areas for team collaboration and resource allocation. This can be useful for understanding user engagement and optimizing team-based features.

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