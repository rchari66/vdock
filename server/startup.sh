#!/bin/bash

# Download vdockPlugin for adding custom snippets
git clone https://github.com/rchari66/vdock-plugin-cloud9.git ${C9_PATH}vdockPlugin
cp -r cp -r ${C9_PATH}vdockPlugin /root/.c9/plugins/

# start the cloud9
node /work/c9sdk/server.js -p 8282 -w ${C9_PATH} >> /dev/null &

# start main server
cd /server/src && /server/src/main
