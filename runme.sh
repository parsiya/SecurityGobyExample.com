echo "Starting build"
eval '.\/tools\/build'
echo "Starting upload"
eval 'python s3cmd sync --acl-public --delete-removed -MP --no-preserve --rr public\/ s3:\/\/securitygobyexample.com'

