# Contribution guide

Everyone is more than welcome to contribute to this project! That's what open source is all about!

In the following, we try to provide some help and guidance on how to participate, contribute and develop on the project.

## Raising an issue or feature request

The simplest way to contribute to the project is to use it!

Whenever you encounter any issues or missing features while using the project, please create a [GitHub Issue](https://github.com/cluetec/lifeboat/issues) in this project and describe what you found or what you need.

## Project Setup & Development

If you are interested in getting your hands dirty, the following subsections provide information and instructions for the setup of your local development environment.

TODO


### Git workflow

The development is orientated on GitHub flow. Therefore every source code contribution needs to be provided through a [GitHub Pull-Request](https://github.com/cluetec/lifeboat/pulls) against the `main` branch.

#### Releases & Versioning

Releases will be made by tagging a specific commit on the `main` branch. For the versioning we are using the [Semantic Versioning Schema](https://semver.org/):

> Given a version number MAJOR.MINOR.PATCH, increment the:
>
> 1. MAJOR version when you make incompatible API changes
> 2. MINOR version when you add functionality in a backwards compatible manner
> 3. PATCH version when you make backwards compatible bug fixes

#### Patching of older release

Sometimes it's happening that security issues appear in older releases. Regarding the fact that the community behind this project is not very large, we are not able to provide patches for each release we have ever published. Therefore, _**we will only maintain the latest minor version with security patches and bug fixes!**_
