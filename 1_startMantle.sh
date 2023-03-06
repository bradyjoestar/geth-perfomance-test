echo "Importing private key"
rm -rf key.prv password /root/.ethereum
echo 6587ae678cf4fc9a33000cdbf9f35226b71dcc6a4684a31203241f9bcfd55d27 > key.prv
echo "pwd" > password
geth_mantle_linux account import --password ./password ./key.prv

echo "Initializing Geth node"
geth_mantle_linux --verbosity=4 "$@" init genesis.json

echo "Starting Geth node"
exec geth_mantle_linux \
  --verbosity=4 \
  --password ./password \
  --allow-insecure-unlock \
  --unlock 0x00000398232E2064F896018496b4b44b3D62751F \
  --mine \
  --miner.etherbase 0x00000398232E2064F896018496b4b44b3D62751F \
  "$@"
