
// data type map, which maps unique keys to values

package main

import(
	"fmt"
)

func main() {

	var mp map[string]string            /* declare a variable, by default map will be nil */
	mp = make(map[string]string)        /* define the map, as nil map can not be assigned any value*/

	mp["india"] = "delhi"
	mp["punjab"] = "CHD"

	for key := range mp {                // one arg to left of range gives key
		fmt.Println(key , mp[key])
	}

	for keys, vals := range mp {       // two args to left of range gives key,value pairs
		fmt.Println(keys , vals)
	}

	value, found_status := mp["usa"]   // check for existence   // if found, second param is "true", else "false"
	if found_status {
		fmt.Println(value)
	}else {
		fmt.Println("Key Not found")
	}

	value, found_status = mp["punjab"]
	if found_status {
		fmt.Println(value)
	}else {
		fmt.Println("Key Not found")
	}

	_ , status := mp["punjab"]          // check for existence
	if status{
		delete(mp , "punjab")          // delete(map_name, key_name)
		fmt.Println("Deleted")
	}else {
		fmt.Println("Key not found")
	}
}