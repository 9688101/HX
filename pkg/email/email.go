package email

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/9688101/HX/config"
)

// SendEmailWithConfig 根据传入的配置发送邮件
func SendEmailWithConfig(mailCfg config.MailConfig, subject, receiver, content string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	// 兼容处理：如果 SMTPFrom 为空，则使用 SMTPAccount
	if mailCfg.SMTPFrom == "" {
		mailCfg.SMTPFrom = mailCfg.SMTPAccount
	}

	// 提取域名，用于生成 Message-ID
	domain := ""
	if parts := strings.Split(mailCfg.SMTPFrom, "@"); len(parts) > 1 {
		domain = parts[1]
	}

	// 生成唯一 Message-ID
	messageID, err := generateMessageID(domain)
	if err != nil {
		return err
	}

	mail, err := generateMail(mailCfg, subject, receiver, content, messageID)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", mailCfg.SMTPAccount, mailCfg.SMTPToken, mailCfg.SMTPServer)
	addr := fmt.Sprintf("%s:%d", mailCfg.SMTPServer, mailCfg.SMTPPort)

	return sendMailWithClient(mailCfg, addr, auth, receiver, mail)
}

// shouldAuth 判断是否需要进行 SMTP 验证
func shouldAuth(mailCfg config.MailConfig) bool {
	return mailCfg.SMTPAccount != "" || mailCfg.SMTPToken != ""
}

// generateMessageID 生成唯一的 Message-ID
func generateMessageID(domain string) (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("<%x@%s>", buf, domain), nil
}

// generateMail 构建邮件消息
func generateMail(mailCfg config.MailConfig, subject, receiver, content, messageID string) ([]byte, error) {
	encodedSubject := fmt.Sprintf("=?UTF-8?B?%s?=", base64.StdEncoding.EncodeToString([]byte(subject)))
	// 使用系统名称作为发件人显示名称
	systemName := "System" // 或者从 config.GetSystemConfig().SystemName 获取
	mail := []byte(fmt.Sprintf("To: %s\r\n"+
		"From: %s<%s>\r\n"+
		"Subject: %s\r\n"+
		"Message-ID: %s\r\n"+
		"Date: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n%s\r\n",
		receiver,
		systemName, mailCfg.SMTPFrom,
		encodedSubject,
		messageID,
		time.Now().Format(time.RFC1123Z),
		content))
	return mail, nil
}

// sendMailWithClient 根据 SMTP 端口选择连接方式，并发送邮件
func sendMailWithClient(mailCfg config.MailConfig, addr string, auth smtp.Auth, receiver string, mail []byte) error {
	var conn net.Conn
	var err error
	// 465 端口一般使用 TLS 连接
	if mailCfg.SMTPPort == 465 {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         mailCfg.SMTPServer,
		}
		conn, err = tls.Dial("tcp", addr, tlsConfig)
	} else {
		conn, err = net.Dial("tcp", addr)
	}
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, mailCfg.SMTPServer)
	if err != nil {
		return err
	}
	defer client.Close()

	if shouldAuth(mailCfg) {
		if err = client.Auth(auth); err != nil {
			return err
		}
	}

	if err = client.Mail(mailCfg.SMTPFrom); err != nil {
		return err
	}

	receiverEmails := strings.Split(receiver, ";")
	for _, email := range receiverEmails {
		if err = client.Rcpt(email); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	if _, err = w.Write(mail); err != nil {
		return err
	}

	return w.Close()
}

// SendEmail 使用全局 SMTP 配置发送邮件
func SendEmail(subject, receiver, content string) error {
	// 从全局配置中获取 MailConfig
	mailCfg := config.GetMailConfig() // 假设 GetSMTPConfig 返回 config.MailConfig 类型
	return SendEmailWithConfig(*mailCfg, subject, receiver, content)
}
