package model

type Business struct {
	Id       int
	Name     string
	Phone    string
	Rating   int
	Price    string

}

func GetAllBusinesses () (businesses []Business, err error) {
	businesses = []Business{
		{1,"Elf","8542244",1,"$$$"},
		{2,"Esso","8546666",2,"$"},
		{3,"Shell","8545647",3,"$$"},
		{4,"BP","1525453",4,"$$$"},
		{5,"Lukoil","1542765",5,"$$"},
		{6,"Neste","8542377",6,"$$"},
		{7,"Mobil1","2366544",7,"$"},
		{8,"ZIC","8544454",8,"$$$"},
		{9,"Gazprom","5678899",9,"$"},
	}
	return
}