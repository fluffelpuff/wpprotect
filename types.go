package main

import "html/template"

type ThirdPartyProvider struct {
}

type LangSpeficData struct {
	WebsiteLang                       string
	EMailUsernameInputFieldHiddenText string
	PasswordInputFieldHiddenText      string
	LogonButtonText                   string
	HelpTextUnderLogonButton          template.HTML
}

type PageConfigData struct {
	Lang                   *LangSpeficData
	PageTitel              string
	ModalTitle             string
	BsThemeDark            string
	WebsiteName            string
	HasThirdPartyProviders bool
	ThirdPartyProviders    []*ThirdPartyProvider
}
