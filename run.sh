DATA_DIR=erigon_data

build/bin/erigon \
  --no-downloader=true \
  --snapshots=false \
  --chain='bor-mainnet' \
  --datadir=$DATA_DIR \
  --log.console.verbosity=4 \
  --http \
  --http.addr '127.0.0.1' \
  --http.vhosts '*' \
  --http.corsdomain '*' \
  --http.port 8545 \
  --http.api 'eth,net,web3,txpool,bor' \
  --txpool.nolocals \
  --txpool.pricelimit '30000000000' \
  --txpool.lifetime '1h30m0s' \
  --maxpeers 250 \
  --metrics \
  --bootnodes='enode://0a754459032d34df969f2c62258a31eb4a0ccc4fd5f3e3e071ab2e7163612feb63186c30b7664e8d2ca393248b472eb453425b4927241949b26fab35a21c5027@69.4.239.34:31410,enode://109be30a5dfcc16b3bdb3916c039464e49bd1ac81b5d04df19cb52dd1bf4c84635216fde36f411ade92373c41797941f91ac34fa3c9b99c2587ac22c6b18ed07@138.2.237.133:30304,enode://efd67dd4d7a4ef358b88c4815da4bca325d6aad046927adfd98e42e4c6bf66a0424c10e97419bcf4bedadf6819a4361fc7f870b7fa25e3c8a544030feba5e1cd@109.235.247.123:30303,enode://85796d187605868c0d956bef1decd27a4b1da0b814eb69ee2250bfcdcb6087631e12a245d868bc22e34d94edb8ecd40847f01143d7cd0dbef839e2a158b4d21c@93.115.25.190:30304,enode://0a754459032d34df969f2c62258a31eb4a0ccc4fd5f3e3e071ab2e7163612feb63186c30b7664e8d2ca393248b472eb453425b4927241949b26fab35a21c5027@69.4.239.34:31410,enode://109be30a5dfcc16b3bdb3916c039464e49bd1ac81b5d04df19cb52dd1bf4c84635216fde36f411ade92373c41797941f91ac34fa3c9b99c2587ac22c6b18ed07@138.2.237.133:30304,enode://efd67dd4d7a4ef358b88c4815da4bca325d6aad046927adfd98e42e4c6bf66a0424c10e97419bcf4bedadf6819a4361fc7f870b7fa25e3c8a544030feba5e1cd@109.235.247.123:30303,enode://85796d187605868c0d956bef1decd27a4b1da0b814eb69ee2250bfcdcb6087631e12a245d868bc22e34d94edb8ecd40847f01143d7cd0dbef839e2a158b4d21c@93.115.25.190:30304' \
  --txpool.accountslots=500 \
  --txpool.globalslots=100000 \
  --txpool.globalbasefeeslots=100000 \
  --txpool.accountqueue=500 \
  --txpool.globalqueue=100000 \
  --rpc.gascap=500000000 \
  --rpc.txfeecap=0 \
  --v5disc=true \
  --maxpeers=500
  # --p2p.protocol value [ --p2p.protocol value ]            Version of eth p2p protocol (default: 68, 67)
