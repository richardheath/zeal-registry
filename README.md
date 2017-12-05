# zeal-registry

Registry service that zeal CLI will interact with.

## Providers

### Auth Provider

Implementation for Auth providers. This provider
is used to authenticate users.

### Data Provider

Implementation for Data provider. This provider is used
to store package metadata.

### Storage Provider

Implementation for Storage provider. This provider is used
to save and download package files.

## Config

You can set what providers to use and it's 
settings on config file.

```json
{
    "auth": {
        "provider": "local",
        "settings": {
            "path": "./server/auth/config.json"
        }
    },
    "data": {
        "provider": "local",
        "settings": {
            "path": "./server/data"
        }
    },
    "storage": {
        "provider": "local",
        "settings": {
            "path": "./server/storage"
        }
    }
}
```

## Usage

```
./zeal-registry ./config.json
```