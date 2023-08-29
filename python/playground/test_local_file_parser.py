_LOCAL_PREFIX = "file://"
uri = "file:///data/1/model/64b7a0a2ad6077101ecb1ab2/serve"
local_path = uri.replace(_LOCAL_PREFIX, "", 1)
print(local_path)