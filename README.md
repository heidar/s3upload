# s3upload.go

s3upload.go is a simple command line tool to upload files to Amazon S3 in
parallel. I created this tool as a part of a concurrency course at university.
Someone at Heroku suggsted to my professor that this would be a good project
so I thought I might as well do it, since I wanted to learn how to use S3 too.

I kind of wanted to do it all from scratch and just do a HTTP PUT with HMAC
authentication but I found this really nice Go library for AWS and S3 from
Ubuntu. I decided to make my life easier and use that instead.

## Usage

### Command Line Options

    -r region (default USWest2)
    -b bucket
    -d directory (default current directory)
    -p permission (default BucketOwnerFull)

### Regions

    APNortheast
    APSoutheast
    APSoutheast2
    EUWest
    SAEast
    USEast
    USWest
    USWest2

### Permissions

    Private
    PublicRead
    PublicReadWrite
    AuthenticatedRead
    BucketOwnerRead
    BucketOwnerFull

### Authentication

Authentication is done through environment variables. Make sure the following
environment variables are set:

    AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY

If your environment variables do not work (like mine) with Go for some reason
then run it with the variables prepended:

    AWS_SECRET_ACCESS_KEY=xyz AWS_SECRET_ACCESS_KEY=abc go run s3upload.go ...

## TODO

 - Change from uploading everything at once to a fixed amount of workers
