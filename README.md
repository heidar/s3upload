== s3upload.go

s3upload.go is a simple command line tool to upload images to Amazon S3 in
parallel. I created this tool as a part of a concurrency course at university.
Someone at Heroku suggsted to my professor that this would be a good project
so I thought I might as well do it, since I wanted to learn how to use S3 too.

== Usage

There are a few command line options that can be passed to the program.

    -r region (default USWest2)
    -b bucket
    -d directory (default current directory)
    -p permission (default BucketOwnerFull)

This is the full list of regions

    APNortheast
    APSoutheast
    APSoutheast2
    EUWest
    SAEast
    USEast
    USWest
    USWest2

This is the full list of permissions:

    Private
    PublicRead
    PublicReadWrite
    AuthenticatedRead
    BucketOwnerRead
    BucketOwnerFull

== Other

I kind of wanted to do it all from scratch and just do a HTTP PUT with HMAC
authentication but I found this really nice Go library for AWS and S3 from
Ubuntu. I decided to make my life easier and use that instead.

== TODO

 - Change from uploading everything at once to a fixed amount of workers
 - Set the correct MIME type (not just JPG)