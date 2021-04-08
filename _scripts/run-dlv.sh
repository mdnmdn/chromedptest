killall __debug_bin dlv
#dlv debug --accept-multiclient --continue --headless --listen=:2345 --api-version=2 --log /app/ -- -test.v=true
dlv debug --accept-multiclient --continue --headless --api-version=2 --log /app/ --