package file

func IsFile(file string) bool{

	if len(file)>0 {
		return  true
	}
	return false
}
