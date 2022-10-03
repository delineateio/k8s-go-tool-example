[![PRs Welcome][pr-welcome-shield]][pr-welcome-url]
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <img alt="delineate.io" src="https://github.com/delineateio/.github/blob/master/assets/logo.png?raw=true" height="75" />
  <h2 align="center">delineate.io</h2>
  <p align="center">portray or describe (something) precisely.</p>

  <h3 align="center">k8s Go Tooling Example</h3>

  <p align="center">
    Demonstrates a combination of k8s related tools for efficient build, test and deployment of Go microservices.
    <br />
    <a href="https://github.com/delineateio/k8s-go-tool-example"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/delineateio/k8s-go-tool-example">View Demo</a>
    ·
    <a href="https://github.com/delineateio/k8s-go-tool-example/issues">Report Bug</a>
    ·
    <a href="https://github.com/delineateio/k8s-go-tool-example/issues">Request Feature</a>
  </p>
</p>

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [About The Project](#about-the-project)
- [What's been built](#whats-been-built)
- [Built With](#built-with)
- [Local Setup](#local-setup)
  - [Dev Container](#dev-container)
  - [Kubernetes](#kubernetes)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

<!-- ABOUT THE PROJECT -->
## About The Project

This repo demonstrates a combination of k8s related tools that when used in combination provide an efficient build, test and deployment pipeline for developing Go microservices.

* Using [ko](https://github.com/ko-build/ko) to create distroless images from [go](https://go.dev) source code
* Using [kustomize](https://kustomize.io/) to dyanmically transform k8s manifests before deployment
* Using [skaffold](https://github.com/GoogleContainerTools/skaffold) to manage a build, test and deploy opinionated pipeline

In addition to this DevContainer has been used to package the local development environment.

## What's been built

Two "unique" services have been containerised that run in the background with no API, one is English and one is Spainish 🇬🇧🇪🇸, and print `hello` to the stdout.

![logs](assets/logs.png "logs")

In addition to these two services, they have been packaged in a pod with a [k8s echo server](https://console.cloud.google.com/gcr/images/kubernetes-e2e-test-images/global/echoserver) container.  When running, the echo server can be accessed using `http :8080` either from the host of from within the dev container.

## Built With

| Syntax | Description |
| --- | ----------- |
| ![pre-commit](https://img.shields.io/badge/precommit-%235835CC.svg?style=for-the-badge&logo=precommit&logoColor=white) | Pre-commit `git` hooks that perform checks before pushes|
| ![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white) | Source control management platform  |
| ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) | Writing of the example containerised services |
| ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) | Used to providr the [devcontainer](https://containers.dev/implementors/json_reference/) |
| ![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white) | Provide the hosting of the containerised applications |

<!-- LOCAL SETUP -->
## Local Setup

### Dev Container

The local development environment is provided using [devcontainer](https://containers.dev/implementors/json_reference/) which is natively supported in VSCode as described [here](https://code.visualstudio.com/docs/remote/containers).  This has been used to show how this approach signficantly reduces local requirements.

### Kubernetes

This examples repo relies on having access to a running local or remote k8s cluster, correctly setup and accessible with [kubectl](https://kubernetes.io/docs/tasks/tools/) using `kubectl cluster-info`.

As part of t`he development container provisioning the [kube-config](.devcontainer/.kube-config.sh) script replicates the hosts `$HOME/.kube/config` within the container.

This particular implementation has only been tested using [minikube](https://minikube.sigs.k8s.io/docs/) configured to use the [VMWare driver](https://minikube.sigs.k8s.io/docs/drivers/vmware/).

<!-- USAGE EXAMPLES -->
## Usage

Inside the container using the VSCode integrated terminal the following commands can be used...

```shell
# builds, deploys (1 pod) and watches for changes
task single:dev

#  alternatively, this builds and deploys once
task single:run

# cleans up the cluster
task cleanup

# builds, deploys a scaled example (3 pods)
task scaled:dev
```

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/delineateio/k8s-go-tool-example/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

If you would like to contribute to any delineate.io OSS projects please read:

* [Code of Conduct](https://github.com/delineateio/.github/blob/master/CODE_OF_CONDUCT.md)
* [Contributing Guidelines](https://github.com/delineateio/.github/blob/master/CONTRIBUTING.md)

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [Markdown Badges](https://github.com/Ileriayo/markdown-badges)
* [DocToc](https://github.com/thlorenz/doctoc)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[pr-welcome-shield]: https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=for-the-badge&logo=github
[pr-welcome-url]: https://github.com/delineateio/k8s-go-tool-example/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue
[contributors-shield]: https://img.shields.io/github/contributors/delineateio/k8s-go-tool-example.svg?style=for-the-badge&logo=github
[contributors-url]: https://github.com/delineateio/k8s-go-tool-example/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/delineateio/k8s-go-tool-example.svg?style=for-the-badge&logo=github
[forks-url]: https://github.com/delineateio/k8s-go-tool-example/network/members
[stars-shield]: https://img.shields.io/github/stars/delineateio/k8s-go-tool-example.svg?style=for-the-badge&logo=github
[stars-url]: https://github.com/delineateio/k8s-go-tool-example/stargazers
[issues-shield]: https://img.shields.io/github/issues/delineateio/k8s-go-tool-example.svg?style=for-the-badge&logo=github
[issues-url]: https://github.com/delineateio/k8s-go-tool-example/issues
[license-shield]: https://img.shields.io/github/license/delineateio/k8s-go-tool-example.svg?style=for-the-badge&logo=github
[license-url]: https://github.com/delineateio/k8s-go-tool-example/blob/master/LICENSE
