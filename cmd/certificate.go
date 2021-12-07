package main

import "github.com/go-gandi/go-gandi/certificate"

type certificateCmd struct {
	List        certificateListCmd        `kong:"cmd,help='List certificates'"`
	Delete      certificateDeleteCmd      `kong:"cmd,help='Delete a certificate'"`
	Create      certificateCreateCmd      `kong:"cmd,help='Create a certificate'"`
	ListPackage certificateListPackageCmd `kong:"cmd,help='List certificate packages'"`
}

type certificateListCmd struct{}
type certificateDeleteCmd struct {
	CertificateId string `kong:"arg,help='The certificate ID'"`
}
type certificateCreateCmd struct {
	CN      string `kong:"arg,help='The certificate CN'"`
	Package string `kong:"arg,help='The certificate package (available packages can be listed with the list-package command)'"`
}

type certificateListPackageCmd struct{}

func (cmd *certificateListCmd) Run(g *globals) error {
	s := g.certificateHandle
	return jsonPrint(s.ListCertificates())
}

func (cmd *certificateDeleteCmd) Run(g *globals) error {
	s := g.certificateHandle
	return jsonPrint(s.DeleteCertificate(cmd.CertificateId))
}

func (cmd *certificateCreateCmd) Run(g *globals) error {
	s := g.certificateHandle
	return jsonPrint(s.CreateCertificate(
		certificate.CreateCertificateRequest{
			CN:      cmd.CN,
			Package: cmd.Package,
		}))
}

func (cmd *certificateListPackageCmd) Run(g *globals) error {
	s := g.certificateHandle
	return jsonPrint(s.ListPackages())
}
