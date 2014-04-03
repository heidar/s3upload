package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

var (
	region     string
	bucketName string
	directory  string
	permission string
)

func init() {
	flag.StringVar(&region,     "r", "USWest2",         "region")
	flag.StringVar(&bucketName, "b", "",                "bucket name")
	flag.StringVar(&directory,  "d", ".",               "directory")
	flag.StringVar(&permission, "p", "BucketOwnerFull", "permission")
}

func upload(directory string, f os.FileInfo, bucket *s3.Bucket, permission s3.ACL, done chan bool) {
	log.Println("uploading " + f.Name())
	data, err := ioutil.ReadFile(path.Join(directory, f.Name()))
	if err != nil {
		panic(err.Error())
	}
	err = bucket.Put(f.Name(), data, "", permission)
	if err != nil {
		panic(err.Error())
	}
	log.Println("finished uploading " + f.Name())
	done <- true
}

func main() {
	flag.Parse()

	// map of all the aws regions
	awsRegions := map[string] aws.Region {
		"APNortheast"  : aws.APNortheast,
		"APSoutheast"  : aws.APSoutheast,
		"APSoutheast2" : aws.APSoutheast2,
		"EUWest"       : aws.EUWest,
		"SAEast"       : aws.SAEast,
		"USEast"       : aws.USEast,
		"USWest"       : aws.USWest,
		"USWest2"      : aws.USWest2,
	}

	// check region exists
	if _, ok := awsRegions[region]; !ok {
		log.Println("error: " + region + " is not a valid aws region")
		return
	}

	// check directory exists
	result, err := os.Stat(directory)
	if err != nil {
		panic(err.Error())
	} else {
		if !result.IsDir() {
			log.Println("error: " + directory + " is not a directory")
			return
		}
	}

	// map of all permission types
	permissions := map[string] s3.ACL {
		"Private"           : s3.Private, 
		"PublicRead"        : s3.PublicRead,
		"PublicReadWrite"   : s3.PublicReadWrite,
		"AuthenticatedRead" : s3.AuthenticatedRead,
		"BucketOwnerRead"   : s3.BucketOwnerRead,
		"BucketOwnerFull"   : s3.BucketOwnerFull,
	}

	// check permission type exists
	if _, ok := permissions[permission]; !ok {
		log.Println("error: " + permission + " is not a valid s3 permission type")
		return
	}

	// AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY env vars are used
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	// authenticate with s3 and acesss the bucket
	s := s3.New(auth, awsRegions[region])
	bucket := s.Bucket(bucketName)

	done := make(chan bool)
	files, _ := ioutil.ReadDir(directory)
	for _, f := range files {
		// send off goroutines doing uploads
		go upload(directory, f, bucket, permissions[permission], done)
	}

	// wait for all goroutines to finish
	for _ = range files {
		<-done
	}
	log.Println("done")
}
