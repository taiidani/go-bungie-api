# Go Bungie API

This repository maintains a Go implementation of the Bungie API documented at https://github.com/Bungie-net/api. It is built based upon the OpenAPI 3.0 specification and adapted to meet the needs of consuming clients.

> [!WARNING]
> This library is still a work in progress and is subject to substantial changes while the Bungie API parsing logic is built. Use at your own risk!

## Contributing

Contributions are welcome! The repository uses [mise](https://mise.jdx.dev) for managing dependencies, tasks, and environment variable expectations.

### Testing

This library may be tested locally. Generate the schema from the "generate" task:

```sh
mise run generate
```

Run the tests via the "test" task:

```sh
mise run test
```
