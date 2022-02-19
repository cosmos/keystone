package main

import (
	"testing"
	"plugin"
	"github.com/stretchr/testify/require"

	pb "github.com/regen-network/keystone2/keystone"
	krplugin "github.com/regen-network/keystone2/plugin"
)

const FILE_PLUGIN_PATH = "./file_keys.so"
const FILE_PLUGIN_ID = Plugin_Type_File_Id

func TestPlugin(t *testing.T) {
	p, err := plugin.Open( FILE_PLUGIN_PATH )
	require.NoError(t, err)
	
	v, err := p.Lookup("TypeIdentifier")
	require.NoError(t, err)
	
	typeId, ok := v.(func() string)
	require.Equal(t, ok, true)
	require.NotZero(t, len(typeId()))
	require.Equal(t, typeId(), FILE_PLUGIN_ID)

	v, err = p.Lookup("Init")
	require.NoError(t, err)

	filePlugin, err := v.(func(string) (kr krplugin.Plugin, err error))("./keys")
	require.NoError(t, err)

	spec := pb.KeySpec{
		Label: "foo123",
		Algo: pb.KeygenAlgorithm_KEYGEN_SECP256R1,
	}
	
	ref, err := filePlugin.NewKey(&spec)
	require.NoError(t, err)
	
	t.Logf("Label: %v", ref.Label)
}

