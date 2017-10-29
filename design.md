Zeal Registry
===========

Registry API that zeal cli will use.

## Metadata

Provide API to grab metadata of package. Zeal will always get
metadata separately than package files.

## Labels/Usage?

Feature to store metadata on where artifact is used/installed?

Also labels on artifact. Like test passed.

Should this be on this service on consume?

## Minimum Files Needed

Provide API to grab package contents based on OS and environment filter.
Server will send compressed file with contents that
satisfies the filter. This is useful for packages that have 
binaries for multiple OS and only one is needed for target OS
where package is installed.

# API

## Commands

### search `/api/v1/:repo/search/:keywords`

Search repo with provided keywords.

### publish `/api/v1/:repo/publish`

Publish a package. Capable to publish a different package per target OS

all/target OS
publish/
    metadata in body

### download `/api/v1/:repo/download/:package/:version/:platform`

Download package per target platform.

Support redirect to download file from different endpoint.

### metadata `/api/v1/:repo/metadata/:package/:version`

Get metadata of package. This includes configuration, dependencies, etc.

### versions `/api/v1/:repo/versions/:package`

Get available versions of package.

# Data Provider

Abstraction for zeal registry data. This allows support for different data Providers in future.

## Local

Store data on locally.

# Storage Provider

Abstraction for zeal file storage. This allows support for different storage Provider in future.

## Local

Store files locally.

# Auth Provider

Abstraction for auth provider. This allows support for different auth Provider in future.

## Local

Load auth config from JSON file.