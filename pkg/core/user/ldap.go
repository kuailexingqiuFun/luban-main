package user

import (
	"crypto/tls"
	"fmt"
	"github.com/dnsjia/luban/cmd/options"
	"github.com/dnsjia/luban/pkg/model"
	"github.com/go-ldap/ldap/v3"
	"gorm.io/gorm"
)

func LdapReq(username, password string) (*ldap.SearchResult, error) {
	var (
		conn *ldap.Conn
		err  error
	)
	// 拼接ldap地址
	lc := options.Config.Ldap
	addr := fmt.Sprintf("%s:%d", lc.Host, lc.Port)
	if lc.Tls {
		conn, err = ldap.DialTLS("tcp", addr, &tls.Config{InsecureSkipVerify: true})
	} else {
		conn, err = ldap.Dial("tcp", addr)
	}

	if err != nil {
		return nil, fmt.Errorf("ldap error, cannot dial ldap(%s), %v", addr, err)
	}
	defer conn.Close()

	// 登陆ldap的管理用户
	if lc.BindUser != "" {
		if err := conn.Bind(lc.BindUser, lc.BindPass); err != nil {
			return nil, fmt.Errorf("ldap error, bind ldap fail: %v, use user(%s) to bind", err, lc.BindUser)
		}
	}

	// search dn
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.AuthFilter, username),
		genLdapAttributeSearchList(), // list attributes to retrieve
		nil,
	)
	// 根据searchRequest条件搜索
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("ldap search fail: %v", err)
	}
	// 判断条目是否为0， 为0则用户不存在
	if len(sr.Entries) == 0 {
		return nil, fmt.Errorf("ldap auth fail, no such user")
	}
	// 若是大于1，则存在多条记录
	if len(sr.Entries) > 1 {
		return nil, fmt.Errorf("Internal server error, try again later please")
	}
	// ldap用户密码
	if err := conn.Bind(sr.Entries[0].DN, password); err != nil {
		return nil, fmt.Errorf("Login fail, check your username and password, %w", err)
	}

	return sr, nil
}

var LDAP model.LdapSection

func genLdapAttributeSearchList() []string {
	var ldapAttributes []string
	attrs := LDAP.Attributes
	if attrs.Nickname != "" {
		ldapAttributes = append(ldapAttributes, attrs.Nickname)
	}
	if attrs.Email != "" {
		ldapAttributes = append(ldapAttributes, attrs.Email)
	}
	if attrs.Phone != "" {
		ldapAttributes = append(ldapAttributes, attrs.Phone)
	}
	if attrs.UID != "" {
		ldapAttributes = append(ldapAttributes, attrs.UID)
	}
	return ldapAttributes
}

func LdapLogin(username, password string) (*model.User, error) {
	sr, err := LdapReq(username, password)
	if err != nil {
		return nil, err
	}

	var user model.User
	attrs := model.LDAP.Attributes
	if attrs.Nickname != "" {
		user.NickName = sr.Entries[0].GetAttributeValue(attrs.Nickname)
	}
	if attrs.Email != "" {
		user.Email = sr.Entries[0].GetAttributeValue(attrs.Email)
	}
	if attrs.Phone != "" {
		user.Email = sr.Entries[0].GetAttributeValue(attrs.Email)
	}

	// 更新数据库
	if err := options.DB.Model(&user).Where("username = ?", username).First(&user).Error; err == gorm.ErrRecordNotFound {
		user.Username = username
		user.Password = password
		if err = options.DB.Model(&user).Create(&user).Error; err != nil {
			return nil, fmt.Errorf("login fail， err:%w", err)
		}
		return &user, nil
	}

	return &user, nil
}
