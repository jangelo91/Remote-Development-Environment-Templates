# Java Remote Development Template

## Overview

The Java Base Template is a set of files that allows the user to develop within a fully isolated remote Java environment. This template utilizes [Visual Studio Code’s](<https://code.visualstudio.com/learn>) Remote – Containers [extension](<https://marketplace.visualstudio.com/VSCode>) in addition to [Docker containers](<https://www.docker.com/resources/what-container>). This template is open to anyone, from students who want to learn the Java programming language to development teams looking to streamline their development environment setup. Docker builds the entire environment within minutes, so there is no need to manually install and configure dependencies within the containerized workspace. Simply deploy a container for each user and, when necessary, modify the config files and update all workspaces simultaneously (advanced). Updating the users' development environments is a simple process that involes modifying the dependencies specfied by the config files and rebuilding the container. This folder contains configuration files to instruct Docker on how to build the Docker container and VS Code what to include (e.g. VS Code extensions, settings, arguments, etc.). In addition, there is a test file for demonstration.

### Devcontainer Properties

VS Code needs to be instructed on how to build the container. This can include the base image variant, the container path for the Java runtime, the name of the workspace folder, any VS Code extension, etc. The properties can be modified in the devcontainer.json file (./.devcontainer/devcontainer.json). A full list of container properties can be found on <https://code.visualstudio.com/docs/remote/devcontainerjson-reference>.

### Dockerfile Properties

The Dockerfile is a script that Docker uses to build the container. It can be modified to run any Linux commands during the image build process. This can include the installation of dependnecies, tools, platforms, etc. For more information, go to <https://docs.docker.com/>.

## Requirements

- [Windows 10 Version 2004 or higher](<https://support.microsoft.com/en-us/windows/get-the-windows-10-october-2020-update-7d20e88c-0568-483a-37bc-c3885390d212>)
- [Windows Subsystem for Linux (WSL2)](<https://docs.microsoft.com/en-us/windows/wsl/install-win10>)
- [Visual Studio Code](<https://code.visualstudio.com/>)
- [Docker Desktop/Engine](<https://www.docker.com/products/docker-desktop>)
- [Visual Studio Code Docker](<https://code.visualstudio.com/docs/containers/overview>)
- [Visual Studio Code Remote Development](<https://code.visualstudio.com/docs/remote/remote-overview>)

See <https://code.visualstudio.com/docs/remote/containers-tutorial> for a complete guide on how to setup remote containers (recommended for beginners).

## How to Use the Java Template

Copy the .devcontainer folder into your project's root directory. VS Code may prompt you to reopen the folder in a container. If you are satisfied with the container configuration, select "Reopen in Container."

## Getting Started with Java Remote Development Containers

1. Clone this repository from GitHub

2. Reopen your local copy of this folder (/Java/Base) in VS Code

3. Open the [command pallette](<https://code.visualstudio.com/docs/getstarted/userinterface#:~:text=The%20most%20important%20key%20combination,provides%20access%20to%20many%20commands.>) (press F1 or Ctrl+Shift+P)

4. Enter "Remote-Containers: Open Folder in Container"

5. Select this folder. Docker should start building the container, and there will be a progress bar to the bottom right of the window

6. Once the container is built, open Test.java and run the code (press F5, hit "run" right above the "main" method (only applicable if the Java extension is installed), or navigate to Run > Start Debugging)