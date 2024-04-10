# CloudQuery SNMP Source Plugin

[![test](https://github.com/francescofuggitti/cq-source-snmp/actions/workflows/test.yaml/badge.svg)](https://github.com/francescofuggitti/cq-source-snmp/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/francescofuggitti/cq-source-snmp)](https://goreportcard.com/report/github.com/francescofuggitti/cq-source-snmp)

A local file source plugin for CloudQuery that loads data from a file (in JSON, YAML, CSV or XLSX format) to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 <!-- - [Supported Tables](docs/tables/README.md) -->

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub  

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
