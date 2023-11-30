---
title: "Steampipe Table: dockerhub_token - Query DockerHub Tokens using SQL"
description: "Allows users to query DockerHub Tokens, specifically the details associated with each token, providing insights into token usage and management."
---

# Table: dockerhub_token - Query DockerHub Tokens using SQL

DockerHub Tokens are a feature of DockerHub, a cloud-based repository service where developers can manage, store, and distribute Docker images. Tokens allow secure access to DockerHub repositories, providing an additional layer of security for images and containers. They are used to authenticate Docker CLI and Docker API requests, replacing the need for using username and password.

## Table Usage Guide

The `dockerhub_token` table provides insights into DockerHub Tokens within DockerHub. As a DevOps engineer, explore token-specific details through this table, including token ID, status, and associated metadata. Utilize it to manage and monitor tokens, such as those currently active, their permissions, and the time of their creation.

## Examples

### Basic info
Explore which DockerHub tokens are active and when they were created. This can be useful for auditing purposes, to track user activity and ensure security compliance.

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
Discover the segments that are associated with inactive tokens in DockerHub. This can be beneficial in identifying potential security risks and maintaining optimal system performance.

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
Identify unused tokens within your DockerHub setup to assess potential security risks or inefficiencies. This helps in maintaining a clean, secure, and efficient environment by removing or updating unused tokens.

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
Explore which tokens have been manually generated. This is beneficial in identifying potential security risks or anomalies related to token generation.

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
Determine the areas in which DockerHub tokens have remained active for more than 90 days. This can be useful for identifying potential security risks associated with outdated tokens.

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