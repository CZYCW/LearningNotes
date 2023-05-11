import boto3
from botocore.exceptions import ClientError
from datetime import datetime, timedelta

# Set up an S3 client
s3 = boto3.client('s3', region_name='us-west-2')

# Set the path and filename for the upload
key = 'uploads/myfile.txt'

# Set the expiration time for the pre-signed URL
expiration = datetime.now() + timedelta(days=7)

try:
    # Generate the pre-signed URL
    url = s3.generate_presigned_url(
        'put_object',
        Params={
            'Bucket': 'luchen-storage',
            'Key': key,
        },
        ExpiresIn=900,  # 900 seconds (15 minutes)
    )
    # Print the pre-signed URL
    print("Upload URL:", url)

    # The client can use this URL to upload the file to S3
    # Example using requests library:
    # response = requests.put(url, data=file_data, headers=headers)
except ClientError as e:
    print("Error creating pre-signed URL:", e)