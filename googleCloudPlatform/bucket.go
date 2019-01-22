package googleCloudPlatform

import (  
        "log"
        "cloud.google.com/go/storage"
)


func CreateBucket(crds Credentials, bucketName, storageClass, location string){

    client, err := storage.NewClient(crds.Ctx)
    if (err != nil) {
        log.Fatalf("[ERROR] Failed to create client: %v", err)
    }

    
    bucket := client.Bucket(bucketName)
    err     = bucket.Create(crds.Ctx, crds.ProjectID, &storage.BucketAttrs{
                StorageClass: storageClass,
                Location:     location,
})
    if (err != nil) {
        log.Fatalf("[ERROR] Failed to create bucket: %v", err)
    }

    log.Printf("[BUCKET CREATED] %v [CLASS] %v [LOCATION] %v \n", bucketName, storageClass, location)
}
