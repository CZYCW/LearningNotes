import requests

# Generate the pre-signed URL using the code from the previous example
url = 'https://luchen-storage.s3.amazonaws.com/uploads/myfile.txt?AWSAccessKeyId=AKIAZUOMWBQ5GSFPF2WF&Signature=K96SRtUzQIQuAMuj31J346g466Q%3D&Expires=1683773538'

# Set the file path for the upload
file_path = './uploads/myfile.txt'

# Open the file and read the contents
with open(file_path, 'r') as file:
    file_data = file.read()

# Set the content type header based on the file extension
content_type = 'application/octet-stream'  # Default content type
if file_path.endswith('.txt'):
    content_type = 'text/plain'
elif file_path.endswith('.jpg'):
    content_type = 'image/jpeg'
# Add more content types as needed

# Set the headers for the upload
headers = {'Content-Type': content_type}

# Upload the file to the pre-signed URL
response = requests.put(url, data=file_data)

# Check if the upload was successful
if response.status_code == 200:
    print("Upload successful!")
else:
    print("Upload failed:", response.text)