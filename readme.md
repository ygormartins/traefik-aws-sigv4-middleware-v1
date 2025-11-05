# Traefik SigV4 Middleware

This is a Traefik plugin that lets you sign your requests sent to an aws-like API.
The plugin can be used to serve static sites from an s3-compatible provider.

## Configuration

| Option       | Required | Description       |
| ------------ | -------- | ----------------- |
| accessKey    | X        | aws Access Key    |
| secretKey    | X        | aws Secret Key    |
| sessionToken |          | aws Session Token |
| service      | X        | aws Service       |
| endpoint     | X        | aws Endpoint      |
| region       | X        | aws Region        |

## Example config

**Static config**

```yaml
# traefik.yml
experimental:
  plugins:
    sigv4middleware:
      moduleName: "github.com/ygormartins/traefikAwsSigv4Middlewarev1"
      # Populate this with the latest release tag.
      version: vX.Y.Z
```

**Dynamic config**

```yaml
http:
  middlewares:
    sigv4:
      plugin:
        sigv4middleware:
          accessKey: ROOTNAME
          secretKey: CHANGEME123
          service: s3
          endpoint: minio.localhost
          region: us-east-1
  routers:
    minio:
      rule: host(`minio.localhost`)
      service: minio@docker
      entryPoints:
        - web
      middlewares:
        - sigv4
```
