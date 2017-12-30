echo "Starting build"
eval '.\/tools\/build'
echo "Starting upload"
eval 'python s3cmd sync --acl-public --delete-removed -MP --no-preserve --rr public\/ s3:\/\/securitygobyexample.com'
eval 'python s3cmd --acl-public --no-preserve --cf-invalidate --add-header="Expires: Sat, 20 Nov 2286 19:00:00 GMT" --mime-type="text/css" put public\/site.css s3:\/\/securitygobyexample.com\/site.css'



