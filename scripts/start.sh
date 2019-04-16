# Initialize configuration files and genesis file
votumchaind init --chain-id testchain

# Copy the `Address` output here and save it for later use
votumchaincli keys add jack

# Copy the `Address` output here and save it for later use
votumchaincli keys add alice

# Add both accounts, with coins to the genesis file
votumchaind add-genesis-account $(votumchaincli keys show jack -a) 1000mycoin,1000jackcoin
votumchaind add-genesis-account $(votumchaincli keys show alice -a) 1000mycoin,1000alicecoin

# Configure your CLI to eliminate need for chain-id flag
votumchaincli config chain-id testchain
votumchaincli config output json
votumchaincli config indent true
votumchaincli config trust-node true
