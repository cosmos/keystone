package main

import (
	"testing"
	"plugin"
	"github.com/stretchr/testify/require"

	pb "github.com/regen-network/keystone2/keystone"
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
	require.Equal(typeId(), Plugin_Type_File_Id)

	v, err = p.Lookup("NewKey")
	require.NoError(t, err)
	
	newKey, ok := v.(func() string)
	require.Equal(t, ok, true)	
}

