package user


/*func (s *service) QueryUserByNname(userName string) (ret *user.User,err error) {
	//queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`
	o := db.GetDB()

	//err := o.Model(&user.User{}).Where(&user.User{Name:userName}).First(&ret)
	ret = &user.User{}
	o.Model(ret).First(ret)

	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}

	return
}*/

