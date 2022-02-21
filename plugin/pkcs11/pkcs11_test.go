package main

import (
	"testing"
	"plugin"
	"github.com/stretchr/testify/require"
	krplugin "github.com/regen-network/keystone2/plugin"
	pb "github.com/regen-network/keystone2/keystone"
)

const PKCS11_PLUGIN_PATH = "./pkcs11_keys.so"
const PKCS11_PLUGIN_ID = krplugin.Plugin_Type_Pkcs11_Id

func TestPlugin(t *testing.T) {
	p, err := plugin.Open( PKCS11_PLUGIN_PATH )
	require.NoError(t, err)
	
	v, err := p.Lookup("TypeIdentifier")
	
	typeId, ok := v.(func() string)
	require.Equal(t, ok, true)
	require.NotZero(t, len(typeId()))
	require.Equal(t, typeId(), PKCS11_PLUGIN_ID)

	v, err = p.Lookup("Init")
	require.NoError(t, err)

	pkcs11Plugin, err := v.(func(string) (kr krplugin.Plugin, err error))("/home/johnk/src/keystoned2/pkcs11-config")
	require.NoError(t, err)

	spec := pb.KeySpec{
		Label: "foo123ab",
		Algo: pb.KeygenAlgorithm_KEYGEN_SECP256R1,
	}
	
	ref, err := pkcs11Plugin.NewKey(&spec)
	require.NoError(t, err)
	
	t.Logf("Label: %v", ref.Label)

}

