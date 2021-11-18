package main

import "github.com/go-gandi/go-gandi/simplehosting"

type simpleHostingCmd struct {
	Instance instanceCmd `kong:"cmd,help='Simple Hosting instance commands'"`
	Vhost    vhostCmd    `kong:"cmd,help='Simple Hosting vhost commands'"`
}

type instanceCmd struct {
	List   instanceListCmd   `kong:"cmd,help='List Simple Hosting instances'"`
	Delete instanceDeleteCmd `kong:"cmd,help='Delete a Simple Hosting instance'"`
	Create instanceCreateCmd `kong:"cmd,help='Create a Simple Hosting instance'"`
}

type vhostCmd struct {
	List   vhostListCmd   `kong:"cmd,help='List Simple Hosting vhosts'"`
	Create vhostCreateCmd `kong:"cmd,help='Create a Simple Hosting vhost'"`
	Delete vhostDeleteCmd `kong:"cmd,help='Delete a Simple Hosting vhost'"`
}

type instanceListCmd struct{}
type instanceDeleteCmd struct {
	InstanceId string `kong:"arg,help='The ID of a Simple Hosting instance'"`
}
type instanceCreateCmd struct {
	Name     string `kong:"arg,help='The name of a Simple Hosting instance'"`
	Location string `kong:"arg,help='The location of a Simple Hosting instance'"`
	Database string `kong:"arg,help='The database type of a Simple Hosting instance'"`
	Language string `kong:"arg,help='The language type of a Simple Hosting instance'"`
}

func (cmd *instanceListCmd) Run(g *globals) error {
	s := g.simpleHostingHandle
	return jsonPrint(s.ListInstances())
}

func (cmd *instanceDeleteCmd) Run(g *globals) error {
	s := g.simpleHostingHandle
	return jsonPrint(s.DeleteInstance(cmd.InstanceId))
}

func (cmd *instanceCreateCmd) Run(g *globals) error {
	s := g.simpleHostingHandle
	return jsonPrint(s.CreateInstance(
		simplehosting.CreateInstanceRequest{
			Name:     cmd.Name,
			Location: cmd.Location,
			Type: &simplehosting.InstanceType{
				Database: &simplehosting.Database{
					Name: cmd.Database,
				},
				Language: &simplehosting.Language{
					Name: cmd.Language,
				},
			},
		}))
}

type vhostListCmd struct {
	InstanceId string `kong:"arg,help='The ID of a Simple Hosting instance'"`
}
type vhostCreateCmd struct {
	InstanceId string `kong:"arg,help='The ID of a Simple Hosting instance'"`
	FQDN       string `kong:"arg,help='The Vhost FQDN'"`
}
type vhostDeleteCmd struct {
	InstanceId string `kong:"arg,help='The ID of a Simple Hosting instance'"`
	FQDN       string `kong:"arg,help='The Vhost FQDN'"`
}

func (cmd *vhostListCmd) Run(g *globals) error {
	s := g.simpleHostingHandle
	return jsonPrint(s.ListVhosts(cmd.InstanceId))
}

func (cmd *vhostCreateCmd) Run(g *globals) error {
	s := g.simpleHostingHandle
	return jsonPrint(s.CreateVhost(
		cmd.InstanceId,
		simplehosting.CreateVhostRequest{
			FQDN: cmd.FQDN,
		}))
}

func (cmd *vhostDeleteCmd) Run(g *globals) error {
	s := g.simpleHostingHandle
	return jsonPrint(s.DeleteVhost(
		cmd.InstanceId,
		cmd.FQDN,
	))
}
