package email

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
)

func SendEmail()error{

	m:=gomail.NewMessage() //实例化gomail
	m.SetHeader("From","15001150581@163.com")
	m.SetHeader("To","wangxin@microdreams.com","dongliang@microdreams.com") //发送给谁 可写多个
	//m.SetAddressHeader("Cc","xxx@163.com") //抄送
	m.SetHeader("Subject","测试") //主题
	m.SetBody("text/html","This is 测试邮件")//内容
	m.Attach("Book1.xlsx")//附件
	d := gomail.NewDialer("smtp.163.com", 25, "15001150581@163.com", "ORURJARPCKUJOEGI")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}