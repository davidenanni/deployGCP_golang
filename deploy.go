package main

import (
    "context"
    gcp "./googleCloudPlatform"
)

func main() {

    ctx := context.Background()

    projectID := "tesidavidenanni"
    bucketName1 := "dav-bucket5"
    bucketName2 := "dav-bucket6"


    crds := gcp.Credentials{projectID, ctx}
    
   
    gcp.CreateBucket(crds,bucketName1,"Regional","us-east1")
    gcp.CreateBucket(crds,bucketName2,"Regional","us-east1")

}
