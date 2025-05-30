# Ubuntu based os
FROM techtaste/dock-c9

LABEL maintainer rchari<techtaste.me>

# SERVER_ROOT_DIR
ENV SERVER_PATH /server
ENV HUGO_BIN ${SERVER_PATH}/scripts/hugo_bin
ENV BLOG_PATH=/ws/Blog-Source
ENV SITE_PATH=/ws/MySite
ENV ALL_POSTS=/ws/AllPosts

ENV C9_PATH /ws/
ENV C9_HOME /root

ENV LOG_FILE=/tmp/server.log
# Add hugo_bin to PATH
ENV PATH=${PATH}:${HUGO_BIN}

ADD server/ /server/
ADD bashrc.sh /tmp/bashrc.sh

RUN mkdir /ws && apt-get update -y \
    && apt-get install hugo -y \
    && apt-get install iproute2 -y \
    && DEBIAN_FRONTEND=noninteractive apt-get install expect -y \
    && cat /tmp/bashrc.sh >> /root/.bashrc && rm /tmp/bashrc.sh \
    && rm /work/proxy* && rm /work/startup.sh \
    && chmod +x $SERVER_PATH/scripts/hugo_bin/*

ADD user.settings ${C9_HOME}/.c9/user.settings
RUN sed -i -e 's/Agent.prototype.connectionTimeout = 10000;/Agent.prototype.connectionTimeout = 50000;/g' /work/c9sdk/node_modules/smith/smith.js

EXPOSE 8288 8286

WORKDIR /server/

CMD ["/bin/bash", "/server/startup.sh"]
