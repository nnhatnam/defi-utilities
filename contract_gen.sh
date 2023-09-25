echo 'Remove .build folder'
rm -rf contract_build 

echo 'Generating ERC20...'
./solc --optimize --bin ./contracts/token/ERC20/ERC20.sol -o contract_build
./solc --optimize --abi ./contracts/token/ERC20/ERC20.sol -o contract_build
./abigen --bin=./contract_build/ERC20.bin --abi=./contract_build/ERC20.abi --pkg=ERC20 --out=./standard/ERC20/ERC20.go
