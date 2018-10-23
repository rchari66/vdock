# start the cloud9
node /work/c9sdk/server.js -p 8282 -w ${C9_PATH} >> /dev/null &

# start main server
cd /server/src && /server/src/main
