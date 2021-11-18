package main

import "github.com/go-gandi/go-gandi/simplehosting"

type simpleHostingCmd struct {
	Instance instanceCmd `kong:"cmd,help='Simple Hosting instance commands'"`
}

type instanceCmd struct {
	List   instanceListCmd   `kong:"cmd,help='List Simple Hosting instances'"`
	Delete instanceDeleteCmd `kong:"cmd,help='Delete a Simple Hosting instance'"`
	Create instanceCreateCmd `kong:"cmd,help='Create a Simple Hosting instance'"`
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
