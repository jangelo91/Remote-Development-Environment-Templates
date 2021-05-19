# Remote Development Environment Templates

## Overview

This project is a framework for remote development templates, which are a set of configuration files with instructions for automatically setting up development environments. This includes details like programming language, developer utilities, and software packages. The automation process uses [Docker containers](<https://www.docker.com/resources/what-container>), a platform for packaging a software application into a “container." For the purposes of this framework, the applications created can be thought of as containerized development environments. The [Visual Studio Code](<https://code.visualstudio.com/learn>) editor provides additional functionality with the Remote – Containers [extension](<https://marketplace.visualstudio.com/VSCode>).

This framework consists of several templates for different development environments, which can be easily modified, standardized, and automated by a course professor or team manager to set up an environment in minutes.  

This framework is open from students who want to learn the a programming language to development teams looking to streamline their development environment setup. Docker builds the entire environment within minutes, so there is no need to manually install and configure dependencies within the containerized workspace.

### Devcontainer Properties

VS Code needs to be instructed on how to build the container. This can include the base image variant, the container path for the Java runtime, the name of the workspace folder, any VS Code extension, etc. The properties can be modified in the devcontainer.json file (./.devcontainer/devcontainer.json). A full list of container properties can be found on <https://code.visualstudio.com/docs/remote/devcontainerjson-reference>.

### Dockerfile Properties

The Dockerfile is a script that Docker uses to build the container. It can be modified to run any Linux commands during the image build process. This can include the installation of dependencies, tools, platforms, etc. For more information, go to <https://docs.docker.com/>.

## Requirements

- [Windows 10 Version 2004 or higher](<https://support.microsoft.com/en-us/windows/get-the-windows-10-october-2020-update-7d20e88c-0568-483a-37bc-c3885390d212>)
- [Windows Subsystem for Linux (WSL2)](<https://docs.microsoft.com/en-us/windows/wsl/install-win10>)
- [Visual Studio Code](<https://code.visualstudio.com/>)
- [Docker Desktop/Engine](<https://www.docker.com/products/docker-desktop>)
- [Visual Studio Code Docker](<https://code.visualstudio.com/docs/containers/overview>)
- [Visual Studio Code Remote Development](<https://code.visualstudio.com/docs/remote/remote-overview>)

See <https://code.visualstudio.com/docs/remote/containers-tutorial> for a complete guide on how to setup remote containers (recommended for beginners).
