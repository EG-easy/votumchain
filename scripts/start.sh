# Initialize configuration files and genesis file
votumd init --chain-id testchain

# Copy the `Address` output here and save it for later use
votumcli keys add jack

# Copy the `Address` output here and save it for later use
votumcli keys add alice

# Add both accounts, with coins to the genesis file
votumd add-genesis-account $(votumcli keys show jack -a) 1000votum
votumd add-genesis-account $(votumcli keys show alice -a) 1000votum

# Configure your CLI to eliminate need for chain-id flag
votumcli config chain-id testchain
votumcli config output json
votumcli config indent true
votumcli config trust-node true
