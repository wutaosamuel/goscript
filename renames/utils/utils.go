package utils

import (
	"os"
)

// IsExist check if file or dir exist
/*
 *	IsExist 
 *		* path 		 			 <- absolute file path
 *		- file exist 		 -> true, nil
 *		- file not exist -> false, nil
 *		- error					 -> _, error
 */
func IsExist(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// IsFile check if it is a file, not a dir
/* IsFile
 *  * name									<- file name with absolute file path
 *  - is file		 						-> true, nil
 *  - is not file or exist	-> false, nil
 *  - error 								-> _, err
 */
func IsFile(name string) (bool, error) {
	f, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	return !f.IsDir(), nil
}

// IsDir check if is a Dir, not a file
/* IsDir
 *  * dir									<- directory path with absolute path 
 *  - is dir 							-> true, nil
 *  - is not dir or exist	-> false, nil
 *  - error 							-> _, err
 */
func IsDir(dir string) (bool, error) {
	d, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	return d.IsDir(), nil
}