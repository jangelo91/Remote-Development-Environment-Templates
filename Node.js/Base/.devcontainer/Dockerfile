# Made by Julius Angelo

# Specifies what version of NodeJS to use
ARG VERSION

# Full list of tags can be found at https://github.com/docker-library/docs/blob/master/openjdk/README.md#supported-tags-and-respective-dockerfile-links
FROM node:${VERSION}

#WORKDIR /home/node/workspace
#COPY /src .

RUN npm install -g npm

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment if you want to install an additional version of node using nvm
# ARG EXTRA_NODE_VERSION=10
# RUN sudo node -c "source /usr/local/share/nvm/nvm.sh && nvm install ${EXTRA_NODE_VERSION}"

# [Optional] Uncomment if you want to install more global node modules
# RUN sudo -u node pm install -g <your-package-list-here>

RUN npm install --save-dev webpack \
    && npm install --save-dev broccoli \
    && npm install --save-dev clinic \
    && npm install --save-dev pm2 \
    && npm install --save-dev mocha \
    && npm install --save-dev chai \
    && npm install --save-dev  sinon