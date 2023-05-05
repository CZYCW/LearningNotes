import os
import zipfile

def zipdir(path, ziph):
    # Walk through all files in the given path
    for root, dirs, files in os.walk(path):
        for file in files:
            # Create the full path of the file
            filepath = os.path.join(root, file)
            print(file, filepath)
            # Add the file to the zip archive
            ziph.write(filepath, os.path.relpath(filepath, path))

# The path of the folder to be compressed
folder_path = "/Users/ziyuanc/Desktop/notes/ray"

# The name of the zip file to be created
zip_name = "/Users/ziyuanc/Desktop/model.zip"

# Create a ZipFile object to write the zip file
zip_file = zipfile.ZipFile(zip_name, "w", zipfile.ZIP_DEFLATED)

# Compress all files in the folder
zipdir(folder_path, zip_file)

# Close the ZipFile object
zip_file.close()