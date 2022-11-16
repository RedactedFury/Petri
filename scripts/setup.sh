make all

rm -rf ~/.sxfury

mkdir ~/.sxfury

comdex init --chain-id redshift redshift-01
comdex keys add test --recover --keyring-backend test<<<"y
wage thunder live sense resemble foil apple course spin horse glass mansion midnight laundry acoustic rhythm loan scale talent push green direct brick please"
comdex add-genesis-account test 100000000000000stake --keyring-backend test
comdex gentx test 1000000000stake --chain-id test --keyring-backend test
comdex collect-gentxs
comdex start