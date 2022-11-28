package config

type Config struct {
	Http  HttpOptions  `mapstructure:"http"  json:"http" yaml:"http"`
	Mysql MysqlOptions `mapstructure:"mysql"  json:"mysql" yaml:"mysql"`
	Ldap  LdapOptions  `mapstructure:"ldap"  json:"ldap" yaml:"ldap"`
}

type HttpOptions struct {
	Mode   string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Listen int    `mapstructure:"listen" json:"listen" yaml:"listen"`
}

type MysqlOptions struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type LdapOptions struct {
	Host            string             `mapstructure:"host" json:"host" yaml:"host"`
	Port            int                `mapstructure:"port" json:"port" yaml:"port"`
	BaseDn          string             `mapstructure:"baseDn" json:"baseDn" yaml:"baseDn"`
	BindUser        string             `mapstructure:"bindUser" json:"bindUser" yaml:"bindUser"`
	BindPass        string             `mapstructure:"bindPass" json:"bindPass" yaml:"bindPass"`
	AuthFilter      string             `mapstructure:"authFilter" json:"authFilter" yaml:"authFilter"`
	Attributes      *AttributesOptions `mapstructure:"attributes" json:"attributes" yaml:"attributes"`
	CoverAttributes bool               `mapstructure:"coverAttributes" json:"coverAttributes" yaml:"coverAttributes"`
	AutoRegist      bool               `mapstructure:"autoRegist" json:"autoRegist" yaml:"autoRegist"`
	Tls             bool               `mapstructure:"tls" json:"tls" yaml:"tls"`
	StartTLS        bool               `mapstructure:"startTLS" json:"startTLS" yaml:"startTLS"`
}

type AttributesOptions struct {
	NickName string `mapstructure:"nickName" json:"nickName" yaml:"nickName"`
	Email    string `mapstructure:"email" json:"email" yaml:"email"`
	Phone    string `mapstructure:"phone" json:"phone" yaml:"phone"`
	Uid      string `mapstructure:"uid" json:"uid" yaml:"uid"`
}
