import requests

# Generate the pre-signed URL using the code from the previous example
url = 'https://luchen-storage.tos-cn-beijing.volces.com/1/job/234/output/checkpoint/train.py?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTOTZiZjY0MmQ0NzNkNGJiOGFhOWEwYWQ0MmJhNzZjNGY%2F20230714%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20230714T064750Z&X-Tos-Expires=3600&X-Tos-Signature=25e059e7dcd21ecc280d2e367e8a853eadb68dc334cf7e5d75f4d074b4bc1b02&X-Tos-SignedHeaders=host'

# Set the file path for the upload
file_path = '/Users/ziyuanc/Desktop/work/train.py'

# Open the file and read the contents
with open(file_path, 'r') as file:
    file_data = file.read()

response = requests.put(url, data=file_data)

# Check if the upload was successful
if response.status_code == 200:
    print("Upload successful!")
else:
    print("Upload failed:", response.text)
