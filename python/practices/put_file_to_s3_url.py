import requests

# Generate the pre-signed URL using the code from the previous example
url = 'https://luchen-storage.s3.us-west-2.amazonaws.com/1/dataset/28/test/b.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAZUOMWBQ5KEYTB27K%2F20230511%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230511T055852Z&X-Amz-Expires=7200&X-Amz-SignedHeaders=host&X-Amz-Signature=f89e1c47d810776658105bfb4ed497668e5d495e9fb277ddfcfa2db96335fded'

# Set the file path for the upload
file_path = './uploads/b.txt'

# Open the file and read the contents
with open(file_path, 'r') as file:
    file_data = file.read()

response = requests.put(url, data=file_data)

# Check if the upload was successful
if response.status_code == 200:
    print("Upload successful!")
else:
    print("Upload failed:", response.text)
