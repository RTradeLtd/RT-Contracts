#! /bin/bash

echo "Please enter the path to your keystore file"
read -r KEYFILE
echo ""

echo "Please enter the password for your keystore file"
read -r KEYPASS
echo ""

echo "Please enter endpoint for ethereum blockchain"
read -r ENDPOINT
echo ""

echo "Please enter merged miner validator contract address"
read -r ADDRESS
echo ""

cat << EOF > config.json
{
 "endpoint": "$ENDPOINT",
 "key_file_path": "$KEYFILE",
 "key_file_pass": "$KEYPASS",
 "contract_address": "$ADDRESS"
}
EOF

echo "Setup complete"