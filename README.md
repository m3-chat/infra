# m3-chat infra CLI

[![Build](https://img.shields.io/github/actions/workflow/status/m3-chat/infra/release.yml?label=Build&logo=githubactions&style=flat-square)](https://github.com/m3-chat/infra/actions)
[![License: AGPL v3](https://img.shields.io/github/license/m3-chat/infra?style=flat-square)](./LICENSE)
[![Release](https://img.shields.io/github/v/release/m3-chat/infra?style=flat-square&label=Latest%20Release)](https://github.com/m3-chat/infra/releases)

Infra is the deployment automation tool for [M3 Chat](https://github.com/m3-chat). It simplifies installing and deploying the `frontend` and `backend` repositories.

---

## Usage

```bash
infra deploy backend   # Clones & sets up the backend locally
infra deploy frontend  # Clones & logs frontend deploy steps
```

---

## Releases

Prebuilt binaries are automatically published on GitHub Releases when a new version tag is pushed (`v*.*.*`).

[Download the latest binary](https://github.com/m3-chat/infra/releases/latest)
