package infrastructure

//Init nitialize configuration file
func Init() (err error) {
	err = initDatabase("conf/databases.conf", 50)
	if err != nil {
		return
	}
	return
}
