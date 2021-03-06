package operator

import (
	"github.com/spf13/cobra"

	"github.com/openshift/cluster-authentication-operator/pkg/operator"
	"github.com/openshift/cluster-authentication-operator/pkg/version"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
)

const componentName = "cluster-authentication-operator"

func NewOperator() *cobra.Command {
	cmd := controllercmd.NewControllerCommandConfig(componentName, version.Get(), operator.RunOperator).NewCommand()
	cmd.Use = "operator"
	cmd.Short = "Start the Authentication Operator"
	return cmd
}
