namespace go auth

service authSer {
	string Register(1:string name, 2:string pwd)
}