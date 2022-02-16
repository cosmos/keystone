package main

import (
	"testing"
	"plugin"
	"github.com/stretchr/testify/require"
)

const PKCS11_PLUGIN_PATH = "./pkcs11_keys.so"
const PKCS11_PLUGIN_ID = Plugin_Type_Pkcs11_Id

func TestPlugin(t *testing.T) {
	p, err := plugin.Open( PKCS11_PLUGIN_PATH )
	require.NoError(t, err)
	
	v, err := p.Lookup("TypeIdentifier")
	
	typeId, ok := v.(func() string)
	require.Equal(t, ok, true)
	require.NotZero(t, len(typeId()))
	require.Equal(t, typeId(), PKCS11_PLUGIN_ID)
}

