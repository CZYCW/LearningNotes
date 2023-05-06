import boto3
from botocore.client import Config

# create an S3 client
s3 = boto3.client('s3', config=Config(signature_version='s3v4'))

# generate a presigned URL for a specific S3 object
bucket_name = 'luchen-storage'
object_key = 'models/model.zip'
url = s3.generate_presigned_url(
    'get_object',
    Params={
        'Bucket': bucket_name,
        'Key': object_key
    },
    ExpiresIn=604800 # URL expires in 1 hour
)

print('Presigned URL: {}'.format(url))