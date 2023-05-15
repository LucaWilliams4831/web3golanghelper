import os

# generate bins and abis from contract
cmd = "solc0.6.6 --overwrite " + "--abi " + "contracts/PancakeMainnetContract.sol " + " -o contracts/build"
run_result = os.system(cmd)
if run_result != 0:
    print("Error generating abi")
    os.abort()

cmd = "solc0.6.6 --overwrite " + "--bin " + "contracts/PancakeMainnetContract.sol " + " -o contracts/build"
run_result = os.system(cmd)
if run_result != 0:
    print("Error generating bin")
    os.abort()

# List all files in a directory using os.listdir
basepath = 'contracts/build/'
for entry in os.listdir(basepath):
    if os.path.isfile(os.path.join(basepath, entry)):
        split_name = entry.split(".")
        filename = split_name[0]
        file_extension = split_name[1]

        # for run only once
        if file_extension == 'abi':
            # [abigen --bin=contracts/build/IPancakeRouter01.bin --abi=contracts/build/IPancakeRouter01.abi --pkg=IPancakeRouter01  --out=contracts/IPancakeRouter01.go]
            cmd = "abigen " + "--bin=contracts/build/" + filename + ".bin " + "--abi=contracts/build/" + filename + ".abi " + "--pkg=" + filename + " --out=contracts/" + filename + ".go"
            run_result = os.system(cmd)

            if run_result != 0:
                print("Error generating golang source code from " + filename)
                os.abort()

            os.mkdir("contracts/" + filename)
            os.rename("contracts/" + filename + ".go", "contracts/" + filename + "/" + filename + ".go")

