package aws

func S3(bucket string) bool{

	if len(bucket)>0{
		return true
	}

	return false
}
