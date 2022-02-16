package main

import (
	"testing"
	"plugin"
	"log"
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
	require.Equal(t, typeId(), FILE_PLUGIN_ID)

	v, err = p.Lookup("NewKey")
	require.NoError(t, err)
	
	newKey, ok := v.(func(*pb.KeySpec) (*pb.KeyRef, error))
	require.Equal(t, ok, true)
	spec := pb.KeySpec{
		Label: "foo123",
	}
	
	ref, err := newKey(&spec)
	require.NoError(t, err)
	
	log.Printf("Label: %v", ref.Label)
}

